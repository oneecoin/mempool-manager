package transaction_service

import (
	"crypto/sha256"
	"fmt"

	transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"
)

func getAmountFromUTxouts(unSpentTxOuts *transaction_model.UTxOutS) int {
	sum := 0
	for _, uTxOut := range *unSpentTxOuts {
		sum += uTxOut.Amount
	}
	return sum
}

func makeTxID(tx *transaction_model.Tx) string {
	bytes := []byte(fmt.Sprintf("%v", tx))
	hash := sha256.Sum256(bytes)
	return fmt.Sprintf("%s", hash)
}
