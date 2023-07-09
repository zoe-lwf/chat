package middlewares

import (
	"chat/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaim, err := util.AnalyseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证未通过",
			})
			return
		}
		c.Set("user_claims", userClaim)
		c.Next() //访问下一个Handler
	}
}
