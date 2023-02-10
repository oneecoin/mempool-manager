package db

import (
	"context"
	"os"

	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_NAME       = "oneecoin-mempool"
	EXAMPLE_CHAIN = "example-chain"
	TRANSACTONS   = "transactions"
)

type Database struct {
	ExampleChain *mongo.Collection
	Transactions *mongo.Collection
}

var DB = &Database{}

func InitDatabase() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("URI does not exist")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	lib.HandleErr(err)

	DB.ExampleChain = client.Database(DB_NAME).Collection(EXAMPLE_CHAIN)
	DB.Transactions = client.Database(DB_NAME).Collection(TRANSACTONS)

	return client
}
