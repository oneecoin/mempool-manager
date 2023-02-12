package messages

type Message struct {
	Kind    MessageKind
	Payload []byte
}

type PayloadPeer struct {
	PeerAddress string
}
