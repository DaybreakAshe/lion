package main

import (
	"github.com/gin-gonic/gin"
	"superlion/controller"
)

func InitRouter(r *gin.Engine) {

	apiRouter := r.Group("/lion")

	apiRouter.GET("/", controller.Index)
	apiRouter.GET("/hello", controller.Hello)
}
