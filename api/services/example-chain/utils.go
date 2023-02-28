package examplechain_service

import (
	"crypto/sha256"
	"fmt"

	models "github.com/onee-only/mempool-manager/api/models/example-chain"
	"github.com/onee-only/mempool-manager/lib"
)

type BlockForHash struct {
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	Height    int    `json:"height"`
	Nonce     int    `json:"nonce"`
	PrevHash  string `json:"prevHash"`
	PublicKey string `json:"publicKey"`
}

func hashBlock(payload *models.ExampleChainBlock) string {
	block := BlockForHash{
		Data:      payload.Data,
		PublicKey: payload.PublicKey,
		Hash:      "",
		PrevHash:  payload.PrevHash,
		Height:    payload.Height,
		Nonce:     payload.Nonce,
	}
	bytes := lib.ToJSON(block)

	hash := sha256.Sum256(bytes)
	return fmt.Sprintf("%x", hash)
}
