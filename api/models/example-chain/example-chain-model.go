package examplechain_model

import (
	"context"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	defaultDifficulty = 4
)

type IExchainModel interface {
	GetSummary() *ExampleChain
	FindAllBlocks() []*ExampleChainBlock
	ExistsByPublicKey(publicKey string) bool
	AddBlock(block *ExampleChainBlock)
}

type exchainModel struct {
}

var summary *ExampleChain = initSummary()

var ExchainModel IExchainModel = &exchainModel{}

func (exchainModel) GetSummary() *ExampleChain {
	return summary
}

func (exchainModel) FindAllBlocks() []*ExampleChainBlock {
	var blocks []*ExampleChainBlock
	cursor, err := db.ExampleChain.Find(context.TODO(), bson.D{})
	lib.HandleErr(err)
	err = cursor.All(context.TODO(), &blocks)
	lib.HandleErr(err)
	return blocks
}

func (exchainModel) ExistsByPublicKey(publicKey string) bool {
	count, err := db.ExampleChain.CountDocuments(context.TODO(), bson.D{{Key: "PublicKey", Value: publicKey}})
	lib.HandleErr(err)
	return count == 1
}

func (exchainModel) AddBlock(block *ExampleChainBlock) {
	_, err := db.ExampleChain.InsertOne(context.TODO(), block)
	lib.HandleErr(err)
}
