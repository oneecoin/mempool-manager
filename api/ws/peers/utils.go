package peers

import (
	"encoding/json"
	"math/rand"

	"github.com/onee-only/mempool-manager/api/ws/messages"
	"github.com/onee-only/mempool-manager/lib"
)

func getRandomPeer() *Peer {
	countLeft := rand.Intn(len(Peers.V))
	var v *Peer
	for _, peer := range Peers.V {
		if countLeft == 0 {
			v = peer
			break
		}
		countLeft--
	}
	return v
}

func marshalPeerMessage(peer *Peer, kind messages.MessageKind) []byte {
	payload, err := json.Marshal(messages.PayloadPeer{
		PeerAddress: peer.GetAddress(),
	})
	lib.HandleErr(err)
	bytes, err := json.Marshal(messages.Message{
		Kind:    messages.MessagePeerRejected,
		Payload: payload,
	})
	lib.HandleErr(err)
	return bytes
}