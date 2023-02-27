package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/routes"
	"github.com/onee-only/mempool-manager/middlewares"
)

func InitServer() {
	os.Setenv("TZ", "Asia/Seoul")
	app := gin.Default()

	// routes
	app.Use(middlewares.CORSMiddleware)
	routes.GetRoutes(app)

	app.Run(":8080")
}
