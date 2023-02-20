package messages

type MessageKind int

const (

	// requests from miner
	MessageMempoolTxsRequest MessageKind = iota

	// responses from miner
	MessageBlocksResponse
	MessageBlockResponse
	MessageUTxOutsResponse

	// requests from mempool
	MessageBlocksRequest
	MessageBlockRequest
	MessageUTxOutsRequest

	// responses from mempool
	MessageMempoolTxsResponse
	MessageTxsDeclined

	// etc.
	MessageRejectPeer
	MessagePeerRejected
	MessageBlockAdded
	MessageNewBlock
)
