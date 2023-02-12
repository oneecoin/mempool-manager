package ws

import "sync"

var Peers peers = peers{
	v: make(map[string]*peer),
	m: sync.Mutex{},
}

