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
var transactions transaction_model.ITxModel = transaction_model.TxModel

func (*TPeers) InitPeer(p *Peer) {
	Peers.M.Lock()
	defer Peers.M.Unlock()
	go p.read()
	go p.write()
	Peers.V[p.GetAddress()] = p
}

func (*TPeers) BroadcastRejectPeer(p *Peer) {
	payload, err := json.Marshal(messages.PayloadPeer{
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
	if block == nil {
		return Peers.RequestBlocks(page)
	}
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
	if block == nil {
		return Peers.RequestBlock(hash)
	}
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

func (*TPeers) GetAllPeers(exclude string) *[]string {
	var peerList []string
	Peers.M.Lock()
	defer Peers.M.Unlock()
	for address := range Peers.V {
		if address != exclude {
			peerList = append(peerList, address)
		}
	}
	return &peerList
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
