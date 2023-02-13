package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/onee-only/mempool-manager/api/ws/peers"
	"github.com/onee-only/mempool-manager/lib"
)

var wsUpgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

var prs *peers.TPeers = peers.Peers

func UpgradeWS(c *gin.Context) {

	publicKey := c.Request.URL.Query().Get("publicKey")
	port := c.Request.URL.Query().Get("port")
	host := strings.Split(c.Request.RemoteAddr, ":")[0]
	address := fmt.Sprintf("%s:%s", host, port)

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		// send http request to the address
		res, err := http.Get("http://" + address)
		if err != nil {
			return false
		}
		if res.StatusCode != http.StatusAccepted {
			return false
		}

		defer res.Body.Close()
		a := struct{ PublicKey string }{}
		err = json.NewDecoder(res.Body).Decode(&a)
		if err != nil {
			return false
		}
		if a.PublicKey != publicKey {
			return false
		}
		return true
	}
	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	lib.HandleErr(err)

	p := &peers.Peer{
		Conn:        conn,
		Inbox:       make(chan []byte),
		BlockInbox:  make(chan []byte),
		RejectCount: 0,
		PublicKey:   publicKey,
		Address: peers.TAddress{
			Host: host,
			Port: port,
		},
	}

	// add this connection to peers map
	prs.BroadcastNewPeer(p)
	prs.InitPeer(p)
}
