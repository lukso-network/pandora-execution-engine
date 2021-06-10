/*
* Developed by: Md. Muhaimin Shah Pahalovi
* Generated: 5/11/21
* This file is generated to support Lukso pandora module.
* Purpose:
 */
package filters

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/pandora_orcclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
)

type pandoraTestBackend struct {
	testBackend
	bc *core.BlockChain
}

// GetPendingHeadsSince implements testBackend only for dummy purpose. So that existing code can run without an issue
func (b *testBackend) GetPendingHeadsSince(ctx context.Context, from common.Hash) []*types.Header {
	return nil
}

// SubscribePendingHeaderEvent implements testBackend only for dummy purpose. So that existing code can run without an issue
func (b *testBackend) SubscribePendingHeaderEvent(ch chan<- core.PendingHeaderEvent) event.Subscription {
	return b.pendingHeaderFeed.Subscribe(ch)
}

// GetPendingHeadsSince returns pending headers from blockchain container
func (b *pandoraTestBackend) GetPendingHeadsSince(ctx context.Context, from common.Hash) []*types.Header {
	return b.bc.GetTempHeadersSince(from)
}

// SubscribePendingHeaderEvent subscribe with the headers to get new headers.
func (b *pandoraTestBackend) SubscribePendingHeaderEvent(ch chan<- core.PendingHeaderEvent) event.Subscription {
	return b.bc.SubscribePendingHeaderEvent(ch)
}

// TestPendingHeaderSubscription tests pending header events. In pending header, a batch of headers are come in insert header.
// pending header event send that batch to the API end.
func TestPendingHeaderSubscription(t *testing.T) {
	t.Parallel()

	// Initialize the backend
	var (
		db                  = rawdb.NewMemoryDatabase()
		backend             = &testBackend{db: db}
		api                 = NewPublicFilterAPI(backend, false, deadline)
		genesis             = new(core.Genesis).MustCommit(db)
		chain, _            = core.GenerateChain(params.TestChainConfig, genesis, ethash.NewFaker(), db, 10, func(i int, gen *core.BlockGen) {})
		pendingHeaderEvents = []core.PendingHeaderEvent{}
	)

	var headers []*types.Header

	// form the header chain from the created blocks
	for _, blk := range chain {
		headers = append(headers, blk.Header())
	}
	pendingHeaderEvents = append(pendingHeaderEvents, core.PendingHeaderEvent{Headers: headers})

	// create two subscriber channels
	chan0 := make(chan *types.Header)
	sub0 := api.events.SubscribePendingHeads(chan0)
	chan1 := make(chan *types.Header)
	sub1 := api.events.SubscribePendingHeads(chan1)

	go func() { // simulate client
		sub0Iterator, sub1Iterator := 0, 0
		// a batch of headers are received as event.
		// but in subscriber end we have to send the batch as one by one header
		for sub0Iterator != len(pendingHeaderEvents[0].Headers) || sub1Iterator != len(pendingHeaderEvents[0].Headers) {
			select {
			// here we will receive a single header from the batch and will process it
			case header := <-chan0:
				if pendingHeaderEvents[0].Headers[sub0Iterator].Hash() != header.Hash() {
					t.Errorf("sub0 received invalid hash on index %d, want %x, got %x", sub0Iterator, pendingHeaderEvents[0].Headers[sub0Iterator].Hash(), header.Hash())
				}
				sub0Iterator++
			case header := <-chan1:
				if pendingHeaderEvents[0].Headers[sub1Iterator].Hash() != header.Hash() {
					t.Errorf("sub1 received invalid hash on index %d, want %x, got %x", sub1Iterator, pendingHeaderEvents[0].Headers[sub1Iterator].Hash(), header.Hash())
				}
				sub1Iterator++
			}
		}

		sub0.Unsubscribe()
		sub1.Unsubscribe()
	}()

	time.Sleep(1 * time.Second)
	for _, e := range pendingHeaderEvents {
		// send the pending header batch to the feed.
		backend.pendingHeaderFeed.Send(e)
	}

	<-sub0.Err()
	<-sub1.Err()
}

// makeBlockChain creates a deterministic chain of blocks rooted at parent.
func makeBlockChain(parent *types.Block, n int, engine consensus.Engine, db ethdb.Database, seed int) []*types.Block {
	blocks, _ := core.GenerateChain(params.TestChainConfig, parent, engine, db, n, func(i int, b *core.BlockGen) {
		b.SetCoinbase(common.Address{0: byte(seed), 19: byte(i)})
	})
	return blocks
}

// TestPendingBlockHeaderFullPath tests backend to API subscription level testing of pandora pending event subscription container.
// The testing procedure is discussed here:
// 1. create block chain and a blockchain backend
// 2. Two clients subscribe with the pending header container
// 3. After inserting chain blocks an event is triggered.
// 4. If two clients can get similar headers then test success.
// 5. Add another client and sync it with the pending header container
// 6. Run the same test for the new client.
func TestPendingBlockHeaderFullPath(t *testing.T) {
	t.Parallel()

	// Initialize the backend
	var (
		db                  = rawdb.NewMemoryDatabase()
		genesis             = new(core.Genesis).MustCommit(db)
		blockchain, _       = core.NewBlockChain(db, &core.CacheConfig{OrcClientEndpoint: pandora_orcclient.DialInProcRPCClient()}, params.AllEthashProtocolChanges, ethash.NewFaker(), vm.Config{}, nil, nil)
		backend             = &pandoraTestBackend{bc: blockchain}
		pendingHeaderEvents = []core.PendingHeaderEvent{}
		chain               = makeBlockChain(genesis, 10, ethash.NewFaker(), db, 1)
	)

	var headers []*types.Header

	for _, block := range chain {
		headers = append(headers, block.Header())
	}

	backend.db = db
	var api = NewPublicFilterAPI(backend, false, deadline)

	pendingHeaderEvents = append(pendingHeaderEvents, core.PendingHeaderEvent{Headers: headers})

	// create two subscriber channels
	chan0 := make(chan *types.Header)
	sub0 := api.events.SubscribePendingHeads(chan0)
	chan1 := make(chan *types.Header)
	sub1 := api.events.SubscribePendingHeads(chan1)

	go func() { // simulate client
		t.Logf("exiting go routine %d", len(pendingHeaderEvents[0].Headers))

		sub0Iterator, sub1Iterator := 0, 0
		// a batch of headers are received as event.
		// but in subscriber end we have to send the batch as one by one header
		for sub0Iterator != len(pendingHeaderEvents[0].Headers) || sub1Iterator != len(pendingHeaderEvents[0].Headers) {
			select {
			// here we will receive a single header from the batch and will process it
			case header := <-chan0:
				if pendingHeaderEvents[0].Headers[sub0Iterator].Hash() != header.Hash() {
					t.Errorf("sub0 received invalid hash on index %d, want %x, got %x", sub0Iterator, pendingHeaderEvents[0].Headers[sub0Iterator].Hash(), header.Hash())
				}
				sub0Iterator++
			case header := <-chan1:
				if pendingHeaderEvents[0].Headers[sub1Iterator].Hash() != header.Hash() {
					t.Errorf("sub1 received invalid hash on index %d, want %x, got %x", sub1Iterator, pendingHeaderEvents[0].Headers[sub1Iterator].Hash(), header.Hash())
				}
				sub1Iterator++
			}
		}

		sub0.Unsubscribe()
		sub1.Unsubscribe()
	}()

	// This waiting is only for dummy purpose. We are pretending that after clients are prepared we are inserting headers.
	// We can omit this sleep.
	time.Sleep(1 * time.Second)
	t.Logf("before inserting chain")
	if _, err := blockchain.InsertChain(chain); err != nil {
		t.Fatalf("insert block chain failed due to %v", err)
	}
	t.Logf("after inserting chain")

	<-sub0.Err()
	<-sub1.Err()

	// Give a few seconds to spin up another client
	time.Sleep(5 * time.Second)
	chain = makeBlockChain(chain[len(chain)-1], 10, ethash.NewFaker(), db, 1)

	headers = []*types.Header{}
	for _, block := range chain {
		headers = append(headers, block.Header())
	}

	pendingHeaderEvents = append(pendingHeaderEvents, core.PendingHeaderEvent{Headers: headers})

	chan2 := make(chan *types.Header)
	sub2 := api.events.SubscribePendingHeads(chan2)

	go func() { // simulate client

		// for new orchestrator it will first take all the headers that is not received yet
		pendingHeadersSince := api.backend.GetPendingHeadsSince(nil, pendingHeaderEvents[0].Headers[0].Hash())

		headerIndex := 0
		for _, pendingHeaderEvent := range pendingHeaderEvents {
			for index := 0; headerIndex < len(pendingHeadersSince) && index < len(pendingHeaderEvent.Headers); index++ {
				if pendingHeadersSince[headerIndex].Hash() != pendingHeaderEvent.Headers[index].Hash() {
					t.Errorf("pending event container returns invalid blocks. received %v and expected %v", pendingHeadersSince[headerIndex].Hash(), pendingHeaderEvent.Headers[index].Hash())
				}
				headerIndex++
			}
		}
		sub2Iterator := 0
		// a batch of headers are received as event.
		// but in subscriber end we have to send the batch as one by one header
		for sub2Iterator != len(pendingHeaderEvents[1].Headers) {
			// here we will receive a single header from the batch and will process it
			header := <-chan2
			if pendingHeaderEvents[1].Headers[sub2Iterator].Hash() != header.Hash() {
				t.Errorf("sub2 received invalid hash on index %d, want %x, got %x", sub2Iterator, pendingHeaderEvents[1].Headers[sub2Iterator].Hash(), header.Hash())
			}
			sub2Iterator++
		}

		sub2.Unsubscribe()
	}()

	if _, err := blockchain.InsertChain(chain); err != nil {
		t.Fatalf("found error while inserting blocks into blockhain %v", err)
	}
	<-sub2.Err()
}
