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

type Tx struct {
	ID        string `json:"id"`
	Timestamp int    `json:"timestamp"`
	TxIns     TxInS  `json:"txIns"`
	TxOuts    TxOutS `json:"txOuts"`
}

type TxInS struct {
	From string  `json:"from"`
	V    []*TxIn `json:"v"`
}
type TxOutS []*TxOut
type TxS []*Tx

func (txs *TxS) String() string {
	s := "["
	for i, tx := range *txs {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", tx)
	}
	return s + "]"
}

func (txIns *TxInS) String() string {
	s := "From: " + txIns.From
	s = ", V: ["
	for i, txIn := range txIns.V {
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
