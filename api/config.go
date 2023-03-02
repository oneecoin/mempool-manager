package api

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/routes"
	"github.com/onee-only/mempool-manager/lib"
	"github.com/onee-only/mempool-manager/middlewares"
)

func InitServer() {
	os.Setenv("TZ", "Asia/Seoul")
	router := gin.Default()

	// Load the SSL certificate
	cert, err := tls.LoadX509KeyPair("/etc/letsencrypt/live/mempool.oneecoin.site/fullchain.pem", "/etc/letsencrypt/live/mempool.oneecoin.site/privkey.pem")
	lib.HandleErr(err)

	// Create a TLS configuration with the certificate
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// routes
	router.Use(middlewares.CORSMiddleware)
	routes.GetRoutes(router)

	server := &http.Server{
		Addr:      ":443",
		Handler:   router,
		TLSConfig: config,
	}

	err = server.ListenAndServeTLS("", "")
	lib.HandleErr(err)
}
