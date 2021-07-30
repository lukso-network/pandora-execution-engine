package pandora

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (pan *Pandora) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}
