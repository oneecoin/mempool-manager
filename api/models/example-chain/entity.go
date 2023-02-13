package examplechain_model

type ExampleChain struct {
	Height     int
	LatestHash string
}

type ExampleChainBlock struct {
	Data      string `json:"data"`
	PublicKey string `json:"publicKey"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevHash"`
	Height    int    `json:"height"`
	Nonce     int    `json:"nonce"`
	Created   int    `json:"created"`
}
