package peers

import (
	"log"

	transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"
	"github.com/onee-only/mempool-manager/api/ws/messages"
	"github.com/onee-only/mempool-manager/lib"
)

var txsInbox = make(chan transaction_model.TxS)
var balanceInbox = make(chan int)

func (*TPeers) handleMessage(m *messages.Message, p *Peer) {
	switch m.Kind {
	case messages.MessageRejectPeer:

		payload := &messages.PayloadPeer{}
		lib.FromJSON(m.Payload, payload)

		peer, exists := Peers.V[payload.PeerAddress]
		if exists {
			peer.RejectCount++
			halfOfPeers := calculateHalfPeers()
			if peer.RejectCount > halfOfPeers {
				close(peer.Inbox)
				peer.closeConn()
				Peers.BroadcastRejectPeer(peer)
			}
		}
	case messages.MessageMempoolTxsRequest:

		payload := &messages.PayloadCount{}
		lib.FromJSON(m.Payload, payload)

		log.Println("got request from ", p.PublicKey, " and need ", payload.Count)
		txs := transactions.GetTxsForMining(p.PublicKey, payload.Count)

		var m []byte
		if txs == nil {
			m = lib.ToJSON(messages.Message{
				Kind:    messages.MessageTxsDeclined,
				Payload: lib.ToJSON(messages.PayloadCount{Count: 0}),
			})
			log.Println("this is unacceptable")
		} else {
			txs = append(txs, transactions.CreateCoinbaseTx(len(txs), p.PublicKey))
			payload := lib.ToJSON(messages.PayloadTxs{
				Txs: txs,
			})

			m = lib.ToJSON(messages.Message{
				Kind:    messages.MessageTxsMempoolResponse,
				Payload: payload,
			})
		}
		p.Inbox <- m

	case messages.MessageBlocksResponse:
		fallthrough
	case messages.MessageBlockResponse:
		p.BlockInbox <- m.Payload
	case messages.MessageUTxOutsResponse:
		data := messages.PayloadUTxOuts{}
		lib.FromJSON(m.Payload, &data)
		p.UTxOutsInbox <- data
	case messages.MessageNewBlock:
		payload := &messages.PayloadPeer{}
		lib.FromJSON(m.Payload, payload)
		handleNewBlock(payload.PeerAddress)
	case messages.MessageNodeTxsResponse:
		payload := &messages.PayloadTxs{}
		lib.FromJSON(m.Payload, payload)
		txsInbox <- payload.Txs
	case messages.MessageInvalidTxsRequest:

		payload := &messages.PayloadTxs{}
		lib.FromJSON(m.Payload, payload)

		transactions.HandleInvalidTxs(payload.Txs, p.PublicKey)
	case messages.MessageBalanceResponse:
		payload := &messages.PayloadCount{}
		lib.FromJSON(m.Payload, payload)

		balanceInbox <- payload.Count
	}
}
