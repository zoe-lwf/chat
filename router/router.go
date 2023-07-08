package router

import (
	"chat/config"
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
	httpAddr := fmt.Sprintf("%s:%s", config.GlobalConfig.App.IP, config.GlobalConfig.App.HTTPServerPort)
	fmt.Println(httpAddr)
	err := r.Run(httpAddr)

	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
