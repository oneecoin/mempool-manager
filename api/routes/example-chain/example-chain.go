package examplechain

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	examplechain_model "github.com/onee-only/mempool-manager/api/models/example-chain"
	examplechain_service "github.com/onee-only/mempool-manager/api/services/example-chain"
	wallet_servie "github.com/onee-only/mempool-manager/api/services/wallet"
	"github.com/onee-only/mempool-manager/lib"
)

var exchain examplechain_service.IExampleChainService = examplechain_service.ExampleChain
var wallet wallet_servie.IWalletService = wallet_servie.WalletService

func GetAllBlocks(c *gin.Context) {
	blocks := exchain.GetAllBlocks()

	blocksRes := &BlocksResponse{
		Count:  len(blocks),
		Blocks: blocks,
	}
	c.JSON(http.StatusOK, blocksRes)
}

func CreateBlock(c *gin.Context) {
	req := &CreateBlockRequest{}
	err := c.BindJSON(req)
	lib.HandleErr(err)

	block := &examplechain_model.ExampleChainBlock{
		Data:      req.Block.Data,
		PublicKey: req.Block.PublicKey,
		Hash:      req.Block.Hash,
		PrevHash:  req.Block.PrevHash,
		Height:    req.Block.Height,
		Nonce:     req.Block.Nonce,
	}

	// validation
	valid := wallet.ValidateWallet(req.Block.PublicKey, req.PrivateKey)
	valid = valid && exchain.ValidateBlock(block)
	if !valid {
		c.Status(http.StatusNotAcceptable)
		return
	}

	block.Created = int(time.Now().Local().Unix())
	err = exchain.AddBlock(block)
	if err != nil {
		c.Status(http.StatusAlreadyReported)
	} else {
		c.JSON(http.StatusCreated, block)
	}
}

func GetSummary(c *gin.Context) {
	summary := exchain.GetSummary()
	c.JSON(http.StatusOK, summary)
}
