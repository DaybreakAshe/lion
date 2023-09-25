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

var (
	loginService = service.NewLoginServiceInstance()
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
	data, err := loginService.GetGoogleAuthBody(req)

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

	data, err := loginService.GetUserInfoByGoId(gid)
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

	luserBean := GetLoginInfoByC(c)

	req := &bean.UpdateUserInfoBean{}
	eor := c.BindJSON(req)
	if eor != nil {
		c.JSONP(400, gin.H{
			"msg": eor.Error(),
		})
		return
	}
	//fmt.Printf("edit infos,n:[%s],a:[%s]\n", req.)
	data := loginService.UpdateUserInfo(luserBean, req)
	writeResponse(c, data, nil)
}
