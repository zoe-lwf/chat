package util

import (
	"chat/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 签名
var sign = "lwf"

type UserClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

// 生成 token

func GenerateToken(userId uint64) (string, error) {
	//签名放在函数外面读取不到，配置文件读不到，hash的可以读取
	//var sign = config.GlobalConfig.JWT.SignKey
	UserClaim := &UserClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.GlobalConfig.JWT.ExpireTime))),
			Issuer:    "test", //签发人
		},
	}
	// 根据签名生成token，NewWithClaims(加密方式,claims) ==》 头部，载荷，签证
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString([]byte(sign))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token

func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	token, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GlobalConfig.JWT.SignKey), nil
	})
	//Valid用于校验鉴权声明。解析出载荷部分
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
