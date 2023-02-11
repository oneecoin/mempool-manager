package examplechain_service

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	models "github.com/onee-only/mempool-manager/api/models/example-chain"
	"github.com/onee-only/mempool-manager/lib"
)

func hashBlock(block models.ExampleChainBlock) string {
	block.Hash = ""
	bytes, err := json.Marshal(block)
	lib.HandleErr(err)

	hash := sha256.Sum256(bytes)
	return fmt.Sprintf("%x", hash)
}
