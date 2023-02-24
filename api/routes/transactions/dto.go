package transactions

import transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"

type TxResponse struct {
	IsProccessing bool                  `json:"isProccessing"`
	Tx            *transaction_model.Tx `json:"tx"`
}

type TxsResponseElement struct {
	IsProccessing bool   `json:"isProccessing"`
	From          string `json:"from"`
	To            string `json:"to"`
	Amount        int    `json:"amount"`
}

type TxsResponse struct {
	Txs []*TxsResponseElement `json:"txs"`
}

type TxCreateRequest struct {
	PrivateKey string `json:"privateKey"`
	To         string `json:"to"`
	Amount     int    `json:"amount"`
}
