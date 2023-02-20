package peers

import (
	"math/rand"
)

func getRandomPeer() *Peer {
	countLeft := 0
	if len(Peers.V) != 0 {
		countLeft = rand.Intn(len(Peers.V))
	}
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

func calculateHalfPeers() int {
	halfOfPeers := len(Peers.V) / 2
	if len(Peers.V)%2 != 0 {
		halfOfPeers++
	}
	return halfOfPeers
}
