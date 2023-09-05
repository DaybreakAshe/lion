package testController

import "github.com/gin-gonic/gin"

func test() {
	server := gin.Default() // 创建服务
	/* 处理请求 */
	server.GET("/", func(context *gin.Context) {
		context.JSONP(200, gin.H{"msg": "Hello World 2!"})
	})
	server.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func getTest() {
	server := gin.Default() // 创建服务
	/* 处理请求 */
	server.GET("/yanjl", func(context *gin.Context) {
		context.JSONP(200, gin.H{"msg": "Hello World 3!"})
	})
	server.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
