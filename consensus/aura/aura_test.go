package aura

import (
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	lru "github.com/hashicorp/golang-lru"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strings"
	"testing"
	"time"
)

var (
	auraChainConfig *params.AuraConfig
	testBankKey, _  = crypto.GenerateKey()
	testBankAddress = crypto.PubkeyToAddress(testBankKey.PublicKey)
	auraEngine *Aura
)

func init() {
	authority1, _ := crypto.GenerateKey()
	authority2, _ := crypto.GenerateKey()
	auraChainConfig = &params.AuraConfig{
		Period: 5,
		Epoch:  500,
		Authorities: []common.Address{
			testBankAddress,
			crypto.PubkeyToAddress(authority1.PublicKey),
			crypto.PubkeyToAddress(authority2.PublicKey),
		},
		Difficulty: big.NewInt(int64(131072)),
		Signatures: nil,
	}

	db := rawdb.NewMemoryDatabase()
	auraEngine = New(auraChainConfig, db)

	signerFunc := func(account accounts.Account, s string, data []byte) ([]byte, error) {
		return crypto.Sign(crypto.Keccak256(data), testBankKey)
	}
	auraEngine.Authorize(testBankAddress, signerFunc)
}

func TestAura_CheckStep(t *testing.T) {
	currentTime := int64(1602588556)

	t.Run("should return true with no tolerance", func(t *testing.T) {
		allowed, currentTurnTimestamp, nextTurnTimestamp := auraEngine.CheckStep(currentTime, 0)
		assert.True(t, allowed)
		// Period is 5 so next time frame started within -1 from unix time
		assert.Equal(t, currentTime - 1, currentTurnTimestamp)
		// Period is 5 so next time frame starts within 4 secs from unix time
		assert.Equal(t, currentTime + 4, nextTurnTimestamp)
	})

	t.Run("should return true with small tolerance", func(t *testing.T) {
		allowed, currentTurnTimestamp, nextTurnTimestamp := auraEngine.CheckStep(
			currentTime,
			time.Unix(currentTime, 25).Unix(),
		)
		assert.True(t, allowed)
		// Period is 5 so next time frame started within -1 from unix time
		assert.Equal(t, currentTime - 1, currentTurnTimestamp)
		// Period is 5 so next time frame starts within 4 secs from unix time
		assert.Equal(t, currentTime + 4, nextTurnTimestamp)
	})

	t.Run("should return false with no tolerance", func(t *testing.T) {
		timeToCheck := currentTime + int64(6)
		allowed, currentTurnTimestamp, nextTurnTimestamp := auraEngine.CheckStep(timeToCheck, 0)
		assert.False(t, allowed)
		assert.Equal(t, timeToCheck - 2, currentTurnTimestamp)
		assert.Equal(t, timeToCheck + 3, nextTurnTimestamp)
	})

	// If base unixTime is invalid fail no matter what tolerance is
	// If you start sealing before its your turn or you have missed your time frame you should resubmit work
	t.Run("should return false with tolerance", func(t *testing.T) {
		timeToCheck := currentTime + int64(5)
		allowed, currentTurnTimestamp, nextTurnTimestamp := auraEngine.CheckStep(
			timeToCheck,
			time.Unix(currentTime + 80, 0).Unix(),
		)
		assert.False(t, allowed)
		assert.Equal(t, timeToCheck - 1, currentTurnTimestamp)
		assert.Equal(t, timeToCheck + 4, nextTurnTimestamp)
	})
}

func TestAura_CountClosestTurn(t *testing.T) {
	currentTime := int64(1602588556)

	t.Run("should return error, because validator wont be able to seal", func(t *testing.T) {
		randomValidatorKey, err := crypto.GenerateKey()
		assert.Nil(t, err)
		auraChainConfig = &params.AuraConfig{
			Period: 5,
			Epoch:  500,
			Authorities: []common.Address{
				crypto.PubkeyToAddress(randomValidatorKey.PublicKey),
			},
			Difficulty: big.NewInt(int64(131072)),
			Signatures: nil,
		}

		db := rawdb.NewMemoryDatabase()
		modifiedAuraEngine := New(auraChainConfig, db)
		closestSealTurnStart, closestSealTurnStop, err := modifiedAuraEngine.CountClosestTurn(
			time.Now().Unix(),
			0,
		)
		assert.Equal(t, errInvalidSigner, err)
		assert.Equal(t, int64(0), closestSealTurnStart)
		assert.Equal(t, int64(0), closestSealTurnStop)
	})

	t.Run("should return current time frame", func(t *testing.T) {
		closestSealTurnStart, closestSealTurnStop, err := auraEngine.CountClosestTurn(currentTime, 0)
		assert.Nil(t, err)
		assert.Equal(t, currentTime - 1, closestSealTurnStart)
		assert.Equal(t, currentTime + 4, closestSealTurnStop)
	})

	t.Run("should return time frame in future", func(t *testing.T) {
		timeModified := currentTime + 5
		closestSealTurnStart, closestSealTurnStop, err := auraEngine.CountClosestTurn(timeModified, 0)
		assert.Nil(t, err)
		assert.Equal(t, timeModified + 9, closestSealTurnStart)
		assert.Equal(t, timeModified + 14, closestSealTurnStop)
	})
}

func TestAura_DecodeSeal(t *testing.T) {
	// Block 1 rlp data
	msg4Node0 := "f90241f9023ea02778716827366f0a5479d7a907800d183c57382fa7142b84fbb71db143cf788ca01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d493479470ad1a5fba52e27173d23ad87ad97c9bbe249abfa040cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000090ffffffffffffffffffffffffeceb197b0183222aa980845f6880949cdb830300018c4f70656e457468657265756d86312e34332e31826c69841314e684b84179d277eb6b97d25776793c1a98639d8d41da413fba24c338ee83bff533eac3695a0afaec6df1b77a48681a6a995798964adec1bb406c91b6bbe35f115a828a4101"
	input, err := hex.DecodeString(msg4Node0)
	assert.Nil(t, err)

	var auraHeaders []*types.AuraHeader
	err = rlp.Decode(bytes.NewReader(input), &auraHeaders)
	assert.Nil(t, err)
	assert.NotEmpty(t, auraHeaders)

	for _, header := range auraHeaders {
		// excepted block 1 hash (from parity rpc)
		hashExpected := "0x4d286e4f0dbce8d54b27ea70c211bc4b00c8a89ac67f132662c6dc74d9b294e4"
		assert.Equal(t, hashExpected, header.Hash().String())
		stdHeader := header.TranslateIntoHeader()
		stdHeaderHash := stdHeader.Hash()
		assert.Equal(t, hashExpected, stdHeaderHash.String())
		if header.Number.Int64() == int64(1) {
			signatureForSeal := new(bytes.Buffer)
			encodeSigHeader(signatureForSeal, stdHeader)
			messageHashForSeal := SealHash(stdHeader).Bytes()
			hexutil.Encode(crypto.Keccak256(signatureForSeal.Bytes()))
			pubkey, err := crypto.Ecrecover(messageHashForSeal, stdHeader.Seal[1])

			assert.Nil(t, err)
			var signer common.Address
			copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
			// 0x70ad1a5fba52e27173d23ad87ad97c9bbe249abf - Block 1 miner
			assert.Equal(t, "0x70ad1a5fba52e27173d23ad87ad97c9bbe249abf", strings.ToLower(signer.Hex()))
		}
	}
}

func TestAura_Seal(t *testing.T) {
	// block hex comes from worker test and is extracted due to unit-level of testing Seal
	blockToSignHex := "0xf902c5f9025ca0f0513bebf98c814b3c28ff61746552f74ed65909a3ca4cc3ea5b56dc6021ee3ea01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347940000000000000000000000000000000000000000a02c6e36b7f66da996dc550a19d56c9994626304dc77e459963c1b4dde768020cda02457516422f685ff3338d36c41f3eaa26c35b53f4d485d8d93543c1c4b8bdf6ba0056b23fbba480696b65fe5a59b8f2148a1299103c4f57df839233af2cf4ca2d2b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000083020000018347e7c4825208845f84393fb86100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000880000000000000000c0f863f8618080825208943da0ae25cdf7004849e352ba1f8b59ea4b6ebd708203e8801ca00ab99fc4760dfddc35ebd4bf4c4be06e3a2b2d6995fa37b674142c573f7683dda008e2b5c9e9c4597b59d639d7d0aba1b0aa4ddeaf4dceb8b89b914272aa340a1ac0"
	blockBytes, err := hexutil.Decode(blockToSignHex)
	assert.Nil(t, err)
	var block types.Block
	err = rlp.DecodeBytes(blockBytes, &block)
	assert.Nil(t, err)

	// Header should not contain Signature and Step because for now it is not signed
	header := block.Header()
	assert.Empty(t, header.Seal)

	// Wait for next turn to start sealing
	timeout := 3
	timeNow := time.Now().Unix()
	closestSealTurnStart, _, err := auraEngine.CountClosestTurn(timeNow, 0)
	assert.Nil(t, err)

	// Seal the block
	chain := core.BlockChain{}
	resultsChan := make(chan *types.Block)
	stopChan := make(chan struct{})
	waitFor := closestSealTurnStart - timeNow

	if waitFor < 1 {
		waitFor = 0
	}

	t.Logf("Test is waiting for proper turn to start sealing. Waiting: %v secs", waitFor)
	time.Sleep(time.Duration(waitFor) * time.Second)
	err = auraEngine.Seal(&chain, &block, resultsChan, stopChan)

	select {
	case receivedBlock := <-resultsChan:
		assert.Nil(t, err)
		assert.IsType(t, &types.Block{}, receivedBlock)
		header := receivedBlock.Header()
		assert.Len(t, header.Seal, 2)
		signatureForSeal := new(bytes.Buffer)
		encodeSigHeader(signatureForSeal, header)
		messageHashForSeal := SealHash(header).Bytes()
		hexutil.Encode(crypto.Keccak256(signatureForSeal.Bytes()))
		pubkey, err := crypto.Ecrecover(messageHashForSeal, header.Seal[1])

		assert.Nil(t, err)
		var signer common.Address
		copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

		// Signer should be equal sealer
		assert.Equal(t, strings.ToLower(testBankAddress.String()), strings.ToLower(signer.Hex()))
	case <- time.After(time.Duration(timeout) * time.Second):
		t.Fatalf("Received timeout")

	case receivedStop := <-stopChan:
		t.Fatalf("Received stop, but did not expect this, %v", receivedStop)
	}
}

func TestAura_VerifySeal(t *testing.T) {
	// Block 1 rlp data
	msg4Node0 := "f90241f9023ea02778716827366f0a5479d7a907800d183c57382fa7142b84fbb71db143cf788ca01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d493479470ad1a5fba52e27173d23ad87ad97c9bbe249abfa040cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000090ffffffffffffffffffffffffeceb197b0183222aa980845f6880949cdb830300018c4f70656e457468657265756d86312e34332e31826c69841314e684b84179d277eb6b97d25776793c1a98639d8d41da413fba24c338ee83bff533eac3695a0afaec6df1b77a48681a6a995798964adec1bb406c91b6bbe35f115a828a4101"
	input, err := hex.DecodeString(msg4Node0)
	assert.Nil(t, err)
	var auraHeaders []*types.AuraHeader
	err = rlp.Decode(bytes.NewReader(input), &auraHeaders)
	assert.Nil(t, err)
	assert.NotEmpty(t, auraHeaders)
	var aura Aura
	auraConfig := &params.AuraConfig{
		Period: uint64(5),
		Authorities: []common.Address{
			common.HexToAddress("0x70ad1a5fba52e27173d23ad87ad97c9bbe249abf"),
			common.HexToAddress("0xafe443af9d1504de4c2d486356c421c160fdd7b1"),
		},
	}
	aura.config = auraConfig
	var auraSignatures *lru.ARCCache
	auraSignatures, err = lru.NewARC(inmemorySignatures)
	assert.Nil(t, err)
	auraSignatures.Add(0, "0x6f17a2ade9f6daed3968b73514466e07e3c1fef2d6350946e1a12d2b577af0aa")
	aura.signatures = auraSignatures
	for _, header := range auraHeaders {
		// excepted block 1 hash (from parity rpc)
		hashExpected := "0x4d286e4f0dbce8d54b27ea70c211bc4b00c8a89ac67f132662c6dc74d9b294e4"
		assert.Equal(t, hashExpected, header.Hash().String())
		stdHeader := header.TranslateIntoHeader()
		stdHeaderHash := stdHeader.Hash()
		assert.Equal(t, hashExpected, stdHeaderHash.String())
		if header.Number.Int64() == int64(1) {
			signatureForSeal := new(bytes.Buffer)
			encodeSigHeader(signatureForSeal, stdHeader)
			messageHashForSeal := SealHash(stdHeader).Bytes()
			hexutil.Encode(crypto.Keccak256(signatureForSeal.Bytes()))
			pubkey, err := crypto.Ecrecover(messageHashForSeal, stdHeader.Seal[1])
			assert.Nil(t, err)
			err = aura.VerifySeal(nil, stdHeader)
			assert.Nil(t, err)
			var signer common.Address
			copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
			// 0x70ad1a5fba52e27173d23ad87ad97c9bbe249abf - Block 1 miner
			assert.Equal(t, "0x70ad1a5fba52e27173d23ad87ad97c9bbe249abf", strings.ToLower(signer.Hex()))
		}
	}
}
