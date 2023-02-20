package peers

import (
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
	payload := lib.ToJSON(messages.PayloadPeer{
		PeerAddress: p.GetAddress(),
	})
	m := lib.ToJSON(messages.Message{
		Kind:    messages.MessagePeerRejected,
		Payload: payload,
	})
	for _, peer := range Peers.V {
		peer.Inbox <- m
	}
}

func (*TPeers) RequestBlocks(page int) []byte {
	peer := getRandomPeer()

	payload := lib.ToJSON(messages.PayloadPage{Page: page})

	m := lib.ToJSON(messages.Message{
		Kind:    messages.MessageBlocksRequest,
		Payload: payload,
	})

	peer.Inbox <- m
	block := <-peer.BlockInbox
	if block == nil {
		return Peers.RequestBlocks(page)
	}
	return block
}

func (*TPeers) RequestBlock(hash string) []byte {
	peer := getRandomPeer()

	payload := lib.ToJSON(messages.PayloadHash{Hash: hash})

	m := lib.ToJSON(messages.Message{
		Kind:    messages.MessageBlockRequest,
		Payload: payload,
	})

	peer.Inbox <- m
	block := <-peer.BlockInbox
	if block == nil {
		return Peers.RequestBlock(hash)
	}
	return block
}

func (*TPeers) GetUnSpentTxOuts(fromPublicKey string, amount int) (*transaction_model.UTxOutS, bool) {
	peer := getRandomPeer()

	payload := lib.ToJSON(messages.PayloadUTxOutsFilter{
		PublicKey: fromPublicKey,
		Amount:    amount,
	})

	m := lib.ToJSON(messages.Message{
		Kind:    messages.MessageUTxOutsRequest,
		Payload: payload,
	})

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

var newBlockMap = make(map[string]int)

func handleNewBlock(address string) {
	newBlockMap[address]++
	if newBlockMap[address] >= calculateHalfPeers() {
		transactions.DeleteTxs(Peers.V[address].PublicKey)
		delete(newBlockMap, address)
	}
}
