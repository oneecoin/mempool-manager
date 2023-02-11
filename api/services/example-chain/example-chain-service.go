package examplechain_service

import (
	models "github.com/onee-only/mempool-manager/api/models/example-chain"
)

type IExampleChainService interface {
	GetAllBlocks() []*models.ExampleChainBlock
	ValidateBlock(*models.ExampleChainBlock) bool
	AddBlock(*models.ExampleChainBlock)
}

type ExampleChainService struct{}

var mExchain models.IExchainModel = models.ExchainModel
var ExampleChain IExampleChainService = ExampleChainService{}

func (ExampleChainService) GetAllBlocks() []*models.ExampleChainBlock {
	return mExchain.FindAllBlocks()
}

func (ExampleChainService) ValidateBlock(block *models.ExampleChainBlock) bool {

	if block.PrevHash != mExchain.GetSummary().LatestHash {
		return false
	}
	if hashBlock(*block) != block.Hash {
		return false
	}
	if mExchain.ExistsByPublicKey(block.PublicKey) {
		return false
	}
	return true
}

func (ExampleChainService) AddBlock(block *models.ExampleChainBlock) {
	mExchain.AddBlock(block)

	current := mExchain.GetSummary()
	current.Height = block.Height
	current.LatestHash = block.Hash
}
