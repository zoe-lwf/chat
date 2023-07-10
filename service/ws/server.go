package ws

import (
	"chat/config"
	"fmt"
	"sync"
)

var (
	ConnManager *Server
	once        sync.Once
)

// Server 连接管理
type Server struct {
	ConnMap   *sync.Map   // 并发安全的 map,登录的用户连接 k-用户userid v-连接
	taskQueue []chan *Req //工作池
}

func GetServer() *Server {
	//once.Do的作用：无论如何更换Do中的方法,保证once只被sync(同步)执行一次
	once.Do(func() {
		ConnManager = &Server{
			// 初始worker队列中，worker个数
			taskQueue: make([]chan *Req, config.GlobalConfig.App.WorkerPoolSize),
		}
	})
	return ConnManager
}

// StartWorkerPool 开启工作池
func (cm *Server) StartWorkerPool() {
	// 初始化并启动 worker 工作池
	for i := 0; i < len(cm.taskQueue); i++ {
		// 初始化worker队列中，每个worker的队列长度
		cm.taskQueue[i] = make(chan *Req, config.GlobalConfig.App.MaxWorkerTask)
		cm.StartOneWorker(i, cm.taskQueue[i])
	}
}

// StartOneWorker 启动 worker 的工作流程
func (cm *Server) StartOneWorker(workId int, taskQueue chan *Req) {
	fmt.Println("Worker ID = ", workId, " is started.")
	for {
		//随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的
		select { //类似于switch，但是主要用于chan
		case req := <-taskQueue: //req接受chan数据
			req.f() //接收到数据，执行handler
		}
	}
}

// RemoveConn 删除连接
func (cm *Server) RemoveConn(userId uint64) {
	cm.ConnMap.Delete(userId)
	fmt.Printf("connection UserId=%d remove from Server\n", userId)
}
