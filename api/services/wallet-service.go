package services

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/onee-only/mempool-manager/api/models"
	"github.com/onee-only/mempool-manager/lib"
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
	hash := []byte("hi")
	r, s, err := ecdsa.Sign(rand.Reader, wallet.GetPrivateKey(), hash)
	lib.HandleErr(err)

	x, y, err := models.EncodePublicKey(publicKey)
	if err != nil {
		return false
	}

	return ecdsa.Verify(&ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}, hash, r, s)
}
