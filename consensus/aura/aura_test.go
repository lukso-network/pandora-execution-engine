package aura

import (
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAura_Seal(t *testing.T) {
	msg4Node0 := "f90241f9023ea02778716827366f0a5479d7a907800d183c57382fa7142b84fbb71db143cf788ca01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d493479470ad1a5fba52e27173d23ad87ad97c9bbe249abfa040cf4430ecaa733787d1a65154a3b9efb560c95d9e324a23b97f0609b539133ba056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421a056e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421b901000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000090ffffffffffffffffffffffffeceb197b0183222aa980845f6880949cdb830300018c4f70656e457468657265756d86312e34332e31826c69841314e684b84179d277eb6b97d25776793c1a98639d8d41da413fba24c338ee83bff533eac3695a0afaec6df1b77a48681a6a995798964adec1bb406c91b6bbe35f115a828a4101"
	input, err := hex.DecodeString(msg4Node0)
	assert.Nil(t, err)

	var auraHeaders []*types.AuraHeader
	err = rlp.Decode(bytes.NewReader(input), &auraHeaders)
	assert.Nil(t, err)
	assert.NotEmpty(t, auraHeaders)

	for _, header := range auraHeaders {
		hashExpected := "0x4d286e4f0dbce8d54b27ea70c211bc4b00c8a89ac67f132662c6dc74d9b294e4"
		assert.Equal(t, hashExpected, header.Hash().String())
		stdHeader := header.TranslateIntoHeader()
		stdHeaderHash := stdHeader.Hash()
		assert.Equal(t, hashExpected, stdHeaderHash.String())
		if header.Number.Int64() == int64(1) {

			signatureForSeal := new(bytes.Buffer)
			encodeSigHeader(signatureForSeal, stdHeader)
			println(SealHash(stdHeader).String())
			println("\n\n")
			messageHashForSeal := SealHash(stdHeader).Bytes()
			hexutil.Encode(crypto.Keccak256(signatureForSeal.Bytes()))
			pubkey, err := crypto.Ecrecover(messageHashForSeal, stdHeader.Seal[1])

			assert.Nil(t, err)
			var signer common.Address

			copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
			println(signer.Hex())
			assert.Equal(t, "0x70ad1a5fba52e27173d23ad87ad97c9bbe249abf", strings.ToLower(signer.Hex()))

		}
	}
}
