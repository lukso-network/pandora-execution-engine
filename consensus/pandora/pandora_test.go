package pandora

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
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
	pandoraEngine := createDummyPandora(t)
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
	pandoraEngine := createDummyPandora(t)
	dummyEndpoint := "https://some.endpoint"
	// TODO: in my opinion Start() should return err when failure is present
	t.Run("should not start with empty endpoint", func(t *testing.T) {
		pandoraEngine.Start(nil)
		assert.Equal(t, "", pandoraEngine.endpoint)
	})

	t.Run("should mark as running with non-empty endpoint", func(t *testing.T) {
		pandoraEngine.endpoint = dummyEndpoint
		pandoraEngine.Start(nil)

		assert.True(t, pandoraEngine.isRunning)
		assert.Nil(t, pandoraEngine.chain)
	})

	makeOrchestratorServer(t)

	t.Run("should wait for connection", func(t *testing.T) {
		t.Parallel()
		waitingPandoraEngine := createDummyPandora(t)
		waitingPandoraEngine.endpoint = ipcTestLocation
		ticker := time.NewTicker(reConPeriod)
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
		_, server, _ := makeOrchestratorServer(t)
		defer server.Stop()
		t.Log("Waiting for reconnection in Pandora Engine")
		<-ticker.C
		assert.Equal(t, uint64(0), waitingPandoraEngine.currentEpoch)
		assert.Nil(t, waitingPandoraEngine.runError)
	})
}

func createDummyPandora(t *testing.T) (pandoraEngine *Pandora) {
	ctx := context.Background()
	cfg := &params.PandoraConfig{
		GenesisStartTime: 0,
		SlotsPerEpoch:    0,
		SlotTimeDuration: 0,
	}
	urls := make([]string, 2)
	dialGrpcFnc := dummyRpcFunc
	pandoraEngine = New(ctx, cfg, urls, dialGrpcFnc)

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
