package controller

import (
	"github.com/gin-gonic/gin"
	"superlion/bean"
	"superlion/service"
)

func Index(c *gin.Context) {

	/* 处理请求 */
	c.JSONP(200, gin.H{"msg": "Hello index 2!"})
}

func Hello(c *gin.Context) {

	/* 处理请求 */
	c.JSONP(200, gin.H{"msg": "Hello World 2!"})
}

func Login(c *gin.Context) {

	req := bean.LoginReq{
		c.Param("name"),
		c.Param("passwd"),
	}
	/* 处理请求 */
	data, err := service.Login(&req)

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
}
