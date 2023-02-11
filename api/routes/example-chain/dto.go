package examplechain

import models "github.com/onee-only/mempool-manager/api/models/example-chain"

type BlocksResponse struct {
	Count  int                         `json:"count"`
	Blocks []*models.ExampleChainBlock `json:"blocks"`
}

type LatestBlockResponse struct {
	Height     int    `json:"height"`
	LatestHash string `json:"latestHash"`
}

type CreateBlockRequest struct {
	PrivateKey string `json:"privateKey" binding:"required"`
	Block      struct {
		Data      string `json:"data" binding:"required"`
		PublicKey string `json:"publicKey" binding:"required"`
		Hash      string `json:"hash" binding:"required"`
		PrevHash  string `json:"prevHash" binding:"required"`
		Height    int    `json:"height" binding:"required"`
		Nonce     int    `json:"nonce" binding:"required"`
	} `json:"block" binding:"required"`
}
