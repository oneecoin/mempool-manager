package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/routes/blocks"
	examplechain "github.com/onee-only/mempool-manager/api/routes/example-chain"
	"github.com/onee-only/mempool-manager/api/routes/transactions"
	"github.com/onee-only/mempool-manager/api/routes/wallets"
	"github.com/onee-only/mempool-manager/api/ws"
)

func GetRoutes(router *gin.Engine) {

	// websocket upgrade
	router.GET("/ws", ws.UpgradeWS)

	// wallets
	w := router.Group("/wallets")
	{
		w.POST("", wallets.CreateWallet)

		w.GET("/:publicKey", wallets.GetTransactions)
		w.GET("/:publicKey/balance", wallets.GetBalance)
	}

	// blocks
	b := router.Group("/blocks")
	{
		b.GET("", blocks.GetBlocks)
		b.GET("/:hash", blocks.GetBlock)
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
