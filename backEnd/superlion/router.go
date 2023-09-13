package main

import (
	"github.com/gin-gonic/gin"
	"superlion/config/webConfig"
	"superlion/controller"
)

func InitRouter(r *gin.Engine) {

	apiRouter := r.Group("/lion")

	apiRouter.GET("/", controller.Index)
	apiRouter.GET("/hello", controller.Hello)
	apiRouter.POST("/login1", controller.Login)
	apiRouter.POST("/login", controller.GetAuthParams)

	// 鉴权
	apiAuthRouter := r.Group("/auth")
	apiAuthRouter.Use(webConfig.LionTokenFilter())
	apiAuthRouter.POST("/user", controller.GetUserInfoByGId)

}
