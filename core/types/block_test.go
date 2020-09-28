// Copyright 2014 The go-ethereum Authors
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

package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"hash"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

// from bcValidBlockTest.json, "SimpleTx"
func TestBlockEncoding(t *testing.T) {
	blockEnc := common.FromHex("f90260f901f9a083cafc574e1f51ba9dc0568fc617a08ea2429fb384059c972f13b19fa1c8dd55a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347948888f1f195afa192cfee860698584c030f4c9db1a0ef1552a40b7165c3cd773806b9e0c165b75356e0314bf0706f279c729f51e017a05fe50b260da6308036625b850b5d6ced6d0a9f814c0688bc91ffb7b7a3a54b67a0bc37d79753ad738a6dac4921e57392f145d8887476de3f783dfa7edae9283e52b90100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008302000001832fefd8825208845506eb0780a0bd4472abb6659ebe3ee06ee4d7b72a00a9f4d001caca51342001075469aff49888a13a5a8c8f2bb1c4f861f85f800a82c35094095e7baea6a6c7c4c2dfeb977efac326af552d870a801ba09bea4c4daac7c7c52e093e6a4c35dbbcf8856f1af7b059ba20253e70848d094fa08a8fae537ce25ed8cb5af9adac3f141af69bd515bd2ba031522df09b97dd72b1c0")
	var block Block
	if err := rlp.DecodeBytes(blockEnc, &block); err != nil {
		t.Fatal("decode error: ", err)
	}

	check := func(f string, got, want interface{}) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s mismatch: got %v, want %v", f, got, want)
		}
	}
	check("Difficulty", block.Difficulty(), big.NewInt(131072))
	check("GasLimit", block.GasLimit(), uint64(3141592))
	check("GasUsed", block.GasUsed(), uint64(21000))
	check("Coinbase", block.Coinbase(), common.HexToAddress("8888f1f195afa192cfee860698584c030f4c9db1"))
	check("MixDigest", block.MixDigest(), common.HexToHash("bd4472abb6659ebe3ee06ee4d7b72a00a9f4d001caca51342001075469aff498"))
	check("Root", block.Root(), common.HexToHash("ef1552a40b7165c3cd773806b9e0c165b75356e0314bf0706f279c729f51e017"))
	check("Hash", block.Hash(), common.HexToHash("0a5843ac1cb04865017cb35a57b50b07084e5fcee39b5acadade33149f4fff9e"))
	check("Nonce", block.Nonce(), uint64(0xa13a5a8c8f2bb1c4))
	check("Time", block.Time(), uint64(1426516743))
	check("Size", block.Size(), common.StorageSize(len(blockEnc)))

	tx1 := NewTransaction(0, common.HexToAddress("095e7baea6a6c7c4c2dfeb977efac326af552d87"), big.NewInt(10), 50000, big.NewInt(10), nil)
	tx1, _ = tx1.WithSignature(HomesteadSigner{}, common.Hex2Bytes("9bea4c4daac7c7c52e093e6a4c35dbbcf8856f1af7b059ba20253e70848d094f8a8fae537ce25ed8cb5af9adac3f141af69bd515bd2ba031522df09b97dd72b100"))
	check("len(Transactions)", len(block.Transactions()), 1)
	check("Transactions[0].Hash", block.Transactions()[0].Hash(), tx1.Hash())

	ourBlockEnc, err := rlp.EncodeToBytes(&block)
	if err != nil {
		t.Fatal("encode error: ", err)
	}
	if !bytes.Equal(ourBlockEnc, blockEnc) {
		t.Errorf("encoded block mismatch:\ngot:  %x\nwant: %x", ourBlockEnc, blockEnc)
	}
}

func TestBlockEncodingAuraHeader(t *testing.T) {
	msg7FromNode0 := "f9023ea02778716827366f0a5479d7a907800d183c57382fa7142b84fbb71db143cf788ca01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d493479470ad1a5fba52e27173d23ad87ad97c9bbe249abfa040cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000090ffffffffffffffffffffffffeceb197b0183222aa980845f6880949cdb830300018c4f70656e457468657265756d86312e34332e31826c69841314e684b84179d277eb6b97d25776793c1a98639d8d41da413fba24c338ee83bff533eac3695a0afaec6df1b77a48681a6a995798964adec1bb406c91b6bbe35f115a828a4101"
	blockEncAuraHeader := common.FromHex(msg7FromNode0)
	var auraHeader AuraHeader
	err := rlp.DecodeBytes(blockEncAuraHeader, &auraHeader)
	assert.Nil(t, err)
}

func TestBlockEncodingAura(t *testing.T) {
	t.Run("Should encode block from json", func(t *testing.T) {
		getLatestBlockResponse := `
		{
			"author": "0x70ad1a5fba52e27173d23ad87ad97c9bbe249abf",
			"difficulty":"0x400",
			"extraData": "0xdb830300018c4f70656e457468657265756d86312e34332e31826c69",
			"gasLimit": "0x8000000",
			"gasUsed": "0x0",
			"hash": "0x3b9384a861bba3bea8f28161f6fabbca4a6215cead9ce7c363536ffd70ffb63f",
			"logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			"miner": "0x70ad1a5fba52e27173d23ad87ad97c9bbe249abf",
			"number": "0xCbCd",
			"parentHash": "0xe076375bfb9bb5eceacbace9562a5b07e299dfc9455b24f9f5a8e08fa703e944",
			"receiptsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
			"sealFields": ["0x8413167e1c", "0xb8414443597ee9435882330f0edfa604a7dfe896a6ab47becb5e8289e995285a170f035a3dae13a510c73be3430ce1ec116138f7fe65e8017e635929dbc32a58853901"],
			"sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
			"signature": "4443597ee9435882330f0edfa604a7dfe896a6ab47becb5e8289e995285a170f035a3dae13a510c73be3430ce1ec116138f7fe65e8017e635929dbc32a58853901",
			"size": 584,
			"stateRoot": "0x40cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133b",
			"step": 320241180,
			"timestamp": "0x5F70768C",
			"totalDifficulty": 1.7753551929366122454274643393537642576131607e+43,
			"transactions": [],
			"transactionsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
			"uncles": []
		}`

		var auraHeader AuraHeader
		jsonBytes := []byte(getLatestBlockResponse)
		err := json.Unmarshal(jsonBytes, &auraHeader)
		assert.Nil(t, err)

		stdHeader := auraHeader.TranslateIntoHeader()
		assert.Nil(t, err)
		assert.IsType(t, &Header{}, stdHeader)

		var buf bytes.Buffer
		block := NewBlock(stdHeader, nil, nil, nil, nil)
		err = block.EncodeRLP(&buf)
		assert.Nil(t, err)

		var auraHeaderRlpBytes bytes.Buffer
		err = rlp.Encode(&auraHeaderRlpBytes, &auraHeader)
		assert.Nil(t, err)

		t.Run("Rlp decode into standard block", func(t *testing.T) {
			var stdBlock Block
			err = rlp.Decode(&buf, &stdBlock)
			assert.Nil(t, err)
			header := stdBlock.Header()
			assert.NotNil(t, header)
			assert.NotEmpty(t, header.Seal)
		})

		t.Run("Rlp decode into aura header", func(t *testing.T) {
			var decodedAuraHeader AuraHeader
			err = rlp.Decode(&auraHeaderRlpBytes, &decodedAuraHeader)
			assert.Nil(t, err)
		})
	})

	t.Run("Message 0x07 - incoming block", func(t *testing.T) {
		msg7FromNode0 := "f9025bf90245f90240a09041480ab2f8b1217f2278e000fd5198d58e8c31b6180946da6bbcc20b516055a01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d493479470ad1a5fba52e27173d23ad87ad97c9bbe249abfa040cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000090fffffffffffffffffffffffffffffffd828d5b837a120080845f6e06189cdb830300018c4f70656e457468657265756d86312e34332e31826c698413160138b841e9669a4e282d5e6fd2e09ae6f4c7253cead13da53456653eec3212e70c61d2be54f370334c684b102d39efeb87543fb84f58f86bdad5f10d0f6e357ee9d093ad01c0c0928d5affffffffffffffffffffffffeceb716d"

		type mappedAura struct {
			Block *AuraBlock
			TD *big.Int
		}

		var mappedAuraResp mappedAura
		input, err := hex.DecodeString(msg7FromNode0)
		assert.Nil(t, err)
		err = rlp.Decode(bytes.NewReader(input), &mappedAuraResp)
		assert.Nil(t, err)

		auraBlock := mappedAuraResp.Block
		assert.NotNil(t, auraBlock)
		err, stdBlock := auraBlock.TranslateIntoBlock()
		assert.Nil(t, err)
		assert.IsType(t, &Block{}, stdBlock)

		t.Run("Block should be valid", func(t *testing.T) {
			stdBlockHash := stdBlock.Hash()
			assert.NotNil(t, auraBlock.Header)
			stdHeader := auraBlock.Header.TranslateIntoHeader()
			stdHeaderHash := stdHeader.Hash()
			assert.Equal(t, stdBlock.header, stdHeader)
			assert.Equal(t, stdHeaderHash, stdBlockHash)
		})
	})

	t.Run("Message 0x04 - incoming batch of headers", func(t *testing.T) {
		msg4Node0 := "f90243f90240a00ca8498075429689026161c395e4238fd4ba4bc61f10b8985f8bad53207472cca01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d493479470ad1a5fba52e27173d23ad87ad97c9bbe249abfa040cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000090fffffffffffffffffffffffffffffffd82d29c837a120080845f70baa29cdb830300018c4f70656e457468657265756d86312e34332e31826c698413168bbab84196d75288c30ee8e025d3e7058511fb887d4830790169a14d1fdfbf53d266473a6a5ef632cd7e7d725cb3db2c2ce8ce253d599bf44a09faabcc9a32905f21502300"
		input, err := hex.DecodeString(msg4Node0)
		assert.Nil(t, err)

		var auraHeaders []*AuraHeader
		err = rlp.Decode(bytes.NewReader(input), &auraHeaders)
		assert.Nil(t, err)
		assert.NotEmpty(t, auraHeaders)
	})

	t.Run("Message 0x04 - incoming batch of headers with block 1 included", func(t *testing.T) {
		msg4Node0 := "f90241f9023ea02778716827366f0a5479d7a907800d183c57382fa7142b84fbb71db143cf788ca01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d493479470ad1a5fba52e27173d23ad87ad97c9bbe249abfa040cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000090ffffffffffffffffffffffffeceb197b0183222aa980845f6880949cdb830300018c4f70656e457468657265756d86312e34332e31826c69841314e684b84179d277eb6b97d25776793c1a98639d8d41da413fba24c338ee83bff533eac3695a0afaec6df1b77a48681a6a995798964adec1bb406c91b6bbe35f115a828a4101"
		input, err := hex.DecodeString(msg4Node0)
		assert.Nil(t, err)

		var auraHeaders []*AuraHeader
		err = rlp.Decode(bytes.NewReader(input), &auraHeaders)
		assert.Nil(t, err)
		assert.NotEmpty(t, auraHeaders)

		for _, header := range auraHeaders {
			hashExpected := "0x4d286e4f0dbce8d54b27ea70c211bc4b00c8a89ac67f132662c6dc74d9b294e4"
			assert.Equal(t, hashExpected, header.Hash().String())
			stdHeader := header.TranslateIntoHeader()
			stdHeaderHash := stdHeader.Hash()
			assert.Equal(t, hashExpected, stdHeaderHash.String())
		}
	})
}

func TestUncleHash(t *testing.T) {
	uncles := make([]*Header, 0)
	h := CalcUncleHash(uncles)
	exp := common.HexToHash("1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")
	if h != exp {
		t.Fatalf("empty uncle hash is wrong, got %x != %x", h, exp)
	}
}

var benchBuffer = bytes.NewBuffer(make([]byte, 0, 32000))

func BenchmarkEncodeBlock(b *testing.B) {
	block := makeBenchBlock()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchBuffer.Reset()
		if err := rlp.Encode(benchBuffer, block); err != nil {
			b.Fatal(err)
		}
	}
}

// testHasher is the helper tool for transaction/receipt list hashing.
// The original hasher is trie, in order to get rid of import cycle,
// use the testing hasher instead.
type testHasher struct {
	hasher hash.Hash
}

func newHasher() *testHasher {
	return &testHasher{hasher: sha3.NewLegacyKeccak256()}
}

func (h *testHasher) Reset() {
	h.hasher.Reset()
}

func (h *testHasher) Update(key, val []byte) {
	h.hasher.Write(key)
	h.hasher.Write(val)
}

func (h *testHasher) Hash() common.Hash {
	return common.BytesToHash(h.hasher.Sum(nil))
}

func makeBenchBlock() *Block {
	var (
		key, _   = crypto.GenerateKey()
		txs      = make([]*Transaction, 70)
		receipts = make([]*Receipt, len(txs))
		signer   = NewEIP155Signer(params.TestChainConfig.ChainID)
		uncles   = make([]*Header, 3)
	)
	header := &Header{
		Difficulty: math.BigPow(11, 11),
		Number:     math.BigPow(2, 9),
		GasLimit:   12345678,
		GasUsed:    1476322,
		Time:       9876543,
		Extra:      []byte("coolest block on chain"),
	}
	for i := range txs {
		amount := math.BigPow(2, int64(i))
		price := big.NewInt(300000)
		data := make([]byte, 100)
		tx := NewTransaction(uint64(i), common.Address{}, amount, 123457, price, data)
		signedTx, err := SignTx(tx, signer, key)
		if err != nil {
			panic(err)
		}
		txs[i] = signedTx
		receipts[i] = NewReceipt(make([]byte, 32), false, tx.Gas())
	}
	for i := range uncles {
		uncles[i] = &Header{
			Difficulty: math.BigPow(11, 11),
			Number:     math.BigPow(2, 9),
			GasLimit:   12345678,
			GasUsed:    1476322,
			Time:       9876543,
			Extra:      []byte("benchmark uncle"),
		}
	}
	return NewBlock(header, txs, uncles, receipts, newHasher())
}
