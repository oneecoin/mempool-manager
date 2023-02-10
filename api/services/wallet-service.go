package services

import (
	"github.com/onee-only/mempool-manager/api/models"
)

type walletService interface {
	New() *models.Wallet
}

type WalletService struct{}

var Wallet walletService = WalletService{}

func (WalletService) New() *models.Wallet {

	wallet := &models.Wallet{}

	privateKey := models.CreatePrivateKey()
	publicKey := models.GetPublicKey(privateKey)

	wallet.PublicKey = publicKey
	wallet.SetPrivateKey(privateKey)

	return wallet
}
