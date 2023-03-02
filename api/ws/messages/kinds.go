package messages

type MessageKind int

const (

	// mempool-node

	// requests from miner
	MessageMempoolTxsRequest MessageKind = iota
	MessageInvalidTxsRequest

	// responses from miner
	MessageBlocksResponse
	MessageBlockResponse
	MessageUTxOutsResponse
	MessageNodeTxsResponse
	MessageBalanceResponse

	// requests from mempool
	MessageBlocksRequest
	MessageBlockRequest
	MessageUTxOutsRequest
	MessageNodeTxsRequest
	MessageBalanceRequest

	// responses from mempool
	MessageTxsMempoolResponse
	MessageTxsDeclined

	// etc.
	MessageRejectPeer
	MessagePeerRejected
	MessageNewBlock

	// node-node

	// download blockchain
	MessageDownloadRequest
)
