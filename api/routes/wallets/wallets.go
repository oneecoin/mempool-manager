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

	walletRes := WalletDTO{}
	walletRes.setKeys(sWallet.GetKeys(wallet))

	c.JSON(http.StatusCreated, walletRes)
}

func GetTransactions(c *gin.Context) {

	publicKey := c.Param("publicKey")

	txs := peers.Peers.RequestTxs(publicKey)
	if len(txs) == 0 {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, txs)
	}
}

func GetBalance(c *gin.Context) {
	publicKey := c.Param("publicKey")

	balance := peers.Peers.RequestBalance(publicKey)
	c.JSON(http.StatusOK, struct {
		Balance int `json:"balance"`
	}{Balance: balance})
}

func VerifyWallet(c *gin.Context) {
	wallet := &WalletDTO{}
	err := c.BindJSON(wallet)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	valid := sWallet.ValidateWallet(wallet.PublicKey, wallet.PrivateKey)
	if !valid {
		c.Status(http.StatusNotAcceptable)
	} else {
		c.Status(http.StatusOK)
	}
}
