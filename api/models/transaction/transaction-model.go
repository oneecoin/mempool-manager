package transaction_model

import (
	"context"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/bson"
)

type ITxModel interface {
	GetAllTxs() *TxS
	CreateTx() *Tx
	GetUnOccupiedTxs() *TxS
	DeleteTxs(txIDs []string)
}

type txModel struct{}

const (
	maxTxsPerBlock = 4
)

var TxModel ITxModel = txModel{}

func (txModel) GetAllTxs() *TxS {

}

func (txModel) CreateTx() *Tx {

}

func (txModel) GetUnOccupiedTxs() *TxS {
	count := 0
	var txIDs []string
	for txID, occupied := range GetTxsOccupation() {
		if !occupied {
			txIDs = append(txIDs, txID)
			count++
		}
		if count == maxTxsPerBlock {
			break
		}
	}

	var txs *TxS

	var arr bson.A
	for _, txID := range txIDs {
		arr = append(arr, bson.M{"ID": txID})
	}

	filter := bson.D{{
		"$or", arr,
	}}

	cursor, err := db.Transactions.Find(context.TODO(), filter)
	lib.HandleErr(err)

	cursor.All(context.TODO(), txs)

	occupyTxs(txIDs)
	return txs
}

func (txModel) DeleteTxs(txIDs []string) {

	deleteTxsOccupation(txIDs)
}
