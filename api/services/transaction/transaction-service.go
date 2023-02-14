package transaction_service

import (
	"errors"
	"time"

	transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"
	wallet_model "github.com/onee-only/mempool-manager/api/models/wallet"
	"github.com/onee-only/mempool-manager/api/ws/peers"
)

type ITxService interface {
	GetTxsForMining() *transaction_model.TxS
	IsTxProcessing(txID string) bool
	CreateTx(privateKey, targetAddress string, amount int) error
	GetAllTxs() *transaction_model.TxS
	GetTx(hash string) *transaction_model.Tx
	DeleteTxs(txIDs []string)
	TryDeleteTx(txID string) error
}

type txService struct{}

var miners *peers.TPeers = peers.Peers
var transactions transaction_model.ITxModel = transaction_model.TxModel
var wallets wallet_model.IWalletModel = wallet_model.WalletModel
var TxService ITxService = txService{}

func (txService) GetTxsForMining() *transaction_model.TxS {
	return transactions.GetUnOccupiedTxs()
}

func (txService) IsTxProcessing(txID string) bool {
	return transactions.IsTxOccupied(txID)
}

func (txService) CreateTx(privateKey, targetAddress string, amount int) error {
	unSpentTxOuts := miners.GetUnSpentTxOuts()

	balance := getBalanceFromUTxouts(unSpentTxOuts)
	balance -= transactions.GetSpentBalance()

	if balance < amount {
		return errors.New("not enough money")
	}

	txIns := transaction_model.TxInS{}
	txOuts := transaction_model.TxOutS{}

	txOuts = append(txOuts, &transaction_model.TxOut{
		PublicKey: targetAddress,
		Amount:    amount,
	})

	

	tx := &transaction_model.Tx{
		ID:        "",
		TxIns:     txIns,
		TxOuts:    txOuts,
		Timestamp: int(time.Now().Local().Unix()),
	}

	tx.ID = makeTxID(tx)

	transactions.CreateTx(tx)
	return nil
}

func (txService) GetAllTxs() *transaction_model.TxS {
	return transactions.GetAllTxs()
}

func (txService) GetTx(txID string) *transaction_model.Tx {
	return transactions.GetTxByTxID(txID)
}

func (txService) DeleteTxs(txIDs []string) {
	transactions.DeleteTxs(txIDs)
}

func (txService) TryDeleteTx(txID string) error {
	if err := transactions.DeleteTx(txID); err != nil {
		return err
	}
	return nil
}
