package messages

type MessageKind int

const (

	// requests from miner
	MessageTxsRequest MessageKind = iota

	// responses from miner
	MessageBlocksResponse
	MessageBlockResponse

	// requests from mempool
	MessageNewPeer
	MessageBlocksRequest
	MessageBlockResquest

	// responses from mempool
	MessageTxsResponse

	// etc.
	MessageRejectPeer
	MessagePeerRejected
)
