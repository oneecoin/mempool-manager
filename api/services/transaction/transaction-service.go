package transaction_service

import (
	transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"
	wallet_model "github.com/onee-only/mempool-manager/api/models/wallet"
)

type ITxService interface {
	GetTxsForMining() *transaction_model.TxS
	IsTxProcessing(txID string) bool
	CreateTx() error
	GetAllTxs() *transaction_model.TxS
	DeleteTxs(txIDs []string)
	TryDeleteTx(txID string) error
}

type txService struct{}

var transaction transaction_model.ITxModel = transaction_model.TxModel
var wallet wallet_model.IWalletModel = wallet_model.WalletModel
var TxService ITxService = txService{}
