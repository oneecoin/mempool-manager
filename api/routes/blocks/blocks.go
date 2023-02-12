package blocks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/ws/peers"
	"github.com/onee-only/mempool-manager/lib"
)

var prs *peers.TPeers = peers.Peers

func GetBlocks(c *gin.Context) {
	bytes, err := prs.RequestBlocks(1)
	if err != nil {
		// c.Status(http.StatusBadRequest)
		lib.HandleErr(err)
		return
	}
	c.JSON(http.StatusOK, bytes)
}

func GetBlock(c *gin.Context) {
	bytes, err := prs.RequestBlock("a")
	if err != nil {
		// c.Status(http.StatusBadRequest)
		lib.HandleErr(err)
		return
	}
	c.JSON(http.StatusOK, bytes)
}
