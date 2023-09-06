package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // 创建服务
	/* 处理请求 */
	server.GET("/hello", func(context *gin.Context) {
		context.JSONP(200, gin.H{"msg": "Hello World !"})
	})

	/* 处理请求 */
	server.GET("/", func(context *gin.Context) {
		context.JSONP(200, gin.H{"msg": "Hello go !"})
	})

	server.Run() // 监听并在 0.0.0.0:8080 上启动服务

	fmt.Printf("app start over")
}
