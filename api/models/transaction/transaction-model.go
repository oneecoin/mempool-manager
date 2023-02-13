package transaction_model

import (
	"context"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/bson"
)

type ITxModel interface {
	GetAllTxs() *TxS
	CreateTx(tx *Tx)
	GetUnOccupiedTxs() *TxS
	DeleteTxs(txIDs []string)
}

type txModel struct{}

const (
	maxTxsPerBlock = 4
)

var TxModel ITxModel = txModel{}

func (txModel) GetAllTxs() *TxS {
	var txs *TxS

	cursor, err := db.Transactions.Find(context.TODO(), bson.D{})
	lib.HandleErr(err)
	err = cursor.All(context.TODO(), txs)
	lib.HandleErr(err)

	return txs
}

func (txModel) CreateTx(tx *Tx) {
	_, err := db.Transactions.InsertOne(context.TODO(), tx)
	lib.HandleErr(err)
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()
	txsMap.v[tx.ID] = false
}

func IsTxOccupied(txId string) bool {
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()
	return txsMap.v[txId]
}

func (txModel) GetUnOccupiedTxs() *TxS {
	count := 0
	var txIDs []string
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()
	for txID, occupied := range txsMap.v {
		if !occupied {
			txIDs = append(txIDs, txID)
			count++
		}
		if count == maxTxsPerBlock {
			break
		}
	}
	if len(txIDs) == 0 {
		return nil
	}

	var txs *TxS

	filter := createFilterByTxIDs(txIDs)

	cursor, err := db.Transactions.Find(context.TODO(), filter)
	lib.HandleErr(err)

	cursor.All(context.TODO(), txs)

	occupyTxs(txIDs)
	return txs
}

func (txModel) DeleteTxs(txIDs []string) {
	filter := createFilterByTxIDs(txIDs)
	_, err := db.ExampleChain.DeleteMany(context.TODO(), filter)
	lib.HandleErr(err)
	deleteTxsOccupation(txIDs)
}
