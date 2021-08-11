package pandora

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"math/big"
	"net"
	"testing"
	"time"
)

var (
	dummyRpcFunc = DialRPCFn(func(endpoint string) (rpcClient *rpc.Client, err error) {
		return rpc.Dial(endpoint)
	})

	ipcTestLocation = "./test.ipc"
)

func Test_New(t *testing.T) {
	pandoraEngine, _ := createDummyPandora(t)
	assert.IsType(t, &Pandora{}, pandoraEngine)

	t.Run("should have default settings", func(t *testing.T) {
		assert.Equal(t, &params.PandoraConfig{
			GenesisStartTime: DefaultGenesisStartTime,
			SlotsPerEpoch:    DefaultSlotsPerEpoch,
			SlotTimeDuration: DefaultSlotTimeDuration,
		}, pandoraEngine.config,
		)
	})
}

func TestPandora_Start(t *testing.T) {
	// TODO: in my opinion Start() should return err when failure is present
	t.Run("should not start with empty endpoint", func(t *testing.T) {
		pandoraEngine, _ := createDummyPandora(t)
		pandoraEngine.Start(nil)
		assert.Equal(t, "", pandoraEngine.endpoint)
	})

	t.Run("should mark as running with non-empty endpoint", func(t *testing.T) {
		pandoraEngine, _ := createDummyPandora(t)
		dummyEndpoint := "https://some.endpoint"
		pandoraEngine.endpoint = dummyEndpoint
		pandoraEngine.Start(nil)
		assert.True(t, pandoraEngine.isRunning)
		assert.Nil(t, pandoraEngine.chain)
	})

	_, server, _ := makeOrchestratorServer(t)
	defer server.Stop()

	t.Run("should wait for connection", func(t *testing.T) {
		waitingPandoraEngine, _ := createDummyPandora(t)
		waitingPandoraEngine.endpoint = ipcTestLocation
		ticker := time.NewTicker(reConPeriod)
		defer ticker.Stop()
		dummyError := fmt.Errorf("dummy Error")

		waitingPandoraEngine.dialRPC = func(endpoint string) (*rpc.Client, error) {
			return nil, dummyError
		}

		waitingPandoraEngine.Start(nil)
		t.Log("Waiting for reconnection in Pandora Engine")
		<-ticker.C
		assert.Equal(t, dummyError, waitingPandoraEngine.runError)
		waitingPandoraEngine.dialRPC = dummyRpcFunc
		t.Log("Waiting for reconnection in Pandora Engine")
		<-ticker.C
		t.Log("Waiting for reconnection in Pandora Engine, pointing to orchestrator server")
		<-ticker.C
		assert.Equal(t, uint64(0), waitingPandoraEngine.currentEpoch)
		assert.Nil(t, waitingPandoraEngine.runError)
	})

	t.Run("should handle seal request", func(t *testing.T) {
		pandoraEngine, _ := createDummyPandora(t)
		dummyEndpoint := ipcTestLocation
		pandoraEngine.endpoint = dummyEndpoint
		pandoraEngine.Start(nil)

		expectedBlockNumber := int64(1)
		firstHeader := &types.Header{Number: big.NewInt(expectedBlockNumber)}
		firstBlock := types.NewBlock(firstHeader, nil, nil, nil, nil)
		results := make(chan *types.Block)
		pandoraEngine.newSealRequestCh <- &sealTask{
			block:   firstBlock,
			results: results,
		}

		time.Sleep(time.Millisecond * 50)

		assert.Equal(t, firstBlock.Number(), pandoraEngine.currentBlock.Number())
	})

	t.Run("should handle sharding info request", func(t *testing.T) {

	})

	t.Run("should handle submitSignatureData", func(t *testing.T) {
		pandoraEngine, _ := createDummyPandora(t)
		dummyEndpoint := ipcTestLocation
		pandoraEngine.endpoint = dummyEndpoint
		pandoraEngine.Start(nil)
		errChannel := make(chan error)

		pandoraEngine.submitShardingInfoCh <- &shardingResult{
			nonce:   types.BlockNonce{},
			hash:    common.Hash{},
			blsSeal: nil,
			errc:    errChannel,
		}

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		select {
		case <-ticker.C:
			assert.Fail(t, "should receive error that work was not submitted")
		case err := <-errChannel:
			assert.NotNil(t, err)
			assert.Equal(t, "invalid submit work request", err.Error())
		}
	})

	t.Run("should handle subscriptionErrCh", func(t *testing.T) {
		pandoraEngine, _ := createDummyPandora(t)
		dummyEndpoint := ipcTestLocation
		pandoraEngine.endpoint = dummyEndpoint
		pandoraEngine.Start(nil)
		time.Sleep(time.Millisecond * 100)
		dummyErr := fmt.Errorf("dummyErr")
		pandoraEngine.subscriptionErrCh <- dummyErr
		time.Sleep(reConPeriod)
		assert.Equal(t, dummyErr, pandoraEngine.runError)
		time.Sleep(reConPeriod)
		assert.NotEqual(t, dummyErr, pandoraEngine.runError)
	})

	t.Run("should handle done event", func(t *testing.T) {
		pandoraEngine, cancel := createDummyPandora(t)
		dummyEndpoint := ipcTestLocation
		pandoraEngine.endpoint = dummyEndpoint
		pandoraEngine.Start(nil)
		assert.True(t, pandoraEngine.isRunning)
		time.Sleep(time.Millisecond * 100)
		cancel()
		time.Sleep(time.Millisecond * 50)
		assert.False(t, pandoraEngine.isRunning)
		assert.Nil(t, pandoraEngine.runError)
	})
}

func createDummyPandora(t *testing.T) (pandoraEngine *Pandora, cancel context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	cfg := &params.PandoraConfig{
		GenesisStartTime: 0,
		SlotsPerEpoch:    0,
		SlotTimeDuration: 0,
	}
	urls := make([]string, 2)
	dialGrpcFnc := dummyRpcFunc
	pandoraEngine = New(ctx, cfg, urls, dialGrpcFnc)

	genesisHeader := &types.Header{Number: big.NewInt(0)}
	genesisBlock := types.NewBlock(genesisHeader, nil, nil, nil, nil)
	pandoraEngine.setCurrentBlock(genesisBlock)

	return
}

type OrchestratorApi struct{}

func (orchestratorApi *OrchestratorApi) MinimalConsensusInfo(
	ctx context.Context,
	epoch uint64,
) (subscription *rpc.Subscription, err error) {
	return
}

func makeOrchestratorServer(
	t *testing.T,
) (listener net.Listener, server *rpc.Server, location string) {
	location = ipcTestLocation
	apis := make([]rpc.API, 0)
	api := &OrchestratorApi{}

	apis = append(apis, rpc.API{
		Namespace: "orc",
		Version:   "1.0",
		Service:   api,
		Public:    true,
	})

	// TODO: change to inproc?
	listener, server, err := rpc.StartIPCEndpoint(location, apis)
	assert.NoError(t, err)

	return
}
