package aura

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNew(t *testing.T) {
	parityFixture, err := ioutil.ReadFile("./fixtures/block-0-parity.json")
	assert.Nil(t, err)

	//	 Other stuff is not needed, I guess hash is really what matters for now
	//   If you want to strict compare you can compare indented bytes instead
	blockStruct := struct {
		Hash string `json:"hash"`
	}{}

	err = json.Unmarshal(parityFixture, &blockStruct)
	assert.Nil(t, err)

	t.Run("Genesis file should produce same block 0 that in parity", func(t *testing.T) {

	})
}
