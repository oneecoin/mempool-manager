package transaction_service

import (
	transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"
)

func getAmountFromUTxouts(unSpentTxOuts *transaction_model.UTxOutS) int {
	sum := 0
	for _, uTxOut := range *unSpentTxOuts {
		sum += uTxOut.Amount
	}
	return sum
}
