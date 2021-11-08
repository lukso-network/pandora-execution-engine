package pandora_orcclient

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/rpc"
)

const (
	MockedHashInvalid = "0xc9a190eb52c18df5ffcb1d817214ecb08f025f8583805cd12064d30e3f9bd9d5"
	MockedHashPending = "0xa99c69a301564970956edd897ff0590f4c0f1031daa464ded655af65ad0906df"
)

// DialInProcRPCClient creates an in process RPC mock client
func DialInProcRPCClient() *OrcClient {
	server := NewMockOrchestratorServer()
	rpcClient := rpc.DialInProc(server)
	return NewOrcClient(rpcClient)
}

// MockOrchestratorService testing mock orchestrator service
type MockOrchestratorService struct{}

// NewMockOrchestratorServer method mock pandora chain apis
func NewMockOrchestratorServer() *rpc.Server {
	server := rpc.NewServer()
	if err := server.RegisterName("orc", new(MockOrchestratorService)); err != nil {
		panic(err)
	}
	return server
}

// ConfirmPanBlockHashes confirms block confirmation
func (OrcClient *MockOrchestratorService) ConfirmPanBlockHashes(ctx context.Context,
	request []*BlockHash) (response []*BlockStatus, err error) {

	if len(request) < 1 {
		err = fmt.Errorf("request has empty slice")

		return
	}
	response = make([]*BlockStatus, 0)

	for _, blockRequest := range request {
		status := Verified

		if MockedHashInvalid == blockRequest.Hash.String() {
			status = Invalid
		}

		if MockedHashPending == blockRequest.Hash.String() {
			status = Pending
		}

		response = append(response, &BlockStatus{
			BlockHash: BlockHash{
				Slot: blockRequest.Slot,
				Hash: blockRequest.Hash,
			},
			Status: status,
		})
	}

	return
}

func (orc *MockOrchestratorService) SteamConfirmedPanBlockHashes(
	ctx context.Context,
	request *BlockHash,
) (*rpc.Subscription, error) {

	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}
	rpcSub := notifier.CreateSubscription()

	hashes := []string{
		"0x436df74e9aeb620e14f53537332b9f641eb1b7870d3cd63c4b68e1408b3f96f8",
		"0x6b28dd08abb32d5a28751e0219ae0506f0a9cae33642ffcbdd5fda05e04dc169",
		"0xa424e271c7bc1c65a8e6864d19f58344f6d3b4d0bcc22279f4f4116e44771b08",
		"0xb4b618b09351a26bc87ec5c74a68d2b5ad368c665f7bf9a7d94b1a20d6d79b47",
		"0x49628542e1c7729dc396c15054eae473957b0e1875ae01883ff941eceb4ff507",
		"0xf1a58167ad8c00f8ed177b8e5f59cf330d33d33222b56b4d9bedd168a10dc098",
		"0x4106d0be79eda3cff3f461a3a038c88ba6acbf2a44e9238e9a243b0f4aae3916",
		"0x91ae125a579af3cbce688be1c56c25581b1250ab2a090085fb751f5e6d1a86a3",
		"0x9b37b5598ea4d6cdcf492ed6eb4f2065ef57ef2f227903107f6c3ca743bea470",
		"0x325632e44c8e18ce47619315c772a4a53327a01e3ff10f67146e549fbb9e5389",
		}

	for i := 0; i < 10; i++ {
		status := &BlockStatus{Status: Verified}
		status.Hash = common.HexToHash(hashes[i])
		notifier.Notify(rpcSub.ID, status)
	}

	return rpcSub, nil
}
