package transaction_model

import "fmt"

type TxIn struct {
	TxID      string `json:"txId"`
	Index     int    `json:"index"`
	Signature string `json:"signature"`
}
type TxOut struct {
	PublicKey string `json:"publicKey"`
	Amount    int    `json:"amount"`
}

type TxInS []*TxIn
type TxOutS []*TxOut

type Tx struct {
	ID        string `json:"id"`
	Timestamp int    `json:"timestamp"`
	TxIns     TxInS  `json:"txIns"`
	TxOuts    TxOutS `json:"txOuts"`
}

func (txIns *TxInS) String() string {
	s := "["
	for i, txIn := range *txIns {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", txIn)
	}
	return s + "]"
}

func (txOuts *TxOutS) String() string {
	s := "["
	for i, txOut := range *txOuts {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", txOut)
	}
	return s + "]"
}
