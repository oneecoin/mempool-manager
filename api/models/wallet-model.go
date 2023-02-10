package models

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"

	"github.com/onee-only/mempool-manager/lib"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  string
}

func (w *Wallet) SetPrivateKey(privateKey *ecdsa.PrivateKey) {
	w.privateKey = privateKey
}

// returns private key with format of 'ecdsa.PrivateKey'
func (w Wallet) GetPrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

// returns private key with format of 'string'
func (w *Wallet) GetPrivateKeyStr() string {
	bytes, err := x509.MarshalECPrivateKey(w.privateKey)
	lib.HandleErr(err)
	return fmt.Sprintf("%x", bytes)
}

func (w *Wallet) SetPublicKey(publicKey string) {
	w.publicKey = publicKey
}

// returns private key with format of 'ecdsa.PrivateKey'
func (w Wallet) GetPublicKey() string {
	return w.publicKey
}

func CreatePrivateKey() (privateKey *ecdsa.PrivateKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	lib.HandleErr(err)
	return
}

func MakePublicKey(key *ecdsa.PrivateKey) string {
	return encodeBigInts(key.X.Bytes(), key.Y.Bytes())
}

func encodeBigInts(a, b []byte) string {
	bytes := append(a, b...)
	return fmt.Sprintf("%x", bytes)
}
