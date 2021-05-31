/*
* Developed by: Md. Muhaimin Shah Pahalovi
* Generated: 5/31/21
* This file is generated to support Lukso pandora module.
* Purpose:
 */
package orcclient_pandora

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	Pending Status = iota
	Verified
	Invalid
)

const (
	MockedHashInvalid = "0xc9a190eb52c18df5ffcb1d817214ecb08f025f8583805cd12064d30e3f9bd9d5"
	MockedHashPending = "0xa99c69a301564970956edd897ff0590f4c0f1031daa464ded655af65ad0906df"
)

// Orchestrator link
var orchestrator_link = "http://127.0.0.1:8545"

// In process rpc server related test

// TestOrcClient_GetConfirmedPanBlockHashes tests in process pandora block confirmation
func TestOrcClient_GetConfirmedPanBlockHashes(t *testing.T) {
	orchestrator := DialInProcRPCClient()
	var request []*BlockHash
	request = append(request, &BlockHash{Hash: common.HexToHash(MockedHashInvalid), Slot: 1}, &BlockHash{Hash: common.HexToHash(MockedHashPending), Slot: 2})
	response, err := orchestrator.GetConfirmedPanBlockHashes(context.Background(), request)
	if err != nil {
		t.Fatalf("error found while getting confirmed pending block hashes")
	}
	t.Log("received block confirmations from orchestrator")
	for _, hash := range response {
		t.Logf("received %v", hash)
	}
}

// DialInProcRPCClient creates an in process RPC mock client
func DialInProcRPCClient() *OrcClient {
	server := NewMockOrchestratorServer()
	rpcClient := rpc.DialInProc(server)
	return NewOrcClient(rpcClient)
}

// GIT WON'T ALLOW TO PASS THIS TEST. DO NOT UN COMMENT IT. IT IS JUST FOR TEST IN LOCAL ENVIRONMENT.
// TestOrcClient_GetConfirmedPanBlockHashesWithHTTP Calls real orchestrator client and fetch mocked hash
//func TestOrcClient_GetConfirmedPanBlockHashesWithHTTP(t *testing.T) {
//	// connect with a remote server and create an orchestrator client
//	orcClient, err := Dial(orchestrator_link)
//	if err != nil {
//		t.Fatalf("error found while dialing orchestrator. error %s", err)
//	}
//
//	var request []*BlockHash
//	request = append(request, &BlockHash{Hash: common.HexToHash(MockedHashInvalid), Slot: 1}, &BlockHash{Hash: common.HexToHash(MockedHashPending), Slot: 2})
//	response, err := orcClient.GetConfirmedPanBlockHashes(context.Background(), request)
//	if err != nil {
//		t.Fatalf("error found while getting confirmed pending block hashes")
//	}
//	t.Log("received block confirmations from orchestrator")
//	for _, hash := range response {
//		t.Logf("received %v", hash)
//	}
//
//}

// TestDial dials real orchestrator to create a client
//func TestDial(t *testing.T) {
//	orcClient, err := Dial(orchestrator_link)
//	if err != nil {
//		t.Fatalf("error found while dialing orchestrator. error %s", err)
//	}
//	t.Logf("orchestrator created %v", orcClient)
//}

// testing mock orchestrator service
type mockOrchestratorService struct{}

// NewMockOrchestratorServer method mock pandora chain apis
func NewMockOrchestratorServer() *rpc.Server {
	server := rpc.NewServer()
	if err := server.RegisterName("orc", new(mockOrchestratorService)); err != nil {
		panic(err)
	}
	return server
}

// ConfirmPanBlockHashes confirms block confirmation
func (OrcClient *mockOrchestratorService) ConfirmPanBlockHashes(ctx context.Context,
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
