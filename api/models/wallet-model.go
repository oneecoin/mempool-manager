package models

import "crypto/ecdsa"

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string
}
