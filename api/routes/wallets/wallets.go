package wallets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/services"
)

func CreateWallet(c *gin.Context) {
	wallet := services.Wallet.New()

	walletRes := WalletResponse{}
	walletRes.setKeys(services.Wallet.GetKeys(wallet))

	c.JSON(http.StatusCreated, walletRes)
}
