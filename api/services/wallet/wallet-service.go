package wallet_servie

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	models "github.com/onee-only/mempool-manager/api/models/wallet"
	"github.com/onee-only/mempool-manager/lib"
)

type IWalletService interface {
	New() *models.Wallet
	GetKeys(w *models.Wallet) (publicKey string, privateKey string)
	ValidateWallet(publicKey string, privateKey string) bool
}

type walletService struct{}

var mWallet models.IWalletModel = models.WalletModel
var WalletService walletService = walletService{}

func (walletService) New() *models.Wallet {

	wallet := &models.Wallet{}

	privateKey := mWallet.CreatePrivateKey()
	publicKey := mWallet.MakePublicKey(privateKey)

	wallet.SetPublicKey(publicKey)
	wallet.SetPrivateKey(privateKey)

	return wallet
}

func (walletService) GetKeys(w *models.Wallet) (publicKey string, privateKey string) {
	publicKey = w.GetPublicKey()
	privateKey = w.GetPrivateKeyStr()
	return
}

func (walletService) ValidateWallet(publicKey string, privateKey string) bool {
	wallet, err := mWallet.RestoreWallet(publicKey, privateKey)
	if lib.HasErr(err) {
		return false
	}
	hash := []byte("hi")
	r, s, err := ecdsa.Sign(rand.Reader, wallet.GetPrivateKey(), hash)
	if lib.HasErr(err) {
		return false
	}

	x, y, err := mWallet.EncodePublicKey(publicKey)
	if lib.HasErr(err) {
		return false
	}

	return ecdsa.Verify(&ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}, hash, r, s)
}
