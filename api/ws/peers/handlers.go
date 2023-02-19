package peers

import (
	"encoding/json"

	"github.com/onee-only/mempool-manager/api/ws/messages"
	"github.com/onee-only/mempool-manager/lib"
)

func (*TPeers) handleMessage(m *messages.Message, p *Peer) {
	switch m.Kind {
	case messages.MessageRejectPeer:

		payload := &messages.PayloadPeer{}
		json.Unmarshal(m.Payload, payload)

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
		json.Unmarshal(m.Payload, payload)

		txs := transactions.GetTxsForMining(p.PublicKey, payload.Count)

		var m []byte
		var err error
		if txs == nil {
			m, err = json.Marshal(messages.Message{
				Kind:    messages.MessageTxsDeclined,
				Payload: nil,
			})
			lib.HandleErr(err)
		} else {
			*txs = append(*txs, transactions.CreateCoinbaseTx(len(*txs), p.PublicKey))
			payload, err := json.Marshal(messages.PayloadTxs{
				Txs: *txs,
			})
			lib.HandleErr(err)

			m, err = json.Marshal(messages.Message{
				Kind:    messages.MessageMempoolTxsResponse,
				Payload: payload,
			})
			lib.HandleErr(err)
		}

		p.Inbox <- m

	case messages.MessageBlocksResponse:
		fallthrough
	case messages.MessageBlockResponse:
		p.BlockInbox <- m.Payload
	case messages.MessageUTxOutsResponse:
		data := messages.PayloadUTxOuts{}
		err := json.Unmarshal(m.Payload, &data)
		lib.HandleErr(err)
		p.UTxOutsInbox <- data

	}
}
