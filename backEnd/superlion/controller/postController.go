//@program: superlion
//@author: yanjl
//@create: 2023-10-08 17:30
package controller

import (
	"github.com/gin-gonic/gin"
	"superlion/bean"
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

// 发布文章（草稿）
func PublishPost(c *gin.Context) {

	user := GetLoginInfoByC(c)
	req := &bean.SavePostReq{}
	eor := c.BindJSON(req)
	if eor != nil {
		writeResponse(c, eor.Error(), nil)
	}

	data, err := postService.PublishPost(req, user)

	writeResponse(c, err, data)
}
