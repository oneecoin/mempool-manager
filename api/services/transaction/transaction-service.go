package transaction_service

import (
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

// var blockchain *peers.TPeers = peers.Peers
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
	// blockchain.

	tx := &transaction_model.Tx{}

	transactions.CreateTx(tx)
}

func (txService) GetAllTxs() *transaction_model.TxS {
	return transactions.GetAllTxs()
}

func (txService) GetTx(hash string) *transaction_model.Tx {
	// return nil
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
