//@program: superlion
//@author: yanjl
//@create: 2023-09-25 14:29
package controller

import (
	"github.com/gin-gonic/gin"
	"superlion/bean"
	"superlion/service"
)

var (
	userService = service.NewUserServiceInstance()
)

// 获取用户所属标签
func GetUserTags(c *gin.Context) {

	user := GetLoginInfoByC(c)

	data, err := userService.GetUserTagList(user.GoId)

	writeResponse(c, err, data)
}

// 用户新增标签
func CreateNewTag(c *gin.Context) {

	user := GetLoginInfoByC(c)

	req := &bean.TagReqBean{}
	eor := c.BindJSON(req)
	if eor != nil {
		c.JSONP(400, gin.H{
			"msg": eor.Error(),
		})
		return
	}
	data, err := userService.CreateNewTag(req, user)

	writeResponse(c, err, data)
}

// 用户新增文章类别
func AddNewArtType(c *gin.Context) {

	user := GetLoginInfoByC(c)

	req := &bean.ArticleTypeReq{}
	eor := c.BindJSON(req)
	if eor != nil {
		c.JSONP(400, gin.H{
			"msg": eor.Error(),
		})
		return
	}
	err := userService.AddNewArtType(req, user)

	writeResponse(c, err, nil)
}

// 用户文章类别列表
func GetUserArtTypeList(c *gin.Context) {

	user := GetLoginInfoByC(c)

	req := &bean.FindArtTypeReq{}

	data, err := userService.GetUserArtTypeList(req, user)

	writeResponse(c, err, data)
}

// 用户删除标签
func DeleteTag(c *gin.Context) {

	user := GetLoginInfoByC(c)

	req := map[string]string{}
	eor := c.BindJSON(&req)
	if eor != nil {
		c.JSONP(400, gin.H{
			"msg": eor.Error(),
		})
		return
	}
	data, err := userService.DeleteTag(req, user)

	writeResponse(c, err, data)
}

// 保存帖子草稿（缓存）
func SavePostCache(c *gin.Context) {
	user := GetLoginInfoByC(c)

	req := bean.PostCacheReq{}
	eor := c.BindJSON(&req)
	if eor != nil {
		c.JSONP(400, gin.H{
			"msg": eor.Error(),
		})
		return
	}
	err := userService.SavePostCache(&req, user)

	writeResponse(c, err, 0)
}

// 读取缓存详细
func GetMyCaches(c *gin.Context) {
	user := GetLoginInfoByC(c)

	req := map[string]any{}
	eor := c.BindJSON(&req)
	if eor != nil {
		c.JSONP(400, gin.H{
			"msg": eor.Error(),
		})
		return
	}
	data, err := userService.GetUserCacheList(req, user)

	writeResponse(c, err, data)
}
