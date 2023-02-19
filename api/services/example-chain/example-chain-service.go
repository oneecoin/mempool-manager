package examplechain_service

import (
	"errors"
	"strings"

	models "github.com/onee-only/mempool-manager/api/models/example-chain"
	transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"
)

const (
	defaultDifficulty = 2
)

type IExampleChainService interface {
	GetAllBlocks() []*models.ExampleChainBlock
	ValidateBlock(*models.ExampleChainBlock) bool
	AddBlock(*models.ExampleChainBlock) error
}

type ExampleChainService struct{}

var mExchain models.IExchainModel = models.ExchainModel
var mTx transaction_model.ITxModel = transaction_model.TxModel
var ExampleChain IExampleChainService = ExampleChainService{}

func (ExampleChainService) GetAllBlocks() []*models.ExampleChainBlock {
	return mExchain.FindAllBlocks()
}

func (ExampleChainService) ValidateBlock(block *models.ExampleChainBlock) bool {

	// haven't added with this public key before
	if !mExchain.ExistsByPublicKey(block.PublicKey) {
		// not up-to-date
		if block.Height != mExchain.GetSummary().Height+1 {
			return false
		}
		// invalid prevHash
		if block.PrevHash != mExchain.GetSummary().LatestHash {
			return false
		}
		// invalid hash
		if hashBlock(block) != block.Hash {
			return false
		}
		// invalid nonce
		if !strings.HasPrefix(block.Hash, strings.Repeat("0", defaultDifficulty)) {
			return false
		}
		return true
	}
	return false

}

func (ExampleChainService) AddBlock(block *models.ExampleChainBlock) error {
	if exists := mExchain.ExistsByPublicKey(block.PublicKey); exists {
		return errors.New("you have already addded block")
	}
	mExchain.AddBlock(block)
	tx := mTx.CreateCoinbaseTx(50, block.PublicKey)
	mTx.CreateTx(tx)
	mExchain.SetSummary(block)
	return nil
}
