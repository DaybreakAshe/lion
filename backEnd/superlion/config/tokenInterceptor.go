//@program: superlion
//@author: yanjl
//@create: 2023-09-12 20:43
package config

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"superlion/constants"
)

const (
	Header   = "Authorization"
	Bearer   = "Bearer"
	RedisPre = "TOKEN:"
)

var (
	ctx = context.Background()
)

func LionTokenFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(Header)
		if len(token) <= 7 || !strings.HasPrefix(token, Bearer) {
			c.JSON(401, gin.H{
				"code": "401",
				"msg":  "#您登录了吗？",
			})
			c.Abort()
			return
		}
		token = string([]rune(token)[7:])
		log.Printf("user token :[%s]\n", token)
		if len(token) == 0 {
			c.JSON(401, gin.H{
				"code": "401",
				"msg":  "@您登录了吗？",
			})
			c.Abort()
			return
		}
		userJson := checkToken(token)
		if len(userJson) <= 0 {
			c.JSON(402, gin.H{
				"code": "402",
				"msg":  "登录过期了",
			})
			c.Abort()
			return
		} else {
			// 用户缓存存在：
			log.Printf("login user : %s\n", userJson)
			c.Set(constants.LoginUser, userJson)
		}
	}
}

func checkToken(token string) string {

	redisP := GetRedisHelper()

	token = RedisPre + token

	data, _ := redisP.Get(ctx, token).Result()

	return data
}
