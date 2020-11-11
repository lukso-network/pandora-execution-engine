// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package Aura implements the proof-of-authority consensus engine.
package aura

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"io"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	lru "github.com/hashicorp/golang-lru"
)

const (
	checkpointInterval = 1024 // Number of blocks after which to save the vote snapshot to the database
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory

	wiggleTime = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers
)

// Aura proof-of-authority protocol constants.
var (
	epochLength = uint64(30000) // Default number of blocks after which to checkpoint and reset the pending votes

	extraVanity = 32                     // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = crypto.SignatureLength // Fixed number of extra-data suffix bytes reserved for signer seal

	nonceAuthVote = hexutil.MustDecode("0xffffffffffffffff") // Magic nonce number to vote on adding a new signer
	nonceDropVote = hexutil.MustDecode("0x0000000000000000") // Magic nonce number to vote on removing a signer.

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.

	diffInTurn = big.NewInt(2) // Block difficulty for in-turn signatures
	diffNoTurn = big.NewInt(1) // Block difficulty for out-of-turn signatures
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errInvalidCheckpointBeneficiary is returned if a checkpoint/epoch transition
	// block has a beneficiary set to non-zeroes.
	errInvalidCheckpointBeneficiary = errors.New("beneficiary in checkpoint block non-zero")

	// errInvalidVote is returned if a nonce value is something else that the two
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidVote = errors.New("vote nonce not 0x00..0 or 0xff..f")

	// errInvalidCheckpointVote is returned if a checkpoint/epoch transition block
	// has a vote nonce set to non-zeroes.
	errInvalidCheckpointVote = errors.New("vote nonce in checkpoint block non-zero")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte signature suffix missing")

	// errExtraSigners is returned if non-checkpoint block contain signer data in
	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// errInvalidValidatorSeal is returned if the extra data field length is not
	// equal to the length of a seal
	errInvalidExtraData = errors.New("extra data field in block header is invalid")

	// errInvalidCheckpointSigners is returned if a checkpoint block contains an
	// invalid list of signers (i.e. non divisible by 20 bytes).
	errInvalidCheckpointSigners = errors.New("invalid signer list on checkpoint block")

	// errMismatchingCheckpointSigners is returned if a checkpoint block contains a
	// list of signers different than the one the local node calculated.
	errMismatchingCheckpointSigners = errors.New("mismatching signer list on checkpoint block")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errInvalidDifficulty is returned if the difficulty of a block neither 1 or 2.
	errInvalidDifficulty = errors.New("invalid difficulty")

	// errWrongDifficulty is returned if the difficulty of a block doesn't match the
	// turn of the signer.
	errWrongDifficulty = errors.New("wrong difficulty")

	// errInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	errInvalidTimestamp = errors.New("invalid timestamp")

	// errInvalidVotingChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errInvalidVotingChain = errors.New("invalid voting chain")

	// errUnauthorizedSigner is returned if a header is signed by a non-authorized entity.
	errUnauthorizedSigner = errors.New("unauthorized signer")

	// errInvalidSigner is returned if signer will not be able to sign due to validator config
	errInvalidSigner = errors.New("unauthorized signer which is not within validators list")

	// errRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	errRecentlySigned = errors.New("recently signed")

	errMissingValidatorSet = errors.New("No validator set found")
)

// SignerFn hashes and signs the data to be signed by a backing account.
type SignerFn func(signer accounts.Account, mimeType string, message []byte) ([]byte, error)

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache) (common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Seal) > 2 || len(header.Seal[1]) < extraSeal {
		return common.Address{}, errMissingSignature
	}

	currentSignature := header.Seal[1]
	signature := currentSignature[len(currentSignature)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

// Aura is the proof-of-authority consensus engine proposed to support the
// Ethereum testnet following the Ropsten attacks.
type Aura struct {
	config *params.AuraConfig // Consensus engine configuration parameters
	db     ethdb.Database     // Database to store and retrieve snapshot checkpoints


	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

	proposals map[common.Address]bool // Current list of proposals we are pushing

	signer common.Address // Ethereum address of the signing key
	signFn SignerFn       // Signer function to authorize hashes with
	lock   sync.RWMutex   // Protects the signer fields

	// The fields below are for testing only
	fakeDiff 				bool // Skip difficulty verifications

	// For validator set contract
	contract	 			*ValidatorSetContract
	validatorList 			[]common.Address
	blockNumList			[]int
	multiSet				map[uint64]*MultiSet
	isContractActivated		bool
	flag 					bool
}

type MultiSet struct {
	list 			[]common.Address
	contractAddr 	common.Address
	hasContractAddr	bool
}

// New creates a AuthorityRound proof-of-authority consensus engine with the initial
// signers set to the ones provided by the user.
func New(config *params.AuraConfig, db ethdb.Database) *Aura {
	// Set any missing consensus parameters to their defaults
	conf := *config
	if conf.Epoch == 0 {
		conf.Epoch = epochLength
	}
	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	auraEngine := &Aura{
		config:     &conf,
		db:         db,
		recents:    recents,
		signatures: signatures,
		proposals:  make(map[common.Address]bool),
		contract: 	nil,
		multiSet: 	make(map[uint64]*MultiSet),
	}

	// parse authorities sturcture
	auraEngine.parseMulti(&config.Authorities)
	return auraEngine
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (a *Aura) Author(header *types.Header) (common.Address, error) {
	return ecrecover(header, a.signatures)
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (a *Aura) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header, seal bool) error {
	if !a.flag {
		a.InitiateValidatorList(chain.(*core.BlockChain), chain.Config())
		a.flag = true
	}
	return a.verifyHeader(chain, header, nil)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (a *Aura) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))
	go func() {
		for i, header := range headers {
			err := a.verifyHeader(chain, header, headers[:i])

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (a *Aura) verifyHeader(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}
	//number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}

	log.Debug("Header difficulty and config difficulty", "header.Difficulty", header.Difficulty, "Aura.GetDifficulty", chain.Config().Aura.GetDifficulty())

	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return a.verifyCascadingFields(chain, header, parents)
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (a *Aura) verifyCascadingFields(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to its parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}

	if parent.Time > header.Time {
		return errInvalidTimestamp
	}

	// All basic checks passed, verify the seal and return
	return a.verifySeal(chain, header, parents)
}

// snapshot retrieves the authorization snapshot at a given point in time.
func (a *Aura) snapshot(chain consensus.ChainHeaderReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *Snapshot
	)
	for snap == nil {
		// If an in-memory snapshot was found, use that
		if s, ok := a.recents.Get(hash); ok {
			snap = s.(*Snapshot)
			break
		}
		// If an on-disk checkpoint snapshot can be found, use that
		if number%checkpointInterval == 0 {
			if s, err := loadSnapshot(a.config, a.signatures, a.db, hash); err == nil {
				log.Trace("Loaded voting snapshot from disk", "number", number, "hash", hash)
				snap = s
				break
			}
		}
		// If we're at the genesis, snapshot the initial state. Alternatively if we're
		// at a checkpoint block without a parent (light client CHT), or we have piled
		// up more headers than allowed to be reorged (chain reinit from a freezer),
		// consider the checkpoint trusted and snapshot it.
		if number == 0 || (number%a.config.Epoch == 0 && (len(headers) > params.FullImmutabilityThreshold || chain.GetHeaderByNumber(number-1) == nil)) {
			checkpoint := chain.GetHeaderByNumber(number)
			if checkpoint != nil {
				hash := checkpoint.Hash()

				signers := make([]common.Address, (len(checkpoint.Extra)-extraVanity-extraSeal)/common.AddressLength)
				for i := 0; i < len(signers); i++ {
					copy(signers[i][:], checkpoint.Extra[extraVanity+i*common.AddressLength:])
				}
				snap = newSnapshot(a.config, a.signatures, number, hash, signers)
				if err := snap.store(a.db); err != nil {
					return nil, err
				}
				log.Info("Stored checkpoint snapshot to disk", "number", number, "hash", hash)
				break
			}
		}
		// No snapshot for this header, gather the header and move backward
		var header *types.Header
		if len(parents) > 0 {
			// If we have explicit parents, pick from there (enforced)
			header = parents[len(parents)-1]
			if header.Hash() != hash || header.Number.Uint64() != number {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			// No explicit parents (or no more left), reach out to the database
			header = chain.GetHeader(hash, number)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}
		headers = append(headers, header)
		number, hash = number-1, header.ParentHash
	}
	// Previous snapshot found, apply any pending headers on top of it
	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}
	snap, err := snap.apply(headers)
	if err != nil {
		return nil, err
	}
	a.recents.Add(snap.Hash, snap)

	// If we've generated a new checkpoint snapshot, save to disk
	if snap.Number%checkpointInterval == 0 && len(headers) > 0 {
		if err = snap.store(a.db); err != nil {
			return nil, err
		}
		log.Trace("Stored voting snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}
	return snap, err
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (a *Aura) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// VerifySeal implements consensus.Engine, checking whether the signature contained
// in the header satisfies the consensus protocol requirements.
func (a *Aura) VerifySeal(chain consensus.ChainHeaderReader, header *types.Header) error {
	return a.verifySeal(chain, header, nil)
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (a *Aura) verifySeal(chain consensus.ChainHeaderReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, a.signatures)
	if err != nil {
		return err
	}
	// Checking authorization
	ts := header.Time

	step := ts / a.config.Period
	println(header.Number.Uint64())

	turn := step % uint64(len(a.validatorList))

	if signer != a.validatorList[turn] {
		// not authorized to sign
		return errUnauthorizedSigner
	}

	return nil
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (a *Aura) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	// Nonce is not used in aura engine
	header.Nonce = types.BlockNonce{}
	number := header.Number.Uint64()

	// Mix digest is not used, set to empty
	header.MixDigest = common.Hash{}

	// Fetch the parent
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// Set the correct difficulty
	calculateExpectedDifficulty := func(parentStep uint64, step uint64, emptyStepsLen uint64) (diff *big.Int) {
		maxInt := big.NewInt(0)
		maxBig128 := maxInt.Sqrt(math.MaxBig256)
		diff = big.NewInt(int64(parentStep - step + emptyStepsLen))
		diff = diff.Add(maxBig128, diff)
		return
	}

	auraHeader := &types.AuraHeader{}

	if len(header.Seal) < 2 {
		header.Seal = make([][]byte, 2)
		step := uint64(time.Now().Unix()) / a.config.Period
		var stepBytes []byte
		stepBytes = make([]byte, 8)
		binary.LittleEndian.PutUint64(stepBytes, step)
		header.Seal[0] = stepBytes
	}

	err := auraHeader.FromHeader(header)

	if nil != err {
		return err
	}

	auraParentHeader := &types.AuraHeader{}
	err = auraParentHeader.FromHeader(parent)
	header.Difficulty = calculateExpectedDifficulty(auraParentHeader.Step, auraHeader.Step, 0)

	return nil
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given.
func (a *Aura) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// No block rewards in PoA, so the state remains as is and uncles are dropped
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// FinalizeAndAssemble implements consensus.Engine, ensuring no uncles are set,
// nor block rewards given, and returns the final block.
func (a *Aura) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// No block rewards in PoA, so the state remains as is and uncles are dropped
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts, new(trie.Trie)), nil
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (a *Aura) Authorize(signer common.Address, signFn SignerFn) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.signer = signer
	a.signFn = signFn
}

// Function should be used if you want to wait until there is current validator turn
// If validator wont be able to seal anytime, function will return error
// Be careful because it can set up very large delay if periods are so long
func (a *Aura) WaitForNextSealerTurn(fromTime int64) (err error) {
	closestSealTurnStart, _, err := a.CountClosestTurn(fromTime, 0)

	if nil != err {
		return
	}

	delay := closestSealTurnStart - fromTime

	if delay < 0 {
		return
	}

	log.Warn(fmt.Sprintf("waiting: %d seconds for sealing turn, time now: %d", delay, fromTime))
	<-time.After(time.Duration(delay) * time.Second)
	log.Warn("this is time now", "timeNow", time.Now().Unix())
	return
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
// You should use Seal only if current sealer is within its turn, otherwise you will get error
func (a *Aura) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	log.Trace("Starting sealing in Aura engine", "block", block.Hash(), "curValidatorList", a.validatorList, "stateRoot", block.Header().Root)
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if a.config.Period == 0 && len(block.Transactions()) == 0 {
		log.Info("Sealing paused, waiting for transactions")
		return nil
	}
	// Don't hold the signer fields for the entire sealing procedure
	a.lock.RLock()
	signer, signFn := a.signer, a.signFn
	a.lock.RUnlock()

	// check if sealer will be ever able to sign
	timeNow := time.Now().Unix()
	_, _, err := a.CountClosestTurn(timeNow, int64(0))

	if nil != err {
		// not authorized to sign ever
		return err
	}

	// check if in good turn time frame
	allowed, _, _ := a.CheckStep(int64(header.Time), 0)

	if !allowed {
		log.Warn(
			"Could not seal, because timestamp of header is invalid: Header time: %d, time now: %d",
			"headerTime",
			header.Time,
			"timeNow",
			time.Now().Unix(),
			"hash",
			SealHash(header),
		)
		return nil
	}

	// Attach time of future execution, not current time
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeAura, AuraRLP(header))
	if err != nil {
		return err
	}

	go func() {
		select {
		case <-stop:
			return
		default:
			header.Seal = make([][]byte, 2)
			step := uint64(time.Now().Unix()) / a.config.Period
			var stepBytes []byte
			stepBytes = make([]byte, 8)
			binary.LittleEndian.PutUint64(stepBytes, step)
			header.Seal[0] = stepBytes
			header.Seal[1] = sighash
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()

	return nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func (a *Aura) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return chain.Config().Aura.Difficulty
}

// SealHash returns the hash of a block prior to it being sealed.
func (a *Aura) SealHash(header *types.Header) common.Hash {
	return SealHash(header)
}

// Close implements consensus.Engine. It's a noop for Aura as there are no background threads.
func (a *Aura) Close() error {
	return nil
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (a *Aura) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return []rpc.API{{
		Namespace: "aura",
		Version:   "1.0",
		Service:   &API{chain: chain, aura: a},
		Public:    false,
	}}
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header) (hash common.Hash) {
	hasher := new(bytes.Buffer)
	encodeSigHeader(hasher, header)
	signatureHash := crypto.Keccak256(hasher.Bytes())
	var arr [32]byte
	copy(arr[:], signatureHash)
	return arr
}

// AuraRLP returns the rlp bytes which needs to be signed for the proof-of-authority
// sealing. The RLP to sign consists of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
func AuraRLP(header *types.Header) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header)
	return b.Bytes()
}

// CheckStep should assure you that current time frame allows you to seal block based on validator set
// UnixTimeToCheck allows you to deduce time not based on current time which might be handy
// TimeTolerance allows you to in-flight deduce that propagation is likely or unlikely to fail. Provide 0 if strict.
// For example if sealing the block is about 1 sec and period is 5 secs you would like to know if your
// committed work will ever have a chance to be accepted by others
// Allowed returns if possible to seal
// currentTurnTimestamp returns when time frame of current turn starts in unixTime
// nextTurnTimestamp returns when time frame of next turn starts in unixTime
func (a *Aura) CheckStep(unixTimeToCheck int64, timeTolerance int64) (
	allowed bool,
	currentTurnTimestamp int64,
	nextTurnTimestamp int64,
) {
	guardStepByUnixTime := func(unixTime int64) (allowed bool) {
		step := uint64(unixTime) / a.config.Period
		turn := step % uint64(len(a.validatorList))

		return a.signer == a.validatorList [turn]
	}

	countTimeFrameForTurn := func(unixTime int64) (turnStart int64, nextTurn int64) {
		timeGap := unixTime % int64(a.config.Period)
		turnStart = unixTime

		if timeGap > 0 {
			turnStart = unixTime - timeGap
		}

		nextTurn = turnStart + int64(a.config.Period)

		return
	}

	checkForProvidedUnix := guardStepByUnixTime(unixTimeToCheck)
	checkForPromisedInterval := guardStepByUnixTime(unixTimeToCheck + timeTolerance)
	currentTurnTimestamp, nextTurnTimestamp = countTimeFrameForTurn(unixTimeToCheck)
	allowed = checkForProvidedUnix && checkForPromisedInterval

	return
}

// CountClosestTurn provides you information should you wait and if so how long for next turn for current validator
// If err is other than nil, it means that you wont be able to seal within this epoch or ever
func (a *Aura) CountClosestTurn(unixTimeToCheck int64, timeTolerance int64) (
	closestSealTurnStart int64,
	closestSealTurnStop int64,
	err error,
) {
	for _, _ = range a.validatorList {
		allowed, turnTimestamp, nextTurnTimestamp := a.CheckStep(unixTimeToCheck, timeTolerance)

		if allowed {
			closestSealTurnStart = turnTimestamp
			closestSealTurnStop = nextTurnTimestamp
			return
		}

		unixTimeToCheck = nextTurnTimestamp
	}

	err = errInvalidSigner

	return
}

// Encode to bare hash
func encodeSigHeader(w io.Writer, header *types.Header) {
	err := rlp.Encode(w, []interface{}{
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
		header.Extra,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}


// TriggerValidatorMode implements some basic logic like -
// when nextTransition = currentBlockNum + 1, then validator set change trigger initiates
// if lastTransition was contract then, stop the previous contract
// Now, new transition is static then just update the validator list or new transition is new contract
// then, initiate new contract and start running new contract to listen validator list
func (a *Aura) TriggerValidatorMode(chain *core.BlockChain, config *params.ChainConfig) error {
	currentBlockNumber := chain.CurrentBlock().Number()
	log.Debug("Current block needs to less than one by defined block height to change validator list", "current", currentBlockNumber, "curValidatorList", a.validatorList)

	lastTransition, nextTransition := a.transitionBlock(currentBlockNumber)
	log.Debug("getting transition block number", "lastTransition", lastTransition, "nextTransition", nextTransition)

	if nextTransition != 0 && currentBlockNumber.Cmp(big.NewInt(int64(nextTransition - 1))) == 0 {
		log.Debug("Trigger initiates and changing in validator set read mode")
		// stop the event watching loop
		if a.isContractActivated {
			a.contract.StopWatchingEvent()
		}

		if !a.multiSet[nextTransition].hasContractAddr {
			a.validatorList = a.multiSet[nextTransition].list
			a.isContractActivated = false
			log.Info("Singling transition to new validator list from genesis", "newValidatorList", a.validatorList)
			return nil
		}

		contractAddr := a.multiSet[nextTransition].contractAddr
		log.Info("Signaling transition to fresh validator set contract", "contractAddr", contractAddr)
		validatorContract, err := NewValidatorSetWithSimBackend(contractAddr, chain, a.db, config)
		if err != nil {
			log.Error("Failed to initiate contract instance", "validatorSetContract", validatorContract)
			return err
		}

		a.contract = validatorContract
		a.isContractActivated = true
		a.validatorList = a.contract.GetValidators()
		a.contract.currentValidatorList = a.validatorList
		a.contract.successCall = false
		log.Info("Getting validator list at certain block height", "blockNumber", currentBlockNumber, "validatorList", a.validatorList)
		return nil
 	}
 	return nil
}

// InitiateValidatorList initialize the validator list for initiating blockchain
// todo Need to change here
func (a *Aura) InitiateValidatorList(chain *core.BlockChain, config *params.ChainConfig) error {
	curBlockNum := chain.CurrentBlock().Number()
	lastTransition, _ := a.transitionBlock(curBlockNum)

	if !a.multiSet[lastTransition].hasContractAddr {
		a.validatorList = a.multiSet[lastTransition].list
		a.isContractActivated = false
		log.Info("initiate new static validator list from genesis", "newValidatorList", a.validatorList)
		return nil
	}

	contractAddr := a.multiSet[lastTransition].contractAddr
	validatorContract, err := NewValidatorSetWithSimBackend(contractAddr, chain, a.db, config)
	if err != nil {
		log.Error("failed to initiate contract instance", "validatorSetContract", validatorContract)
		return err
	}

	a.contract = validatorContract
	a.isContractActivated = true
	a.validatorList = a.contract.GetValidators()
	a.contract.currentValidatorList = a.validatorList
	a.contract.successCall = true
	log.Info("initiate validator from validator set contract", "validatorList", a.validatorList)
	return nil
}

// CheckChange method will check the validator set changs
func (a *Aura) CheckChange(chain *core.BlockChain, block *types.Block, stateDB *state.StateDB) error {
	if a.isContractActivated {
		if !a.contract.successCall {
			_, err := a.contract.FinalizeChange(block.Header(), stateDB)
			if err != nil {
				log.Error("Getting error from calling finalize change method", "err", err)
				return err
			}
			a.contract.successCall = true
		} else {
			log.Debug("From CheckAndFinalizeChange method..................")
			pendingValidatorList, _ := a.contract.CheckAndFinalizeChange(block.Header(), stateDB)
			if pendingValidatorList == nil {
				log.Debug("no change in validator set contract")
				return nil
			}
			a.validatorList = pendingValidatorList
			a.contract.currentValidatorList = pendingValidatorList
		}
	}

	header := block.GetBlockHeader()
	prevStateRoot := header.Root
	header.Root = stateDB.IntermediateRoot(chain.Config().IsEIP158(block.Header().Number))
	log.Debug("successfully calling finalizeChange method", "prevStateRoot", prevStateRoot, "curStateRoot", header.Root)
	return nil
}

// todo - Need to implement according to recursive fashion
// parseMulti retrieves validator list and makes a decision of using validator set contract
func (a *Aura) parseMulti(set *params.ValidatorSet) error {
	log.Debug("getting validator set", "validatorSet", set)
	hasContractAddrFn := func (staticList []common.Address) bool {
		if len(staticList) == 0 { return true }
		return false
	}
	if set.Multi == nil {
		if len(set.List) == 0 || len(set.Contract.Bytes()) == 0 { return errMissingValidatorSet }
		a.multiSet[0] = &MultiSet{
			list: 				set.List,
			contractAddr: 		set.Contract,
			hasContractAddr: 	hasContractAddrFn(set.List),
		}
		a.blockNumList = append(a.blockNumList, 0)
		return nil
	}

	if len(set.Multi) == 0 { return errMissingValidatorSet }
	for key, value := range set.Multi {
		a.blockNumList = append(a.blockNumList, int(key))
		a.multiSet[key] = &MultiSet{
			list: value.List,
			contractAddr: value.Contract,
			hasContractAddr: hasContractAddrFn(value.List),
		}
	}
	sort.Ints(a.blockNumList)
	log.Debug("parsed validator set", "multiSetLen", len(a.multiSet), "blockList", a.blockNumList)
	return nil
}

// transitionBlock gives transition block number with respect to current block number
// if current block number 5 and in multi, there are 0, 10, 15 block height exist then, this
// func gives result as 0, 5
func (a *Aura) transitionBlock(curBlockNum *big.Int) (uint64, uint64){
	if len(a.blockNumList) == 1 {
		return uint64(0), uint64(0)
	}
	for index, value := range a.blockNumList {
		if curBlockNum.Cmp(big.NewInt(int64(value))) == 0 {
			return uint64(value), uint64(value)
		} else if curBlockNum.Cmp(big.NewInt(int64(value))) < 0 {
			return uint64(a.blockNumList[index - 1]), uint64(value)
		}
	}
	return uint64(a.blockNumList[len(a.blockNumList) - 1]), uint64(0)
}