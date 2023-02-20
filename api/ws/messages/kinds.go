package messages

type MessageKind int

const (

	// requests from miner
	MessageMempoolTxsRequest MessageKind = iota
	MessageInvalidTxsRequest

	// responses from miner
	MessageBlocksResponse
	MessageBlockResponse
	MessageUTxOutsResponse
	MessageNodeTxsResponse

	// requests from mempool
	MessageBlocksRequest
	MessageBlockRequest
	MessageUTxOutsRequest
	MessageNodeTxsRequest

	// responses from mempool
	MessageMempoolTxsResponse
	MessageTxsDeclined

	// etc.
	MessageRejectPeer
	MessagePeerRejected
	MessageNewBlock
)
