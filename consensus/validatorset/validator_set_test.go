package validatorset

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/params"
	"io/ioutil"
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
	NewValidatorSet(multiMap, &authority, nil, nil, nil)
	//multi := authority.Multi[20]
	//contractAddr := multi.Contract
	//assert.Equal(t, contractAddr, common.HexToAddress("0xc6d9d2cd449a754c494264e1809c50e34d64562b"))
}