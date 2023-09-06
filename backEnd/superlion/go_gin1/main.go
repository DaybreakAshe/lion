package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

func main() {
	server := gin.Default() // 创建服务
	server.Use(Cors())
	/* 处理请求 */
	server.GET("/hello", func(context *gin.Context) {
		fmt.Println("access path:/hello")
		context.JSONP(200, gin.H{"msg": "Hello World !"})
	})

	/* 处理请求 */
	server.GET("/", func(context *gin.Context) {
		fmt.Println("access path:/")
		context.JSONP(200, gin.H{"msg": "Hello go !"})
	})

	server.Run() // 监听并在 0.0.0.0:8080 上启动服务

	fmt.Printf("app start over")
}

// 开启跨域函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Panic info is: %v\n", err)
				fmt.Printf("Panic info is: %s\n", debug.Stack())
			}
		}()

		c.Next()
	}
}
