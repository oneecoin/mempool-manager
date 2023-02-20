package wallets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	wallet_servie "github.com/onee-only/mempool-manager/api/services/wallet"
	"github.com/onee-only/mempool-manager/api/ws/peers"
)

var sWallet wallet_servie.IWalletService = wallet_servie.WalletService

func CreateWallet(c *gin.Context) {
	wallet := sWallet.New()

	walletRes := WalletResponse{}
	walletRes.setKeys(sWallet.GetKeys(wallet))

	c.JSON(http.StatusCreated, walletRes)
}

func GetTransactions(c *gin.Context) {
	txs := <-peers.TxsInbox
	if len(txs) == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, txs)
	}
}

func GetBalance(c *gin.Context) {

}
