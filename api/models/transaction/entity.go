package transaction_model

import "fmt"

type TxIn struct {
	BlockHash string `json:"blockHash"`
	TxID      string `json:"txId"`
	Index     int    `json:"index"`
	Signature string `json:"signature"`
}
type TxOut struct {
	PublicKey string `json:"publicKey"`
	Amount    int    `json:"amount"`
}

type UTxOut struct {
	BlockHash string `json:"blockHash"`
	TxID      string `json:"txId"`
	Index     int    `json:"index"`
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
type UTxOutS []*UTxOut
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

func (uTxOuts *UTxOutS) String() string {
	s := "["
	for i, uTxOut := range *uTxOuts {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", uTxOut)
	}
	return s + "]"
}

func (txIns *TxInS) String() string {
	s := "From: " + txIns.From
	s += ", V: ["
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
