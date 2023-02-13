package examplechain_model

type ExampleChain struct {
	Height     int
	LatestHash string
}

type ExampleChainBlock struct {
	Data      string
	PublicKey string
	Hash      string
	PrevHash  string
	Height    int
	Nonce     int
	Created   int
}
