package main

import (
	"github.com/gin-gonic/gin"
	"superlion/config"
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
	apiAuthRouter.Use(config.LionTokenFilter())
	apiAuthRouter.POST("/user", controller.GetUserInfoByGId)
	apiAuthRouter.POST("/user/edit", controller.UpdateUserInfo)
	apiRouter.POST("/upload", controller.PictureUpload)

}
