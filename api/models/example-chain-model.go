package models

import (
	"context"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	defaultDifficulty = 4
)

type ExampleChain struct {
	Height     int
	LatestHash string
}

var ChainSummary = ExampleChain{
	Height:     0,
	LatestHash: "",
}

type ExampleChainBlock struct {
	Data      string
	PublicKey string
	Hash      string
	PrevHash  string
	Height    int
	Nonce     int
	Created   string
}

func FindAllBlocks() []*ExampleChainBlock {
	var blocks []*ExampleChainBlock
	cursor, err := db.ExampleChain.Find(context.TODO(), bson.D{})
	lib.HandleErr(err)
	err = cursor.All(context.TODO(), &blocks)
	lib.HandleErr(err)
	return blocks
}
