package peers

import (
	"encoding/json"
	"sync"

	transaction_model "github.com/onee-only/mempool-manager/api/models/transaction"
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

func (*TPeers) InitPeer(p *Peer) {
	Peers.M.Lock()
	defer Peers.M.Unlock()
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

	payload, err := json.Marshal(messages.PayloadPage{Page: page})
	lib.HandleErr(err)

	m, err := json.Marshal(messages.Message{
		Kind:    messages.MessageBlocksRequest,
		Payload: payload,
	})
	lib.HandleErr(err)

	peer.Inbox <- m
	block := <-peer.BlockInbox
	return block
}

func (*TPeers) RequestBlock(hash string) []byte {
	peer := getRandomPeer()

	payload, err := json.Marshal(messages.PayloadHash{Hash: hash})
	lib.HandleErr(err)

	m, err := json.Marshal(messages.Message{
		Kind:    messages.MessageBlockRequest,
		Payload: payload,
	})
	lib.HandleErr(err)

	peer.Inbox <- m
	block := <-peer.BlockInbox
	return block
}

func (*TPeers) GetUnSpentTxOuts(fromPublicKey string, amount int) (*transaction_model.UTxOutS, bool) {
	peer := getRandomPeer()

	payload, err := json.Marshal(messages.PayloadUTxOutsFilter{
		PublicKey: fromPublicKey,
		Amount:    amount,
	})
	lib.HandleErr(err)

	m, err := json.Marshal(messages.Message{
		Kind:    messages.MessageUTxOutsRequest,
		Payload: payload,
	})
	lib.HandleErr(err)

	peer.Inbox <- m

	uTxOuts := <-peer.UTxOutsInbox
	return &uTxOuts.UTxOuts, uTxOuts.Available
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

	case messages.MessageBlocksResponse:
		fallthrough
	case messages.MessageBlockResponse:
		p.BlockInbox <- m.Payload
	}
}
