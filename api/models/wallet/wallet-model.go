package wallet_model

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"math/big"

	"github.com/onee-only/mempool-manager/lib"
)

type IWalletModel interface {
	CreatePrivateKey() (privateKey *ecdsa.PrivateKey)
	MakePublicKey(key *ecdsa.PrivateKey) string
	EncodePublicKey(publicKey string) (*big.Int, *big.Int, error)
	RestoreWallet(publicKey string, privateKey string) (*Wallet, error)
}

type walletModel struct{}

var WalletModel IWalletModel = &walletModel{}

func (walletModel) CreatePrivateKey() (privateKey *ecdsa.PrivateKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	lib.HandleErr(err)
	return
}

func (walletModel) MakePublicKey(key *ecdsa.PrivateKey) string {
	return encodeBigInts(key.X.Bytes(), key.Y.Bytes())
}

func (walletModel) EncodePublicKey(publicKey string) (*big.Int, *big.Int, error) {
	return restoreBigInts(publicKey)
}

func (walletModel) RestoreWallet(publicKey string, privateKey string) (*Wallet, error) {
	privKeyBytes, err := hex.DecodeString(privateKey)
	lib.HandleErr(err)
	privKeyObj, err := x509.ParseECPrivateKey(privKeyBytes)
	if err != nil {
		return nil, err
	}

	wallet := &Wallet{}
	wallet.SetPrivateKey(privKeyObj)
	wallet.SetPublicKey(publicKey)

	return wallet, nil
}
