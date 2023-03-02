package ws

import (
	"fmt"
	"net/http"

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

	publicKey := c.Query("publicKey")
	port := c.Query("port")
	host := c.ClientIP()
	// address := fmt.Sprintf("%s:%s", host, port)

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		// send http request to the address
		// res, err := http.Get("http://" + address + "/check")
		// if err != nil {
		// 	fmt.Println(err)
		// 	return false
		// }
		// if res.StatusCode != http.StatusAccepted {
		// 	return false
		// }

		// defer res.Body.Close()
		// a := struct{ PublicKey string }{}
		// err = json.NewDecoder(res.Body).Decode(&a)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return false
		// }
		// if a.PublicKey != publicKey {
		// 	return false
		// }
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
	prs.InitPeer(p)
}

func GetPeersCount(c *gin.Context) {
	count := len(prs.V)
	c.JSON(http.StatusOK, struct {
		Count int `json:"count"`
	}{
		Count: count,
	})
}

func GetPeers(c *gin.Context) {

	host := c.ClientIP()
	port := c.Query("port")
	addr := fmt.Sprintf("%s:%s", host, port)
	if _, exists := prs.V[addr]; !exists {
		c.Status(http.StatusUnauthorized)
		return
	}

	peerList := prs.GetAllPeers(addr)
	c.JSON(http.StatusOK, peerList)
}
