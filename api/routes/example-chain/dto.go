package examplechain

import "github.com/onee-only/mempool-manager/api/models"

type BlocksResponse struct {
	Count  int
	Blocks []*models.ExampleChainBlock
}

type LatestBlockResponse struct {
	Height     int
	LatestHash string
}
