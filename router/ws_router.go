package router

import (
	"chat/config"
	"chat/service/ws"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	server := ws.GetServer()
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
		conn := ws.NewConnection(server, WsConn, connID)
		connID++
		// 开启读写线程
		go conn.Start()
		//conn.Start()

	})

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.GlobalConfig.App.IP, config.GlobalConfig.App.WebsocketPort),
		Handler: r,
	}
	go func() {
		fmt.Println("websocket 启动：", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	fmt.Println("ws_router starting...")

	//Set up channel on which to send signal notifications.
	quit := make(chan os.Signal, 1)
	//Notify方法很重要
	//causes package signal to relay incoming signals to c
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Block until a signal is received.
	s := <-quit
	fmt.Println("Got signal:", s)

	// 关闭服务
	//server.Stop()

}
