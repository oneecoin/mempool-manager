package blocks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/ws/peers"
)

var prs *peers.TPeers = peers.Peers

func GetBlocks(c *gin.Context) {
	bytes := prs.RequestBlocks(1)
	c.JSON(http.StatusOK, bytes)
}

func GetBlock(c *gin.Context) {
	bytes := prs.RequestBlock("a")
	c.JSON(http.StatusOK, bytes)
}
