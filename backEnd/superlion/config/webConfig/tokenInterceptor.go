//@program: superlion
//@author: yanjl
//@create: 2023-09-12 20:43
package webConfig

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/u2takey/go-utils/json"
	"log"
	"strings"
	"superlion/config"
	"superlion/service"
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
		if len(token) <= 7 || !strings.HasPrefix(token, "Bearer ") {
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
			lUser := &service.LionUserInfo{}
			eor := json.Unmarshal([]byte(userJson), lUser)
			log.Printf("json format error :%s\n", eor)
			c.Set("lUser", lUser)
		}
	}
}

func checkToken(token string) string {

	redisP := config.GetRedisHelper()

	token = RedisPre + token

	data, _ := redisP.Get(ctx, token).Result()

	return data
}
