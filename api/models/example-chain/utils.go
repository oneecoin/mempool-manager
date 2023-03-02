package examplechain_model

import (
	"context"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initSummary() *ExampleChain {
	summary := &ExampleChain{
		Height:     0,
		LatestHash: "Genesis",
	}

	// is Collection empty?
	count, err := db.ExampleChain.EstimatedDocumentCount(context.TODO())
	lib.HandleErr(err)
	if count == 0 {
		return summary
	}

	// get latest block
	opts := options.Find()
	opts.SetLimit(1)
	opts.SetSort(bson.D{{Key: "height", Value: -1}})

	cursor, err := db.ExampleChain.Find(context.TODO(), bson.D{}, opts)
	lib.HandleErr(err)

	var result []*ExampleChainBlock
	cursor.All(context.TODO(), &result)

	if len(result) != 0 {
		summary.Height = result[0].Height
		summary.LatestHash = result[0].Hash
	}

	return summary
}
