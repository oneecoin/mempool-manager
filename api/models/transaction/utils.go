package transaction_model

import (
	"context"
	"sync"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type txIDsResult struct {
	ID string
}

type txInsAmountResult struct {
	TxOuts []struct {
		Amount int
	}
}

type txsOccupationMap struct {
	v map[string]string
	m sync.Mutex
}

var txsOccupation txsOccupationMap = txsOccupationMap{
	v: make(map[string]string),
	m: sync.Mutex{},
}

func GetTxsOccupation() *txsOccupationMap {
	if txsOccupation.v == nil {
		txsOccupation.m.Lock()
		defer txsOccupation.m.Lock()
		initTxsOccupation()
	}
	return &txsOccupation
}

func initTxsOccupation() {
	opts := options.Find().SetProjection(bson.D{{Key: "ID", Value: 1}})

	results := []txIDsResult{}
	cursor, err := db.Transactions.Find(context.TODO(), bson.D{}, opts)
	lib.HandleErr(err)

	cursor.All(context.TODO(), &results)

	for _, result := range results {
		txsOccupation.v[result.ID] = ""
	}
}

func deleteTxsOccupation(txIDs []string) {
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()
	for _, txID := range txIDs {
		delete(txsMap.v, txID)
	}
}

func occupyTxs(txIDs []string, minerPublicKey string) {
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()
	for _, txID := range txIDs {
		txsMap.v[txID] = minerPublicKey
	}
}

func createFilterByTxIDs(txIDs []string) primitive.D {
	var arr bson.A
	for _, txID := range txIDs {
		arr = append(arr, bson.M{"ID": txID})
	}

	return bson.D{{
		Key: "$or", Value: arr,
	}}
}

func createFilterByTxInsFrom(publicKey string) primitive.D {
	return bson.D{{
		Key: "TxIns.From", Value: publicKey,
	}}
}
