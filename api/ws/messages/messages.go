package messages

type Message struct {
	Kind    MessageKind
	Payload []byte
}
