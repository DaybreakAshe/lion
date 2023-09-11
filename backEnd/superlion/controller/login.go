//@program: superlion
//@author: yanjl
//@create: 2023-09-07 14:47
package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"superlion/service"
	"superlion/util"
)

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
	util.PrintLog("login func :" + string(jstr))
	//	log.Panicf("receive body :%s\n", data)
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
	log.Panicf("receive body :%s\n", data)
}
