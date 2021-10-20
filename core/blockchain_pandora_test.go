package core

import (
	"context"
	"math/big"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/pandora_orcclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	lru "github.com/hashicorp/golang-lru"
)

type Orchestrator struct {
	ctx context.Context
}

func fakeOrcServer(ctx context.Context) *rpc.Server {
	server := rpc.NewServer()
	server.RegisterName("orc", &Orchestrator{ctx: ctx})
	return server
}

func (orc *Orchestrator) SteamConfirmedPanBlockHashes(request *pandora_orcclient.BlockHash) (*rpc.Subscription, error) {

	timer := time.NewTicker(2 * time.Second)
	defer timer.Stop()
	notifier, supported := rpc.NotifierFromContext(orc.ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}
	rpcSub := notifier.CreateSubscription()
	go func() {
		for {
			select {
			case <-timer.C:
				retVal := &pandora_orcclient.BlockStatus{Status: pandora_orcclient.Verified}
				retVal.Hash = request.Hash
				notifier.Notify(rpcSub.ID, retVal)
			case <-orc.ctx.Done():
				log.Info("exiting streamConfirmedPanBlockHashes")
				return
			}
		}
	}()
	return rpcSub, nil
}

func fakeOrcClient(ctx context.Context) *pandora_orcclient.OrcClient {
	server := fakeOrcServer(ctx)
	client := rpc.DialInProc(server)
	return pandora_orcclient.NewOrcClient(client)
}

func getWSlinkOfTestOrcServer(ctx context.Context) string {
	server := fakeOrcServer(ctx)
	httpsrv := httptest.NewServer(server.WebsocketHandler(nil))
	return "ws:" + strings.TrimPrefix(httpsrv.URL, "http:")
}

func TestPandoraBlockHashConfirmationFetcher(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	db := rawdb.NewMemoryDatabase()
	gspec := &Genesis{BaseFee: big.NewInt(params.InitialBaseFee)}
	genesis := gspec.MustCommit(db)
	prefix, _ := GenerateChain(params.TestChainConfig, genesis, ethash.NewFaker(), db, 10, func(i int, gen *BlockGen) {})

	chain, _ := NewBlockChain(db, nil, params.AllCliqueProtocolChanges, ethash.NewFaker(), vm.Config{}, nil, nil)

	chain.InsertChain(prefix)

	bc := &BlockChain{blockConfirmationCh: make(chan struct{})}
	t.Run("orchestrator client with ws connection", func(t *testing.T) {
		serverLocation := getWSlinkOfTestOrcServer(ctx)
		//targetHash := types.EmptyRootHash
		orchCache, _ := lru.New(1 << 10)
		bc.orchestratorConfirmationCache = orchCache
		bc.cacheConfig = new(CacheConfig)
		bc.cacheConfig.OrcClientEndpoint = []string{serverLocation}

		err := bc.pandoraBlockHashConfirmationFetcher(ctx)
		if err != nil {
			t.Error("error", err)
		}
	})
}
