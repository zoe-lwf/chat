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
	connMap   *sync.Map   // 并发安全的 map,登录的用户连接 k-用户userid v-连接
	taskQueue []chan *Req //工作池
}

func GetServer() *Server {
	fmt.Println("GetServer")
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
		// 启动worker,启动一个协程去处理，不然会阻塞主线程
		go cm.StartOneWorker(i, cm.taskQueue[i])
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
	cm.connMap.Delete(userId)
	fmt.Printf("connection UserId=%d remove from Server\n", userId)
}

// AddConn 添加连接
func (cm *Server) AddConn(userId uint64, conn *Conn) {
	cm.connMap.Store(userId, conn)
	fmt.Printf("connection UserId=%d add to Server\n", userId)
}

// GetConn 根据userid获取相应的连接
func (cm *Server) GetConn(userId uint64) *Conn {
	value, ok := cm.connMap.Load(userId)
	if ok {
		return value.(*Conn)
	}
	return nil
}

// GetConnAll 获取全部连接
func (cm *Server) GetConnAll() []*Conn {
	conns := make([]*Conn, 0)
	cm.connMap.Range(func(key, value interface{}) bool {
		conn := value.(*Conn)
		conns = append(conns, conn)
		return true
	})
	return conns
}

// Stop 关闭服务
func (cm *Server) Stop() {
	fmt.Println("server stop ...")
	ch := make(chan struct{}, 1000) // 控制并发数
	var wg sync.WaitGroup
	connAll := cm.GetConnAll()
	for _, conn := range connAll {
		ch <- struct{}{}
		wg.Add(1)
		c := conn
		go func() {
			defer func() {
				wg.Done()
				<-ch
			}()
			c.Stop()
		}()
	}
	close(ch)
	wg.Wait()
}

// SendMsgToTaskQueue 将消息交给 taskQueue，由 worker 调度处理
func (cm *Server) SendMsgToTaskQueue(req *Req) {
	if len(cm.taskQueue) > 0 {
		// 根据ConnID来分配当前的连接应该由哪个worker负责处理，保证同一个连接的消息处理串行
		// 轮询的平均分配法则
		//得到需要处理此条连接的workerID
		workerID := req.conn.ConnId % uint64(len(cm.taskQueue))
		// 将消息发给对应的 taskQueue
		cm.taskQueue[workerID] <- req
	} else {
		// 可能导致消息乱序
		go req.f()
	}

}
