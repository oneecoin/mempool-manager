package api

import (
	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/routes"
)

func InitServer() {

	app := gin.Default()

	// routes
	routes.GetRoutes(app)

	app.Run(":8080")
}
