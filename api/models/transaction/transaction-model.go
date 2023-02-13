package transaction_model

type ITxModel interface {
	GetAllTxs() *TxS
	CreateTx() *Tx
	GetUnOccupiedTxs() *TxS
	OccupyTxs(txIDs []string) error
}

type txModel struct{}

var TxModel ITxModel = txModel{}
