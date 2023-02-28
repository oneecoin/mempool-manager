package transaction_model

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/onee-only/mempool-manager/db"
	"github.com/onee-only/mempool-manager/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ITxModel interface {
	GetAllTxs() TxS
	CreateTx(tx *Tx)
	IsTxOccupied(txID string) bool
	GetTxsForMining(minerPublicKey string, minCount int) TxS
	DeleteTxs(minerPublicKey string)
	DeleteTx(txID string) error
	GetSpentBalanceAmount(fromPublicKey string) int
	GetTxByTxID(txID string) *Tx
	CreateCoinbaseTx(txCount int, targetPublicKey string) *Tx
	MakeTxID(tx *Tx) string
	HandleInvalidTxs(txs TxS, publicKey string)
}

type txModel struct{}

const (
	maxTxsPerBlock = 4
	rewardPerTx    = 10
)

var TxModel ITxModel = txModel{}

func (txModel) GetAllTxs() TxS {
	var txs TxS

	cursor, err := db.Transactions.Find(context.TODO(), bson.D{})
	lib.HandleErr(err)
	err = cursor.All(context.TODO(), &txs)
	lib.HandleErr(err)

	return txs
}

func (txModel) CreateTx(tx *Tx) {
	_, err := db.Transactions.InsertOne(context.TODO(), tx)
	lib.HandleErr(err)
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()
	txsMap.v[tx.ID] = ""
}

func (txModel) IsTxOccupied(txID string) bool {
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()
	return txsMap.v[txID] != ""
}

func (txModel) GetTxsForMining(minerPublicKey string, minCount int) TxS {
	count := 0
	var txIDs []string
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()
	for txID, publicKey := range txsMap.v {
		if publicKey == "" {
			txIDs = append(txIDs, txID)
			count++
		}
		if count == maxTxsPerBlock {
			break
		}
	}
	if count <= minCount {
		return nil
	}

	var txs TxS

	filter := createFilterByTxIDs(txIDs)

	cursor, err := db.Transactions.Find(context.TODO(), filter)
	lib.HandleErr(err)

	cursor.All(context.TODO(), txs)

	occupyTxs(txIDs, minerPublicKey)
	return txs
}

func (txModel) DeleteTxs(minerPublicKey string) {
	txIDs := []string{}

	for txID, minerPK := range GetTxsOccupation().v {
		if minerPK == minerPublicKey {
			txIDs = append(txIDs, txID)
		}
	}

	filter := createFilterByTxIDs(txIDs)
	_, err := db.ExampleChain.DeleteMany(context.TODO(), filter)
	lib.HandleErr(err)
	deleteTxsOccupation(txIDs)
}

func (txModel) DeleteTx(txID string) error {
	txsMap := GetTxsOccupation()
	txsMap.m.Lock()
	defer txsMap.m.Unlock()
	if publicKey := txsMap.v[txID]; publicKey != "" {
		return errors.New("already taken")
	}
	_, err := db.Transactions.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: txID}})
	lib.HandleErr(err)
	return nil
}

func (txModel) GetSpentBalanceAmount(fromPublicKey string) int {

	filter := createFilterByTxInsFrom(fromPublicKey)
	opts := options.Find().SetProjection(bson.D{{Key: "txouts.amount", Value: 1}})
	cursor, err := db.Transactions.Find(context.TODO(), filter, opts)
	lib.HandleErr(err)

	var result txInsAmountResult
	err = cursor.All(context.TODO(), &result)
	lib.HandleErr(err)

	amount := 0
	for _, txOut := range result.TxOuts {
		amount += txOut.Amount
	}
	return amount
}

func (txModel) GetTxByTxID(txID string) *Tx {
	var tx *Tx
	cursor := db.Transactions.FindOne(context.TODO(), bson.D{{Key: "id", Value: txID}})
	err := cursor.Decode(tx)
	lib.HandleErr(err)
	return tx
}

func (txModel) CreateCoinbaseTx(txCount int, targetPublicKey string) *Tx {
	tx := &Tx{
		Timestamp: int(time.Now().Local().Unix()),
		TxIns: TxInS{
			From: "COINBASE",
			V: []*TxIn{
				{
					BlockHash: "",
					TxID:      "",
					Index:     -1,
					Signature: "",
				},
			},
		},
		TxOuts: TxOutS{
			{
				PublicKey: targetPublicKey,
				Amount:    txCount * rewardPerTx,
			},
		},
	}
	tx.ID = TxModel.MakeTxID(tx)
	return tx
}

func (txModel) MakeTxID(tx *Tx) string {
	bytes := []byte(fmt.Sprintf("%v", tx))
	hash := sha256.Sum256(bytes)
	return fmt.Sprintf("%x", hash)
}

func (txModel) HandleInvalidTxs(txs TxS, publicKey string) {
	var txIDs []string
	for _, tx := range txs {
		txIDs = append(txIDs, tx.ID)
	}
	unOccupyTxs(publicKey, txIDs)
	TxModel.DeleteTxs(publicKey)
}
