package services

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/onee-only/mempool-manager/api/models"
	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
)

type exampleChainService interface {
	GetAllBlocks() []*models.ExampleChainBlock
	ValidateBlock(*models.ExampleChainBlock) bool
	AddBlock(*models.ExampleChainBlock)
}

type ExampleChainService struct{}

var ExampleChain exampleChainService = ExampleChainService{}

func (ExampleChainService) GetAllBlocks() []*models.ExampleChainBlock {
	return models.FindAllBlocks()
}

func (ExampleChainService) ValidateBlock(block *models.ExampleChainBlock) bool {

	current := &models.ChainSummary

	if block.PrevHash != current.LatestHash {
		return false
	}
	if HashBlock(block) != block.Hash {
		return false
	}
	return true
}
func (ExampleChainService) AddBlock(block *models.ExampleChainBlock) {
	_, err := db.ExampleChain.InsertOne(context.TODO(), block)
	lib.HandleErr(err)
}

func HashBlock(block *models.ExampleChainBlock) string {
	bytes, err := json.Marshal(block)
	lib.HandleErr(err)

	hash := sha256.Sum256(bytes)
	return fmt.Sprintf("%x", hash)
}
