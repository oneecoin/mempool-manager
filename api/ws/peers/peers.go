package peers

import (
	"encoding/json"
	"sync"

	"github.com/onee-only/mempool-manager/api/ws/messages"
	"github.com/onee-only/mempool-manager/lib"
)

type TPeers struct {
	V map[string]*Peer
	M sync.Mutex
}

var Peers *TPeers = &TPeers{
	V: make(map[string]*Peer),
	M: sync.Mutex{},
}

func (*TPeers) BroadcastNewPeer(p *Peer) {

}
func (*TPeers) BroadcastRejectPeer(p *Peer) {
	payload, err := json.Marshal(messages.PayloadRejectPeer{
		PeerAddress: p.GetAddress(),
	})
	lib.HandleErr(err)
	m, err := json.Marshal(messages.Message{
		Kind:    messages.MessagePeerRejected,
		Payload: payload,
	})
	lib.HandleErr(err)
	for _, peer := range Peers.V {
		peer.Inbox <- m
	}
}
func (*TPeers) BroadcastNewTx() {

}
func (*TPeers) RequestBlocks(page int) {

}
func (*TPeers) handleMessage(m *messages.Message, p *Peer) {
	switch m.Kind {
	case messages.MessageRejectPeer:

		payload := &messages.PayloadRejectPeer{}
		json.Unmarshal(m.Payload, payload)

		peer, exists := Peers.V[payload.PeerAddress]
		if exists {
			peer.RejectCount++
			halfOfPeers := len(Peers.V) / 2
			if len(Peers.V)%2 != 0 {
				halfOfPeers++
			}
			if peer.RejectCount > halfOfPeers {
				close(peer.Inbox)
				peer.closeConn()
				Peers.BroadcastRejectPeer(peer)
			}
		}
	case messages.MessageTxsRequest:
	}
}
