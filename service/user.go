package service

import (
	"chat/model"
	"chat/pkg/db"
	"chat/util"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	//查询
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
		Password:    util.GetMD5(password), //TODO 加密，加盐
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
	token, err := util.GenerateToken(ub.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}
	// 发放 token
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"data": gin.H{
			"token": token,
			"id":    util.Uint64ToStr(ub.ID),
		},
	})
}

func Login(c *gin.Context) {
	phoneNumber := c.PostForm("phone_number")
	password := c.PostForm("password")
	if phoneNumber == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "参数不正确",
		})
		return
	}
	// 验证账号名和密码是否正确
	ub, err := model.GetUserFromPhoneAndPWD(phoneNumber, util.GetMD5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "手机号或密码错误",
		})
		return
	}
	// 生成 token
	token, err := util.GenerateToken(ub.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}

	// 发放 token
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token":   token,
			"user_id": util.Uint64ToStr(ub.ID),
		},
	})
}

// 防止跨域站点伪造请求
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

// SendMsg 发送消息
func SendMsg(c *gin.Context) {
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer func(ws *websocket.Conn) {
		ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	MsgHandler(ws)
}

func MsgHandler(ws *websocket.Conn) {
	_, p, _ := ws.ReadMessage()
	msg, err := db.RDB.Set(context.Background(), db.PublicKey, string(p), 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)
	//msg, err := db.Subscribe(c, db.PublicKey)
	//tm := time.Now().Format("2006-01-02 15:04:05")
	//m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
	m, _ := db.RDB.Get(context.Background(), db.PublicKey).Result()
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println(err)
	}
}
