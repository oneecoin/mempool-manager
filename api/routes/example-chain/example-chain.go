package examplechain

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onee-only/mempool-manager/api/models"
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
	req := &CreateBlockRequest{}
	c.BindJSON(req)

	block := &models.ExampleChainBlock{
		Data:      req.Data,
		PublicKey: req.PublicKey,
		Hash:      req.Hash,
		PrevHash:  req.PrevHash,
		Height:    req.Height,
		Nonce:     req.Nonce,
	}

	valid := services.Wallet.ValidateWallet(req.PrivateKey, req.PublicKey)
	valid = valid && services.ExampleChain.ValidateBlock(block)
	if !valid {
		c.Status(http.StatusNotAcceptable)
		return
	}
	block.Created = time.Now().Local().String()
	services.ExampleChain.AddBlock(block)
	c.Status(http.StatusCreated)
}
