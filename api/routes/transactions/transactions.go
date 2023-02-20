package transactions

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	transaction_service "github.com/onee-only/mempool-manager/api/services/transaction"
	"github.com/onee-only/mempool-manager/lib"
)

var transactions transaction_service.ITxService = transaction_service.TxService

func GetAllTransactions(c *gin.Context) {
	txs := transactions.GetAllTxs()
	c.JSON(http.StatusOK, txs)
}

func CreateTransaction(c *gin.Context) {
	req, err := io.ReadAll(c.Request.Body)
	lib.HandleErr(err)

	var txCreateReq TxCreateRequest
	lib.FromJSON(req, &txCreateReq)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if txCreateReq.Amount == 0 {
		c.Status(http.StatusForbidden)
		return
	}

	transactions.CreateTx(
		txCreateReq.PrivateKey,
		txCreateReq.To,
		txCreateReq.Amount,
	)
	c.Status(http.StatusCreated)
}

func GetTransaction(c *gin.Context) {
	hash := c.Param("hash")
	tx := transactions.GetTx(hash)
	isProccessing := transactions.IsTxProcessing(hash)

	// should handle 404
	res := TxResponse{
		IsProccessing: isProccessing,
		Tx:            tx,
	}
	c.JSON(http.StatusOK, res)
}

func DeleteTransaction(c *gin.Context) {
	hash := c.Param("hash")
	err := transactions.TryDeleteTx(hash)
	if err != nil {
		c.Status(http.StatusAlreadyReported)
	} else {
		c.Status(http.StatusNoContent)
	}
}
