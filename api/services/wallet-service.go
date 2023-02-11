package services

import (
	"github.com/onee-only/mempool-manager/api/models"
)

type walletService interface {
	New() *models.Wallet
	GetKeys(w *models.Wallet) (publicKey string, privateKey string)
	ValidateWallet(publicKey string, privateKey string) bool
}

type WalletService struct{}

var Wallet walletService = WalletService{}

func (WalletService) New() *models.Wallet {

	wallet := &models.Wallet{}

	privateKey := models.CreatePrivateKey()
	publicKey := models.MakePublicKey(privateKey)

	wallet.SetPublicKey(publicKey)
	wallet.SetPrivateKey(privateKey)

	return wallet
}

func (WalletService) GetKeys(w *models.Wallet) (publicKey string, privateKey string) {
	publicKey = w.GetPublicKey()
	privateKey = w.GetPrivateKeyStr()
	return
}

func (WalletService) ValidateWallet(publicKey string, privateKey string) bool {
	wallet := models.RestoreWallet(publicKey, privateKey)
}
