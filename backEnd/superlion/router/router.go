package router

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
	// 无需认证//
	// 获取文章（草稿？）
	apiRouter.GET("/post/:id", controller.GetPostContent)
	// 获取文章列表
	apiRouter.GET("/posts", controller.GetPostList)

	// 鉴权
	apiAuthRouter := r.Group("/auth")
	apiAuthRouter.Use(config.LionTokenFilter())
	apiAuthRouter.POST("/user", controller.GetUserInfoByGId)
	apiAuthRouter.POST("/user/edit", controller.UpdateUserInfo)
	apiAuthRouter.POST("/upload", controller.PictureUpload)

	// 文章:
	apiUserRouter := r.Group("/user")
	apiUserRouter.Use(config.LionTokenFilter())
	apiUserRouter.POST("/newTag", controller.CreateNewTag)
	apiUserRouter.GET("/tags", controller.GetUserTags)
	apiUserRouter.POST("/delTag", controller.DeleteTag)
	apiUserRouter.POST("/saveCache", controller.SavePostCache)
	apiUserRouter.GET("/cache", controller.GetMyCaches)

	// 新增文章类型
	apiUserRouter.POST("/newType", controller.AddNewArtType)
	// 文章类型
	apiUserRouter.GET("/types", controller.GetUserArtTypeList)
	// 获取我的文章列表
	apiUserRouter.GET("/posts", controller.GetMyPostList)
	// 发布文章
	apiUserRouter.POST("/publish", controller.PublishPost)

	// 通用接口
	// 文件上传
	commonRouter := r.Group("/common")
	commonRouter.POST("/upload", controller.PictureUploadR2)
	// 图片预览

}
