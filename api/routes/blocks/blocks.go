package blocks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/ws/peers"
)

var prs *peers.TPeers = peers.Peers

func GetBlocks(c *gin.Context) {
	var bq BlocksQuery
	err := c.BindQuery(&bq)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	bytes := prs.RequestBlocks(bq.Page)
	c.Data(http.StatusOK, "application/json", bytes)
}

func GetBlock(c *gin.Context) {
	hash := c.Param("hash")
	bytes := prs.RequestBlock(hash)
	// should add error handling. like 404
	c.Data(http.StatusOK, "application/json", bytes)
}

func GetSummary(c *gin.Context) {
	bytes := prs.GetLatest()
	c.Data(http.StatusOK, "application/json", bytes)
}
