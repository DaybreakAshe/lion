//@program: superlion
//@author: yanjl
//@create: 2023-09-07 14:47
package controller

import (
	"github.com/gin-gonic/gin"
	"superlion/service"
)

func GetAuthParams(c *gin.Context) {

	req := service.LoginParmas{}

	// 获取body参数
	eor := c.BindJSON(&req)
	if eor != nil {
		return
	}

	// 解析：
	data, err := service.GetGoogleAuthBody(req)

	if len(err) != 0 {
		c.JSONP(400, gin.H{
			"msg": err,
		})
	} else {
		c.JSONP(200, gin.H{
			"msg":  "",
			"data": data,
		})
	}
	return
}
