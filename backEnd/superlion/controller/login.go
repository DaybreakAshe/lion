//@program: superlion
//@author: yanjl
//@create: 2023-09-07 14:47
package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"superlion/bean"
	"superlion/service"
)

// GetAuthParams 获取谷歌用户信息
func GetAuthParams(c *gin.Context) {

	req := service.LoginParmas{}

	// 获取body参数
	eor := c.BindJSON(&req)
	if eor != nil {
		c.JSONP(400, gin.H{
			"msg": eor.Error,
		})
	}

	// 解析：
	data, err := service.GetGoogleAuthBody(req)

	jstr, eor := json.Marshal(*data)
	fmt.Printf("login func :%s\n", string(jstr))
	//	log.Panicf("receive body :%s\n", data)
	if len(err) != 0 {
		c.JSONP(400, gin.H{
			"msg": err,
		})
	} else {
		c.JSONP(200, gin.H{
			"msg":  "",
			"data": *data,
		})
	}
	fmt.Printf("receive body :%s\n", *data)
}

/**
根据传入gid查询用户
*/
func GetUserInfoByGId(c *gin.Context) {

	gid := c.Query("gid")

	data, err := service.GetUserInfoByGoId(gid)
	if err == nil {
		c.JSON(http.StatusOK, bean.CommonResponse{
			Code: 200,
			Data: *data,
		})
	} else {
		c.JSON(http.StatusOK, bean.CommonResponse{
			Code: 608,
			Msg:  err.Error(),
		})
	}

}

/**
用户更新头像或昵称
*/
func UpdateUserInfo(c *gin.Context) {

	lUser, _ := c.Get("lUser")

	luserBean := GetLoginInfoByC(lUser)

	req := UpdateUserInfoBean{}
	eor := c.BindJSON(&req)
	if eor != nil {
		c.JSONP(400, gin.H{
			"msg": eor.Error,
		})
	}
	data := service.UpdateUserInfo(luserBean, req.nickName, req.avatar)

	// 失败
	if len(data) != 0 {
		c.JSON(608, bean.CommonResponse{
			Code: 608,
			Msg:  data,
		})
	} else {
		c.JSON(200, bean.CommonResponse{
			Code: 200,
			Msg:  "修改成功",
		})
	}
}

/**
修改信息请求bean
*/
type UpdateUserInfoBean struct {
	nickName string
	avatar   string
}
