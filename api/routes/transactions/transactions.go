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

	txsResElems := []*TxsResponseElement{}

	for _, tx := range txs {
		txElem := TxsResponseElement{
			TxID:          tx.ID,
			IsProccessing: transactions.IsTxProcessing(tx.ID),
			From:          tx.TxIns.From,
			To:            tx.TxOuts[0].PublicKey,
			Amount:        tx.TxOuts[0].Amount,
		}
		txsResElems = append(txsResElems, &txElem)
	}

	c.JSON(http.StatusOK, txsResElems)
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

	err = transactions.CreateTx(
		txCreateReq.PrivateKey,
		txCreateReq.To,
		txCreateReq.Amount,
	)
	if err != nil {
		c.Status(http.StatusForbidden)
		c.Header("err", err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

func GetTransaction(c *gin.Context) {
	hash := c.Param("hash")
	tx := transactions.GetTx(hash)
	isProccessing := transactions.IsTxProcessing(hash)
	if tx.ID == "" {
		c.Status(http.StatusNotFound)
		return
	}
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
