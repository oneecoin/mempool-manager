package messages

type Message struct {
	Kind    MessageKind
	Payload []byte
}

type PayloadPeer struct {
	PeerAddress string
}

type PayloadPage struct {
	Page int
}

type PayloadHash struct {
	Hash string
}
