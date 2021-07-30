package pandora

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/sha3"
	"sync"
	"time"
)

var (
	reConPeriod             = 15 * time.Second
	DefaultGenesisStartTime = uint64(time.Now().Unix())
	DefaultSlotsPerEpoch    = uint64(32)
	DefaultSlotTimeDuration = 6 * time.Second

	errInvalidValidatorSize = errors.New("invalid length of validator list")
	errInvalidEpochInfo     = errors.New("invalid epoch info")
	errEmptyOrchestratorUrl = errors.New("orchestrator url is empty")
	errNoShardingBlock      = errors.New("no pandora sharding block available yet")
)

// DialRPCFn dials to the given endpoint
type DialRPCFn func(endpoint string) (*rpc.Client, error)

// Pandora
type Pandora struct {
	isRunning      bool
	processingLock sync.RWMutex
	ctx            context.Context
	cancel         context.CancelFunc
	runError       error

	chain             consensus.ChainHeaderReader
	config            *params.PandoraConfig // Consensus engine configuration parameters
	epochInfoCache    *EpochInfoCache
	currentEpoch      uint64
	currentEpochInfo  *EpochInfo
	dialRPC           DialRPCFn
	endpoint          string
	connected         bool
	rpcClient         *rpc.Client
	namespace         string
	subscription      *rpc.ClientSubscription
	subscriptionErrCh chan error

	apiResponse          [4]string
	results              chan<- *types.Block
	fetchShardingInfoCh  chan *shardingInfoReq // Channel used for remote sealer to fetch mining work
	submitShardingInfoCh chan *shardingResult

	lock      sync.Mutex // Ensures thread safety for the in-memory caches and mining fields
	closeOnce sync.Once  // Ensures exit channel will not be closed twice.
}

func New(
	ctx context.Context,
	cfg *params.PandoraConfig,
	chain consensus.ChainHeaderReader,
	urls []string,
	dialRPCFn DialRPCFn,
) (*Pandora, error) {

	ctx, cancel := context.WithCancel(ctx)
	_ = cancel // govet fix for lost cancel. Cancel is handled in service.Stop()
	if cfg.SlotsPerEpoch == 0 {
		cfg.SlotsPerEpoch = DefaultSlotsPerEpoch
	}
	if cfg.GenesisStartTime == 0 {
		cfg.GenesisStartTime = DefaultGenesisStartTime
	}
	if cfg.SlotTimeDuration == 0 {
		cfg.SlotTimeDuration = DefaultSlotTimeDuration
	}
	if urls[0] == "" {
		return nil, errEmptyOrchestratorUrl
	}

	pandora := &Pandora{
		ctx:            ctx,
		cancel:         cancel,
		chain:          chain,
		config:         cfg,
		epochInfoCache: NewEpochInfoCache(),
		dialRPC:        dialRPCFn,
		endpoint:       urls[0],
		namespace:      "orc",

		fetchShardingInfoCh: make(chan *shardingInfoReq),
		subscriptionErrCh:   make(chan error),
	}

	pandora.start()
	return pandora, nil
}

func (p *Pandora) start() {
	// Exit early if pandora endpoint is not set.
	if p.endpoint == "" {
		return
	}
	go func() {
		p.isRunning = true
		p.waitForConnection()
		if p.ctx.Err() != nil {
			log.Info("Context closed, exiting pandora goroutine")
			return
		}
		p.run(p.ctx.Done())
	}()
}

// run subscribes to all the services for the ETH1.0 chain.
func (p *Pandora) run(done <-chan struct{}) {
	log.Debug("Pandora chain service is starting")
	p.runError = nil

	// the loop waits for any error which comes from consensus info subscription
	// if any subscription error happens, it will try to reconnect and re-subscribe with pandora chain again.
	for {
		select {
		case shardingInfoReq := <-p.fetchShardingInfoCh:
			curHeader := getDummyHeader()
			hash := p.SealHash(curHeader)
			shardingInfo := prepareShardingInfo(curHeader, hash)
			shardingInfoReq.res <- shardingInfo

		case submitSignatureData := <-p.submitShardingInfoCh:

		case err := <-p.subscriptionErrCh:
			log.Debug("Got subscription error", "err", err)
			log.Debug("Starting retry to connect and subscribe to orchestrator chain")
			// Try to check the connection and retry to establish the connection
			p.retryToConnectAndSubscribe(err)
			continue
		case <-done:
			p.isRunning = false
			p.runError = nil
			log.Debug("Context closed, exiting goroutine", "ctx", "pandora-consensus")
			return
		}
	}
}

// Close closes the exit channel to notify all backend threads exiting.
func (p *Pandora) Close() error {
	if p.cancel != nil {
		defer p.cancel()
	}
	return nil
}

// SealHash returns the hash of a block prior to it being sealed.
func (p *Pandora) SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	extraData := header.Extra
	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		extraData,
	})
	hasher.Sum(hash[:0])
	return hash
}
