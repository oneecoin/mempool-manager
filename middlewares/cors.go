package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
)

var CORSMiddleware = cors.New(cors.Config{
	AllowAllOrigins: true,
	AllowMethods:    []string{"POST", "GET", "PUT", "DELETE", "HEAD"},
	AllowWebSockets: true,
	MaxAge:          12 * time.Hour,
})
