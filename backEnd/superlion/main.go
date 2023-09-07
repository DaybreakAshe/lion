package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

func main() {

	// 1、连接数据库：
	//err := repository.InitMysqlDB()
	//if err != nil {
	//	// mysql连接失败：
	//	fmt.Printf("mysql connect failed :{%s}...", err)
	//	os.Exit(-1)
	//}

	// 2、连接redis:

	// 、启动服务端
	server := gin.Default() // 创建服务
	server.Use(cors())
	/* 处理请求 */

	InitRouter(server)

	server.Run(":8080")

	fmt.Printf("app start over")
}

// 开启跨域函数
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 接收特定
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "https://superlion.vercel.app/")
		// 接收所有
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
