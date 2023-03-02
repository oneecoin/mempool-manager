package peers

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/onee-only/mempool-manager/api/ws/messages"
	"github.com/onee-only/mempool-manager/lib"
)

type TAddress struct {
	Host string
	Port string
}

type Peer struct {
	Conn         *websocket.Conn
	Inbox        chan []byte
	BlockInbox   chan []byte
	UTxOutsInbox chan messages.PayloadUTxOuts
	RejectCount  int
	PublicKey    string
	Address      TAddress
}

func (p Peer) GetAddress() string {
	return fmt.Sprintf("%s:%s", p.Address.Host, p.Address.Port)
}

func (p *Peer) closeConn() {
	log.Println("closing connection...")
	Peers.M.Lock()
	defer Peers.M.Unlock()
	p.Conn.Close()
	delete(Peers.V, p.GetAddress())
}

func (p *Peer) read() {
	defer p.closeConn()
	for {
		m := &messages.Message{}
		err := p.Conn.ReadJSON(m)
		if err != nil {
			log.Println("got err", err)
			break
		}
		Peers.handleMessage(m, p)
	}
}

func (p *Peer) write() {
	defer p.closeConn()
	for {
		m, ok := <-p.Inbox
		if !ok {
			break
		}
		err := p.Conn.WriteMessage(websocket.TextMessage, m)
		lib.HandleErr(err)
	}
}
