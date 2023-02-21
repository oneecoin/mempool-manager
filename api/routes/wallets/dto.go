package wallets

type WalletDTO struct {
	PrivateKey string `json:"privateKey" binding:"required"`
	PublicKey  string `json:"publicKey" binding:"required"`
}

func (w *WalletDTO) setKeys(publicKey string, privateKey string) {
	w.PublicKey = publicKey
	w.PrivateKey = privateKey
}
