package services

import "github.com/onee-only/mempool-manager/api/models"

type exampleChainService interface {
	GetAllBlocks() []*models.ExampleChainBlock
}

type ExampleChainService struct{}

var ExampleChain exampleChainService = ExampleChainService{}

func (ExampleChainService) GetAllBlocks() []*models.ExampleChainBlock {
	return models.FindAllBlocks()
}
