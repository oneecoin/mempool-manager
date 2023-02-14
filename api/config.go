package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/routes"
)

func InitServer() {
	os.Setenv("TZ", "Asia/Seoul")
	app := gin.Default()

	// routes
	routes.GetRoutes(app)

	app.Run(":8080")
}
