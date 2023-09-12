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

	gid := c.Param("gid")

	data, err := service.GetUserInfoByGoId(gid)
	if err != nil {
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
