package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

const (
	messageNewPeer = iota
)

type message struct {
	kind    string
	payload []byte
}

type peers struct {
	v map[string]*peer
	m sync.Mutex
}

type tAddress struct {
	host string
	port string
}

type peer struct {
	conn        *websocket.Conn
	inbox       chan []byte
	rejectCount int
	publicKey   string
	address     tAddress
}
