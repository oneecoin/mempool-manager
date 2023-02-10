package examplechain

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/services"
)

func GetAllBlocks(c *gin.Context) {
	blocks := services.ExampleChain.GetAllBlocks()

	blocksRes := &BlocksResponse{
		Count:  len(blocks),
		Blocks: blocks,
	}
	c.JSON(http.StatusOK, blocksRes)
}

func CreateBlock(c *gin.Context) {

}
