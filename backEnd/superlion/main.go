package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"superlion/config"
	"superlion/repository"
	"superlion/router"
)

var (
	ctx = context.Background()
)

func main() {

	// uploadFile()
	server()
	// util.Upload()

	fmt.Println("app start ....")
}

func server() {
	// 1、连接数据库：
	err := repository.InitMysqlDB()
	if err != nil {
		// mysql连接失败：
		fmt.Printf("mysql connect failed :{%s}...", err)
		os.Exit(-1)
	}

	// 2、连接redis:
	rdb := config.NewRedisHelper()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("redis connect failed :{%s}...", err.Error())
		return
	}

	// 、启动服务端
	server := gin.Default() // 创建服务
	server.Use(cors())

	router.InitRouter(server)

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
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		//// 统一error拦截？
		//defer func() {
		//	if r := recover(); r != nil {
		//		//打印错误堆栈信息
		//		log.Printf("panic: %v\n", r)
		//		debug.PrintStack()
		//		//封装通用json返回
		//		//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
		//		//Result.Fail不是本例的重点，因此用下面代码代替
		//		c.JSON(http.StatusOK, gin.H{
		//			"code": "608",
		//			"msg":  errorToString(r),
		//			"data": nil,
		//		})
		//		//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
		//		c.Abort()
		//	}
		//}()
		c.Next()
	}
}

func uploadFile() {

	//str := "C:\\Users\\Yanjilong\\Desktop\\LinkinPark\\Lucian_25.jpg"+"Lucian.jpg";
	// util.UploadPicToImagse()

}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用json返回
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail不是本例的重点，因此用下面代码代替
			c.JSON(http.StatusOK, gin.H{
				"code": "608",
				"msg":  errorToString(r),
				"data": nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
