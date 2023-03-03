package transaction_model

import (
	"context"
	"log"
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
	v: nil,
	m: sync.Mutex{},
}

func GetTxsOccupation() *txsOccupationMap {
	if txsOccupation.v == nil {
		txsOccupation.v = make(map[string]string)
		txsOccupation.m.Lock()
		defer txsOccupation.m.Unlock()
		initTxsOccupation()
	}
	return &txsOccupation
}

func initTxsOccupation() {
	log.Println("initializing map")
	opts := options.Find().SetProjection(bson.M{"id": 1})

	results := []txIDsResult{}
	cursor, err := db.Transactions.Find(context.TODO(), bson.D{}, opts)
	lib.HandleErr(err)
	log.Println("got map data")

	err = cursor.All(context.TODO(), &results)
	lib.HandleErr(err)
	log.Println("got map data again", results)

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
	for _, txID := range txIDs {
		txsMap.v[txID] = minerPublicKey
	}
}

func unOccupyTxs(publicKey string, txIDs []string) {
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()

	for txID, minerPK := range txsMap.v {
		if minerPK == publicKey {
			invalid := false
			for _, aTxID := range txIDs {
				if txID == aTxID {
					invalid = true
				}
			}
			if !invalid {
				txsMap.v[txID] = ""
			}
		}
	}
}

func createFilterByTxIDs(txIDs []string) primitive.D {
	var arr bson.A
	for _, txID := range txIDs {
		arr = append(arr, bson.M{"id": txID})
	}

	return bson.D{{
		Key: "$or", Value: arr,
	}}
}

func createFilterByTxInsFrom(publicKey string) primitive.D {
	return bson.D{{
		Key: "txins.from", Value: publicKey,
	}}
}
