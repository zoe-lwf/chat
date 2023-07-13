package ws

import (
	"chat/common"
	"chat/config"
	"chat/lib/cache"
	"chat/protocol/pb"
	"chat/util"
	"fmt"
	"google.golang.org/protobuf/proto"
)

// Handler 路由函数
type Handler func()

// Req 请求
type Req struct {
	conn *Conn   // 连接
	data []byte  // 客户端发送的请求数据
	f    Handler // 该请求需要执行的路由函数
}

func (r *Req) Login() {
	// 检查用户是否已登录 防止同一个连接多次调用 Login
	if r.conn.GetUserId() != 0 {
		fmt.Println("[用户登录] 用户已登录")
		return
	}
	// 消息解析 proto string -> struct
	loginMsg := new(pb.LoginMsg)
	//string -> struct
	err := proto.Unmarshal(r.data, loginMsg)
	if err != nil {
		fmt.Println("[用户登录] unmarshal error,err:", err)
		return
	}
	// 登录校验
	userClaims, err := util.AnalyseToken(string(loginMsg.Token))
	if err != nil {
		fmt.Println("[用户登录] AnalyseToken err:", err)
		return
	}
	// 检查用户是否已经在其他连接登录
	onlineAddr, err := cache.GetUserOnline(userClaims.UserId)
	if onlineAddr != "" {
		// TODO 更友好的提示
		fmt.Println("[用户登录] 用户已经在其他连接登录")
		r.conn.Stop()
		return
	}
	// Redis 存储用户数据 k: userId,  v: grpc地址，方便用户能直接通过这个地址进行 rpc 方法调用
	grpcServerAddr := fmt.Sprintf("%s:%s", config.GlobalConfig.App.IP, config.GlobalConfig.App.RPCPort)
	err = cache.SetUserOnline(userClaims.UserId, grpcServerAddr)
	if err != nil {
		fmt.Println("[用户登录] 系统错误")
		return
	}

	// 设置 user_id
	r.conn.SetUserId(userClaims.UserId)

	// 加入到 connMap 中
	r.conn.server.AddConn(userClaims.UserId, r.conn)

	// 回复ACK
	bytes, err := GetOutputMsg(pb.CmdType_CT_ACK, int32(common.OK), &pb.ACKMsg{Type: pb.ACKType_AT_Login})
	if err != nil {
		fmt.Println("[用户登录] proto.Marshal err:", err)
		return
	}

	// 回复发送 Login 请求的客户端
	r.conn.SendMsg(userClaims.UserId, bytes)

}
