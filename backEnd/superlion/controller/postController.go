//@program: superlion
//@author: yanjl
//@create: 2023-10-08 17:30
package controller

import (
	"github.com/gin-gonic/gin"
	"superlion/service"
	"superlion/util"
)

var (
	postService = service.GetPostServiceInstance()
)

// GetPostContent 获取文章内容 /article/{id}
func GetPostContent(c *gin.Context) {

	id := c.Param("id")

	data, err := postService.FindPostById(util.StrToInt(id))

	writeResponse(c, err, data)
}

// GetMyPostList 获取文章内容 /article/{id}
func GetMyPostList(c *gin.Context) {

	user := GetLoginInfoByC(c)

	data, err := postService.GetMyPostList(user)

	writeResponse(c, err, data)
}
