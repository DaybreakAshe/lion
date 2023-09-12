package main

import (
	"github.com/gin-gonic/gin"
	"superlion/controller"
)

func InitRouter(r *gin.Engine) {

	apiRouter := r.Group("/lion")

	apiRouter.GET("/", controller.Index)
	apiRouter.GET("/hello", controller.Hello)
	apiRouter.POST("/login1", controller.Login)
	apiRouter.POST("/login", controller.GetAuthParams)
	apiRouter.POST("/user", controller.GetUserInfoByGId)
}
