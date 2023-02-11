package routes

import (
	"github.com/gin-gonic/gin"
	examplechain "github.com/onee-only/mempool-manager/api/routes/example-chain"
	"github.com/onee-only/mempool-manager/api/routes/transactions"
	"github.com/onee-only/mempool-manager/api/routes/wallets"
	"github.com/onee-only/mempool-manager/api/ws"
)

func GetRoutes(router *gin.Engine) {

	// websocket upgrade
	router.POST("/ws", ws.UpgradeWS)

	// wallets
	w := router.Group("/wallets")
	{
		w.POST("", wallets.CreateWallet)

		w.GET("/:publicKey", wallets.GetTransactions)
		w.GET("/:publicKey/balance", wallets.GetBalance)
	}

	// mempool
	mp := router.Group("/mempool")
	{
		mp.GET("", transactions.GetAllTransactions)
		mp.POST("", transactions.CreateTransaction)

		mp.GET("/:hash", transactions.GetTransaction)
	}

	// example blockchain
	exchain := router.Group("/example-chain")
	{
		exchain.GET("", examplechain.GetAllBlocks)
		exchain.POST("", examplechain.CreateBlock)
	}
}
