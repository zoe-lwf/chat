package router

import (
	"chat/config"
	"chat/pkg/middlewares"
	"chat/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HTTPRouter() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	//用户注册
	r.POST("/register", service.Register)
	//用户登录
	r.POST("/login", service.Login)
	//加入group的所有请求都要鉴权
	auth := r.Group("/u", middlewares.AuthCheck()) //中间件校验权限
	{
		//添加好友
		auth.POST("/friend/add", service.AddFriend)
		// 创建群聊
		auth.POST("/group/create", service.CreateGroup)
		//获取群成员列表
		//auth.POST("/group_user/list", service.GroupUserList)
	}
	//发送消息
	r.GET("/user/sendMsg", service.SendMsg)

	httpAddr := fmt.Sprintf("%s:%s", config.GlobalConfig.App.IP, config.GlobalConfig.App.HTTPServerPort)
	fmt.Println(httpAddr)
	err := r.Run(httpAddr)

	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	fmt.Println("http router starting...")

}
