package ws

import "sync"

const (
	MessageNewPeer = iota
)

type peers struct {
	v map[string]*peer
	m sync.Mutex
}

type peer struct {
	address struct {
		host string
		port string
	}
}
