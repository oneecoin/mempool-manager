package transaction_model

import (
	"context"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type txIDsResult struct {
	ID string
}

var TxsOccupation map[string]bool

func InitTxsOccupation() map[string]bool {
	txsOccupation := make(map[string]bool)

	opts := options.Find().SetProjection(bson.D{{"ID", 1}})

	results := []txIDsResult{}
	cursor, err := db.Transactions.Find(context.TODO(), bson.D{}, opts)
	lib.HandleErr(err)

	cursor.All(context.TODO(), &results)

	for _, result := range results {
		txsOccupation[result.ID] = false
	}

	return txsOccupation
}
