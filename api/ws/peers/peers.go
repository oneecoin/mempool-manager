package peers

import (
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

func (*TPeers) BroadcastNewPeer(p *Peer) {
}
func (*TPeers) BroadcastRejectPeer(p *Peer) {

}
func (*TPeers) BroadcastNewTx() {

}
func (*TPeers) RequestBlocks(page int) {

}
func (*TPeers) handleMessage(m *messages.Message, p *Peer) {
	switch m.Kind {
	case messages.MessageRejectPeer:
		
	case messages.MessageTxsRequest:
	}
}
