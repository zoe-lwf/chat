package router

import (
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
	httpAddr := fmt.Sprintf("%s:%s", "127.0.0.1", "8080")
	fmt.Println(httpAddr)
	err := r.Run(httpAddr)

	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
