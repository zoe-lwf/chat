package service

import (
	"chat/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	// 获取参数并验证
	phoneNumber := c.PostForm("phone_number")
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")

	if phoneNumber == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "参数不正确",
		})
		return
	}
	cnt, err := model.GetUserCountByPhone(phoneNumber)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号已被注册",
		})
		return
	}
	// 插入用户信息
	ub := model.User{
		PhoneNumber: phoneNumber,
		Nickname:    nickname,
		Password:    password, //TODO 加密，加盐
	}
	err = model.CreateUser(&ub)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误" + err.Error(),
		})
		return
	}

	// 生成 token
	// 发放 token
}
