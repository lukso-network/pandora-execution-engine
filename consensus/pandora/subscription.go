package pandora

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"time"
)

// waitForConnection waits for a connection with pandora chain. Until a successful connection and subscription with
// pandora chain, it retries again and again.
func (p *Pandora) waitForConnection() {
	log.Debug("Waiting for the connection with orchestrator client")
	var err error
	if err = p.connectToOrchestrator(); err == nil {
		log.Info("Connected and subscribed to orchestrator client", "endpoint", p.endpoint)
		p.connected = true
		return
	}
	log.Warn("Could not connect or subscribe to orchestrator client", "err", err)
	p.runError = err
	ticker := time.NewTicker(reConPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Debug("Dialing orchestrator node", "endpoint", p.endpoint)
			var errConnect error
			if errConnect = p.connectToOrchestrator(); errConnect != nil {
				log.Warn("Could not connect or subscribe to orchestrator client", "err", errConnect)
				p.runError = errConnect
				continue
			}
			p.connected = true
			p.runError = nil
			log.Info("Connected and subscribed to orchestrator client", "endpoint", p.endpoint)
			return
		case <-p.ctx.Done():
			log.Debug("Received cancelled context,closing existing waiting connection loop")
			return
		}
	}
}

// connectToChain dials to pandora chain and creates rpcClient and subscribe
func (p *Pandora) connectToOrchestrator() error {
	if p.rpcClient == nil {
		panRPCClient, err := p.dialRPC(p.endpoint)
		if err != nil {
			return err
		}
		p.rpcClient = panRPCClient
	}
	cs, err := p.subscribe()
	if err != nil {
		return err
	}
	p.subscription = cs
	return nil
}

func (p *Pandora) subscribe() (*rpc.ClientSubscription, error) {
	curHeader := p.chain.CurrentHeader()
	log.Debug("Retrieved current header from local chain", "curHeader", fmt.Sprintf("%+v", curHeader))
	extraData := new(ExtraData)
	err := rlp.DecodeBytes(curHeader.Extra, extraData)
	if err != nil {
		return nil, err
	}
	// subscribing from previous epoch
	p.currentEpoch = extraData.Epoch - 1
	if p.currentEpoch < 0 {
		p.currentEpoch = 0
	}
	// connect to pandora subscription
	sub, err := p.SubscribeEpochInfo(p.ctx, p.currentEpoch, p.namespace, p.rpcClient)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

// retryToConnectAndSubscribe retries to pandora chain in case of any failure.
func (p *Pandora) retryToConnectAndSubscribe(err error) {
	p.runError = err
	p.connected = false
	// Back off for a while before resuming dialing the pandora node.
	time.Sleep(reConPeriod)
	p.waitForConnection()
	// Reset run error in the event of a successful connection.
	p.runError = nil
}

// subscribePendingHeaders subscribes to pandora client from latest saved slot using given rpc client
func (p *Pandora) SubscribeEpochInfo(
	ctx context.Context,
	fromEpoch uint64,
	namespace string,
	client *rpc.Client,
) (*rpc.ClientSubscription, error) {
	ch := make(chan *EpochInfo)
	sub, err := client.Subscribe(ctx, namespace, ch, "minimalConsensusInfo", fromEpoch)
	if nil != err {
		return nil, err
	}
	log.Debug("subscribed to orchestrator for new epoch info", "fromEpoch", fromEpoch)

	// Start up a dispatcher to feed into the callback
	go func() {
		for {
			select {
			case epochInfo := <-ch:
				log.Debug("Received new epoch info", "epochInfo", fmt.Sprintf("%+v", epochInfo))
				// dispatch newPendingHeader to handler
				err = p.processEpochInfo(epochInfo)
				if nil != err {
					log.Error("Failed to process epoch info", "err", err)
				}
			case err := <-sub.Err():
				if err != nil {
					log.Debug("Got subscription error", "err", err)
					p.subscriptionErrCh <- err
				}
				return
			}
		}
	}()

	return sub, nil
}

// processEpochInfo
func (p *Pandora) processEpochInfo(info *EpochInfo) error {
	// checking proper length of the validator list.
	if uint64(len(info.ValidatorList)) != p.config.SlotsPerEpoch {
		return errInvalidValidatorSize
	}

	// store epoch info in in-memeory cache
	if err := p.epochInfoCache.put(info.Epoch, info); err != nil {
		return err
	}

	return nil
}
