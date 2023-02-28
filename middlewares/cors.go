package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
)

var CORSMiddleware = cors.New(cors.Config{
	AllowOrigins:    []string{"*"},
	AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Origin"},
	AllowWebSockets: true,
	MaxAge:          12 * time.Hour,
})
