package transaction_model

import (
	"context"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
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

}

func (txModel) CreateTx(tx *Tx) {
	
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

	deleteTxsOccupation(txIDs)
}
