package wallets

type WalletResponse struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

func (w *WalletResponse) setKeys(publicKey string, privateKey string) {
	w.PublicKey = publicKey
	w.PrivateKey = privateKey
}
