package router

import (
	ws2 "chat/service/ws"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// websocket服务
var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connID uint64

func WSRouter() {
	server := ws2.GetServer()
	// 开启worker工作池
	server.StartWorkerPool()

	//TODO 心跳

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	r.GET("/ws", func(c *gin.Context) {
		// 升级协议  http -> websocket
		WsConn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("websocket conn err :", err)
			return
		}
		// 初始化连接
		conn := ws2.NewConnection(server, WsConn, connID)
		connID++
		// 开启读写线程
		go conn.Start()

	})
}
