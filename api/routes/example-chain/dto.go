package examplechain

import "github.com/onee-only/mempool-manager/api/models"

type BlocksResponse struct {
	Count  int                         `json:"count"`
	Blocks []*models.ExampleChainBlock `json:"blocks"`
}

type LatestBlockResponse struct {
	Height     int    `json:"height"`
	LatestHash string `json:"latestHash"`
}

type CreateBlockRequest struct {
	Data       string `binding:"required"`
	PrivateKey string `binding:"required"`
	PublicKey  string `binding:"required"`
	Hash       string `binding:"required"`
	PrevHash   string `binding:"required"`
	Height     int    `binding:"required"`
	Nonce      int    `binding:"required"`
}
