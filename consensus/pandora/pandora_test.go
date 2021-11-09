package pandora

import (
	"context"
	"fmt"
	"math/big"
	"net"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	bls_common "github.com/prysmaticlabs/prysm/shared/bls/common"
	"github.com/prysmaticlabs/prysm/shared/bls/herumi"
	"github.com/stretchr/testify/assert"
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
		pandoraEngine, _ := createDummyPandora(t)
		dummyEndpoint := ipcTestLocation
		pandoraEngine.endpoint = dummyEndpoint
		pandoraEngine.Start(nil)

		t.Run("should handle sharding info for block number less than 2", func(t *testing.T) {
			errChannel := make(chan error)
			resChannel := make(chan [4]string)
			pandoraEngine.fetchShardingInfoCh <- &shardingInfoReq{
				slot:        0,
				epoch:       0,
				blockNumber: 0,
				parentHash:  common.Hash{},
				errc:        errChannel,
				res:         resChannel,
			}

			select {
			case err := <-errChannel:
				assert.Nil(t, err)
			case response := <-resChannel:
				assert.Equal(t, "0x2e07d67c7eebfc74fcc5ca07f25d995a3a67c497deaee9d60fb0b26b549a8d3b", response[0])
				assert.Equal(t, "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421", response[1])
				assert.Equal(t, "0xf901f1a00000000000000000000000000000000000000000000000000000000000000000a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347940000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000808080808084c3808080a00000000000000000000000000000000000000000000000000000000000000000880000000000000000", response[2])
				assert.Equal(t, "0x", response[3])
				break
			}

			pandoraEngine.fetchShardingInfoCh <- &shardingInfoReq{
				slot:        1,
				epoch:       0,
				blockNumber: 1,
				parentHash:  common.Hash{},
				errc:        errChannel,
				res:         resChannel,
			}

			select {
			case err := <-errChannel:
				assert.Nil(t, err)
			case response := <-resChannel:
				assert.Equal(t, "0x4203bd2ead3d91541fd49fb40cc2a16bce624be5074917d9d28ce6e164860197", response[0])
				assert.Equal(t, "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421", response[1])
				assert.Equal(t, "0xf901f1a00000000000000000000000000000000000000000000000000000000000000000a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347940000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000808080808084c3018001a00000000000000000000000000000000000000000000000000000000000000000880000000000000000", response[2])
				assert.Equal(t, "0x", response[3])
				break
			}
		})

		t.Run(
			"should return err when parent hash does not match and block number is greater than 1",
			func(t *testing.T) {
				errChannel := make(chan error)
				resChannel := make(chan [4]string)

				pandoraEngine.fetchShardingInfoCh <- &shardingInfoReq{
					slot:        3,
					epoch:       0,
					blockNumber: 2,
					parentHash:  common.HexToHash("0x4203bd2ead3d91541fd49fb40cc2a16bce624be5074917d9d28ce6e164860192"),
					errc:        errChannel,
					res:         resChannel,
				}

				err := <-errChannel
				assert.Equal(t, errInvalidParentHash, err)
			})

		t.Run(
			"should return err when block number does not match and block number is greater than 1",
			func(t *testing.T) {
				errChannel := make(chan error)
				resChannel := make(chan [4]string)

				pandoraEngine.fetchShardingInfoCh <- &shardingInfoReq{
					slot:        3,
					epoch:       0,
					blockNumber: 2,
					parentHash:  common.Hash{},
					errc:        errChannel,
					res:         resChannel,
				}

				err := <-errChannel
				assert.Equal(t, errInvalidBlockNumber, err)
			})

		t.Run("should handle sharding info when block is greater than 1", func(t *testing.T) {
			errChannel := make(chan error)
			resChannel := make(chan [4]string)
			header := &types.Header{
				Number:     big.NewInt(2),
				ParentHash: common.HexToHash("0x4203bd2ead3d91541fd49fb40cc2a16bce624be5074917d9d28ce6e164860192"),
			}
			block := types.NewBlock(header, nil, nil, nil, nil)
			pandoraEngine.setCurrentBlock(block)

			pandoraEngine.fetchShardingInfoCh <- &shardingInfoReq{
				slot:        3,
				epoch:       0,
				blockNumber: header.Number.Uint64(),
				parentHash:  header.ParentHash,
				errc:        errChannel,
				res:         resChannel,
			}

			select {
			case err := <-errChannel:
				assert.Nil(t, err)
			case response := <-resChannel:
				assert.Equal(t, "0xe64f708a495767942054021a27efd25730ba1a2fef2dc6f68503ed8dc36d3951", response[0])
				assert.Equal(t, "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421", response[1])
				assert.Equal(t, "0xf901f1a04203bd2ead3d91541fd49fb40cc2a16bce624be5074917d9d28ce6e164860192a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347940000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b9010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800280808084c3038003a00000000000000000000000000000000000000000000000000000000000000000880000000000000000", response[2])
				assert.Equal(t, "0x02", response[3])
				break
			}
		})

		t.Run("should return err when there is no sharding work", func(t *testing.T) {
			pandoraEngine.currentBlock = nil
			errChannel := make(chan error)
			pandoraEngine.fetchShardingInfoCh <- &shardingInfoReq{
				slot:        0,
				epoch:       0,
				blockNumber: 0,
				parentHash:  common.Hash{},
				errc:        errChannel,
				res:         nil,
			}

			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()

			select {
			case <-ticker.C:
				assert.Fail(t, "should receive error that there was no sharding block")
				ticker.Stop()
			case err := <-errChannel:
				assert.NotNil(t, err)
				assert.Equal(t, errNoShardingBlock, err)
				ticker.Stop()
			}
		})
	})

	t.Run("should handle submitSignatureData", func(t *testing.T) {
		pandoraEngine, _ := createDummyPandora(t)
		dummyEndpoint := ipcTestLocation
		pandoraEngine.endpoint = dummyEndpoint
		pandoraEngine.Start(nil)

		t.Run("should react to invalid work", func(t *testing.T) {
			shardingWorkHash := common.Hash{}
			errChannel := make(chan error)
			pandoraEngine.submitShardingInfoCh <- &shardingResult{
				nonce:   types.BlockNonce{},
				hash:    shardingWorkHash,
				blsSeal: nil,
				errc:    errChannel,
			}

			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()

			select {
			case <-ticker.C:
				assert.Fail(t, "should receive error that work was not submitted")
				ticker.Stop()
			case err := <-errChannel:
				assert.NotNil(t, err)
				assert.Equal(t, "Work submitted but none pending", err.Error())
				ticker.Stop()
			}
		})

		t.Run("should pass valid work", func(t *testing.T) {
			errChannel := make(chan error)
			expectedBlockNumber := int64(1)
			firstHeader := &types.Header{Number: big.NewInt(expectedBlockNumber)}
			firstHeaderExtra := &ExtraData{
				Slot:  1,
				Epoch: 0,
				Turn:  1,
			}
			firstHeaderExtraBytes, err := rlp.EncodeToBytes(firstHeaderExtra)
			assert.Nil(t, err)
			firstHeader.Extra = firstHeaderExtraBytes

			blsSeal := &BlsSignatureBytes{}
			privKeyHex := "0x3a3cae36df7b66019442f8f38acb080c680780dc8ee2d430cc801903be0b651e"
			privKey, err := herumi.SecretKeyFromBytes(hexutil.MustDecode(privKeyHex))
			assert.Nil(t, err)

			publicKeys := [32]bls_common.PublicKey{}
			publicKeys[1] = privKey.PublicKey()

			pandoraEngine.epochInfos.Add(uint64(0), &EpochInfo{
				Epoch:            0,
				ValidatorList:    publicKeys,
				EpochTimeStart:   0,
				SlotTimeDuration: DefaultSlotTimeDuration,
			})

			firstBlock := types.NewBlock(firstHeader, nil, nil, nil, nil)

			// I dont know why but in test I couldnt match sealhash via SealHash() so I do it manually
			expectedSealHash := hexutil.MustDecode("0x85d2ed6f97f0e8ebfcc7f1badaf6522748b2488a45cc90f56c6c48a7290658f6")
			shardingWorkHash := common.BytesToHash(expectedSealHash)
			pandoraEngine.works[shardingWorkHash] = firstBlock
			fullResultsChan := make(chan *types.Block)
			pandoraEngine.results = fullResultsChan
			signature := privKey.Sign(shardingWorkHash.Bytes())
			copy(blsSeal[:], signature.Marshal())
			signatureFromBytes, err := herumi.SignatureFromBytes(blsSeal[:])

			assert.Nil(t, err)
			assert.NotNil(t, signatureFromBytes)

			signature, err = herumi.SignatureFromBytes(blsSeal[:])
			assert.Nil(t, err)
			assert.True(t, signature.Verify(publicKeys[1], expectedSealHash))

			pandoraEngine.submitShardingInfoCh <- &shardingResult{
				nonce:   types.BlockNonce{},
				hash:    shardingWorkHash,
				blsSeal: blsSeal,
				errc:    errChannel,
			}

			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()

			select {
			case block := <-fullResultsChan:
				assert.Equal(t, firstHeader.Number, block.Number())
			case <-ticker.C:
				assert.Fail(t, "should receive error that work was not submitted")
				ticker.Stop()
			case err := <-errChannel:
				assert.Nil(t, err)
				ticker.Stop()
			}
		})
	})

	t.Run("should handle subscriptionErrCh", func(t *testing.T) {
		pandoraEngine, _ := createDummyPandora(t)
		dummyEndpoint := ipcTestLocation
		pandoraEngine.endpoint = dummyEndpoint
		pandoraEngine.Start(nil)
		time.Sleep(time.Millisecond * 100)
		dummyErr := error(nil)
		pandoraEngine.subscriptionErrCh <- dummyErr
		time.Sleep(reConPeriod)
		assert.Equal(t, dummyErr, pandoraEngine.runError)
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

	listener, server, err := rpc.StartIPCEndpoint(location, apis)
	assert.NoError(t, err)

	return
}
