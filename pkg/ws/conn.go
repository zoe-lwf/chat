package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

// Conn 连接实例
// 1. 启动读写线程
// 2. 读线程读到数据后，根据数据类型获取处理函数，交给 worker 队列调度执行
type Conn struct {
	ConnId           uint64          // 连接编号，通过对编号取余，能够让 Conn 始终进入同一个 worker，保持有序性
	server           *Server         // 当前连接属于哪个 server
	UserId           uint64          // 连接所属用户id
	UserIdMutex      sync.RWMutex    // 保护 userId 的锁
	Socket           *websocket.Conn // 用户连接
	sendCh           chan []byte     // 用户要发送的数据
	isClose          bool            // 连接状态
	isCloseMutex     sync.RWMutex    // 保护 isClose 的锁
	exitCh           chan struct{}   // 通知 writer 退出
	maxClientId      uint64          // 该连接收到的最大 clientId，确保消息的可靠性
	maxClientIdMutex sync.Mutex      // 保护 maxClientId 的锁

	lastHeartBeatTime time.Time  // 最后活跃时间
	heartMutex        sync.Mutex // 保护最后活跃时间的锁
}

func NewConnection(server *Server, wsConn *websocket.Conn, ConnId uint64) *Conn {
	return &Conn{
		ConnId:            ConnId,
		server:            server,
		UserId:            0, // 此时用户未登录， userID 为 0
		Socket:            wsConn,
		sendCh:            make(chan []byte, 10),
		isClose:           false,
		exitCh:            make(chan struct{}, 1),
		lastHeartBeatTime: time.Now(), // 刚连接时初始化，避免正好遇到清理执行，如果连接没有后续操作，将会在下次被心跳检测踢出
	}
}

func (c *Conn) Start() {
	// 开启从客户端读取数据流程的 goroutine
	go c.StartReader()
}

// StartReader 用于从客户端中读取数据
func (c *Conn) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println(c.RemoteAddr(), "[conn Reader exit!]")
	defer c.Stop()
	for {
		// 阻塞读
		_, data, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println("read msg data error ", err)
			return
		}

		// 消息处理
		c.HandlerMessage(data)
	}
}

func (c *Conn) HandlerMessage(data []byte) {
	//  所有错误都需要写回给客户端
	// 消息解析 proto string -> struct

}

func (c *Conn) Stop() {
	c.isCloseMutex.Lock()
	defer c.isCloseMutex.Unlock()
	if c.isClose {
		return
	}
	// 关闭 socket 连接
	c.Socket.Close()
	// 关闭 writer
	c.exitCh <- struct{}{}
	if c.GetUserId() != 0 {
		c.server.RemoveConn(c.GetUserId())
		// 用户下线
		//_ = cache.DelUserOnline(c.GetUserId())
	}
	c.isClose = true
	// 关闭管道
	close(c.exitCh)
	close(c.sendCh)
	fmt.Println("Conn Stop() ... UserId = ", c.GetUserId())
}

// GetUserId 获取 userId
func (c *Conn) GetUserId() uint64 {
	c.UserIdMutex.RLock()
	defer c.UserIdMutex.RUnlock()

	return c.UserId
}

// RemoteAddr 获取远程客户端地址
func (c *Conn) RemoteAddr() string {
	return c.Socket.RemoteAddr().String()
}
