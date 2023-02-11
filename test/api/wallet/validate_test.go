package wallet_test

import (
	"testing"

	wallet_servie "github.com/onee-only/mempool-manager/api/services/wallet"
)

const (
	invalidPrivateKey = "207702010104207f2175b6e420ab2a995add6d45eef4de9b73418984a3b3c5601f472b89054976a00a06082a8648ce3d030107a1440342000413390be6f2805a24070a0eb890008a64dfae722911f7bceb466925652d73315a14a195f336d954644b2c2548a478582c3584ea52afbe992b163733fec374ea8e"
	validPrivateKey   = "307702010104207f2175b6e420ab2a995add6d45eef4de9b73418984a3b3c5601f472b89054976a00a06082a8648ce3d030107a1440342000413390be6f2805a24070a0eb890008a64dfae722911f7bceb466925652d73315a14a195f336d954644b2c2548a478582c3584ea52afbe992b163733fec374ea8e"
	publicKey         = "13390be6f2805a24070a0eb890008a64dfae722911f7bceb466925652d73315a14a195f336d954644b2c2548a478582c3584ea52afbe992b163733fec374ea8e"
)

func TestValidateWallet(t *testing.T) {
	t.Run("valid wallet", func(t *testing.T) {
		valid := wallet_servie.WalletService.ValidateWallet(publicKey, validPrivateKey)
		if !valid {
			t.Error("wallet isn't valid")
		}
	})
	t.Run("invalid wallet", func(t *testing.T) {
		valid := wallet_servie.WalletService.ValidateWallet(publicKey, invalidPrivateKey)
		if valid {
			t.Error("wallet is valid")
		}
	})
}
