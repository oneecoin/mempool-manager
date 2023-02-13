package transaction_model

import (
	"context"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type txIDsResult struct {
	ID string
}

var txsOccupation map[string]bool

func GetTxsOccupation() map[string]bool {
	if txsOccupation == nil {
		initTxsOccupation()
	}
	return txsOccupation
}

func initTxsOccupation() {
	txsOccupation = make(map[string]bool)

	opts := options.Find().SetProjection(bson.D{{"ID", 1}})

	results := []txIDsResult{}
	cursor, err := db.Transactions.Find(context.TODO(), bson.D{}, opts)
	lib.HandleErr(err)

	cursor.All(context.TODO(), &results)

	for _, result := range results {
		txsOccupation[result.ID] = false
	}
}

func deleteTxsOccupation(txIDs []string) {
	txsMap := GetTxsOccupation()
	for _, txID := range txIDs {
		delete(txsMap, txID)
	}
}

func occupyTxs(txIDs []string) {
	txsMap := GetTxsOccupation()
	for _, txID := range txIDs {
		txsMap[txID] = true
	}
}

func createFilterByTxIDs(txIDs []string) primitive.D {
	var arr bson.A
	for _, txID := range txIDs {
		arr = append(arr, bson.M{"ID": txID})
	}

	return bson.D{{
		"$or", arr,
	}}
}
