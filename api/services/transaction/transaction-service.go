package transaction_service

import (
	"errors"
	"time"

	transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"
	wallet_model "github.com/onee-only/mempool-manager/api/models/wallet"
	"github.com/onee-only/mempool-manager/api/ws/peers"
)

type ITxService interface {
	IsTxProcessing(txID string) bool
	CreateTx(privateKey, targetAddress string, amount int) error
	GetAllTxs() *transaction_model.TxS
	GetTx(hash string) *transaction_model.Tx
	DeleteTxs(minerPublicKey string)
	TryDeleteTx(txID string) error
}

type txService struct{}

var miners *peers.TPeers = peers.Peers
var transactions transaction_model.ITxModel = transaction_model.TxModel
var wallets wallet_model.IWalletModel = wallet_model.WalletModel
var TxService ITxService = txService{}

func (txService) IsTxProcessing(txID string) bool {
	return transactions.IsTxOccupied(txID)
}

func (txService) CreateTx(privateKey, targetAddress string, amount int) error {

	privKeyObj, err := wallets.GetPrivKeyObjFromString(privateKey)
	if err != nil {
		return err
	}

	fromPublicKey := wallets.GetPublicFromPrivate(privKeyObj)

	spent := transactions.GetSpentBalanceAmount(fromPublicKey)
	unSpentTxOuts, available := miners.GetUnSpentTxOuts(fromPublicKey, amount+spent)

	if !available {
		return errors.New("not enough money")
	}
	inputAmount := getAmountFromUTxouts(unSpentTxOuts)

	change := inputAmount - amount

	txIns := transaction_model.TxInS{From: fromPublicKey}
	txOuts := transaction_model.TxOutS{}

	for _, uTxOut := range *unSpentTxOuts {
		txIns.V = append(txIns.V, &transaction_model.TxIn{
			BlockHash: uTxOut.BlockHash,
			TxID:      uTxOut.TxID,
			Index:     uTxOut.Index,
			Signature: "",
		})
	}

	if change != 0 {
		txOuts = append(txOuts, &transaction_model.TxOut{
			PublicKey: fromPublicKey,
			Amount:    change,
		})
	}
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

	tx.ID = transactions.MakeTxID(tx)

	for _, txIn := range tx.TxIns.V {
		txIn.Signature = wallets.Sign(privKeyObj, tx.ID)
	}

	transactions.CreateTx(tx)
	return nil
}

func (txService) GetAllTxs() *transaction_model.TxS {
	return transactions.GetAllTxs()
}

func (txService) GetTx(txID string) *transaction_model.Tx {
	return transactions.GetTxByTxID(txID)
}

func (txService) DeleteTxs(minerPublicKey string) {
	transactions.DeleteTxs(minerPublicKey)
}

func (txService) TryDeleteTx(txID string) error {
	if err := transactions.DeleteTx(txID); err != nil {
		return err
	}
	return nil
}
