package peers

import (
	"encoding/json"
	"sync"

	"github.com/onee-only/mempool-manager/api/ws/messages"
)

type TPeers struct {
	V map[string]*Peer
	M sync.Mutex
}

var Peers *TPeers = &TPeers{
	V: make(map[string]*Peer),
	M: sync.Mutex{},
}

func (*TPeers) InitPeer(p *Peer) {
	go p.read()
	go p.write()
	Peers.V[p.GetAddress()] = p
}

func (*TPeers) BroadcastNewPeer(p *Peer) {
	m := marshalPeerMessage(p, messages.MessageNewPeer)
	for _, peer := range Peers.V {
		peer.Inbox <- m
	}
}
func (*TPeers) BroadcastRejectPeer(p *Peer) {
	m := marshalPeerMessage(p, messages.MessagePeerRejected)
	for _, peer := range Peers.V {
		peer.Inbox <- m
	}
}

func (*TPeers) BroadcastNewTx() {
	// send count of transaction
}

func (*TPeers) RequestBlocks(page int) []byte {
	peer := getRandomPeer()
	// send request to random miner node
}

func (*TPeers) RequestBlock(hash string) []byte {
	peer := getRandomPeer()
	// send request to random miner node
}

func (*TPeers) handleMessage(m *messages.Message, p *Peer) {
	switch m.Kind {
	case messages.MessageRejectPeer:

		payload := &messages.PayloadPeer{}
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
		// retrieve transactions with coinbase transaction added
	}
}
