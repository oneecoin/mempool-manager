package routes

import (
	"github.com/gin-gonic/gin"
	examplechain "github.com/onee-only/mempool-manager/api/routes/example-chain"
	"github.com/onee-only/mempool-manager/api/routes/transactions"
	"github.com/onee-only/mempool-manager/api/routes/wallets"
)

func GetRoutes(router *gin.Engine) {

	// wallets
	router.POST("/wallets", wallets.CreateWallet)

	// transactions
	tx := router.Group("/transactions")
	{
		tx.GET("", transactions.GetAllTransactions)
		tx.POST("", transactions.CreateTransaction)

		tx.GET("/:hash", transactions.GetTransaction)
	}

	// example blockchain
	exchain := router.Group("/example-chain")
	{
		exchain.GET("", examplechain.GetAllBlocks)
		exchain.POST("", examplechain.CreateBlock)
	}
}
