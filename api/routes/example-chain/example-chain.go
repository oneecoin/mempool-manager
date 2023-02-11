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
		Data:      req.Data,
		PublicKey: req.PublicKey,
		Hash:      req.Hash,
		PrevHash:  req.PrevHash,
		Height:    req.Height,
		Nonce:     req.Nonce,
	}

	valid := wallet.ValidateWallet(req.PublicKey, req.PrivateKey)
	valid = valid && exchain.ValidateBlock(block)
	if !valid {
		c.Status(http.StatusNotAcceptable)
		return
	}
	block.Created = time.Now().Local().String()
	exchain.AddBlock(block)
	c.JSON(http.StatusCreated, block)
}
