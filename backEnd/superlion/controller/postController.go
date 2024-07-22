//@program: superlion
//@author: yanjl
//@create: 2023-10-08 17:30
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
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

// GetMyPostList 获取我的文章列表 /posts/
func GetMyPostList(c *gin.Context) {

	user := GetLoginInfoByC(c)

	params := &bean.PostListParams{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		writeResponse(c, err.Error(), nil)
	}
	data, esr := postService.GetMyPostList(user, params)

	writeResponse(c, esr, data)
}

// GetMyPostList 获取文章列表 /posts/
func GetPostList(c *gin.Context) {

	// user := GetLoginInfoByC(c)

	params := &bean.PostListParams{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		writeResponse(c, err.Error(), nil)
	}
	str, _ := json.Marshal(params)
	log.Printf("get request params:%s", str)
	data, esr := postService.GetPostList(params)

	writeResponse(c, esr, data)
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
