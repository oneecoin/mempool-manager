package messages

import transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"

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

type PayloadUTxOutsFilter struct {
	PublicKey string
	Amount    int
}

type PayloadUTxOuts struct {
	Available bool
	UTxOuts   transaction_model.UTxOutS
}
