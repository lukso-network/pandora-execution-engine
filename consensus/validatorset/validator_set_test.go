package validatorset

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/big"
	"testing"
)

func TestNewValidatorSet(t *testing.T)  {
	validatorSetJSON, err := ioutil.ReadFile("res/validatorset.json")
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	var authority params.ValidatorSet
	err = json.Unmarshal(validatorSetJSON, &authority)
	if err != nil {
		t.Errorf("could not get code at test addr: %v", err)
	}

	multiMap := make(map[uint64]ValidatorSet)
	validators := NewValidatorSet(multiMap, &authority)
	validatorSet := validators.GetValidatorsByCaller(big.NewInt(11))

	assert.Equal(t, validatorSet[0], common.HexToAddress("0xd6d9d2cd449a754c494264e1809c50e34d64562b"))
}