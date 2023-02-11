package ws

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/onee-only/mempool-manager/lib"
)

var wsUpgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

func UpgradeWS(c *gin.Context) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	lib.HandleErr(err)

	fmt.Println(c.Request.RemoteAddr)
	fmt.Println(c.Request.Host)
	conn.Close()
}
