package controller

import (
	"github.com/gin-gonic/gin"
	"superlion/bean"
	"superlion/service"
)

/**
首页测试
*/
func Index(c *gin.Context) {

	/* 处理请求 */
	c.JSONP(200, gin.H{"msg": "Hello index 2!"})
}

/**
测试接口
*/
func Hello(c *gin.Context) {

	/* 处理请求 */
	c.JSONP(200, gin.H{"msg": "Hello World 2!"})
}

/**
登录测试
*/
func Login(c *gin.Context) {

	req := bean.LoginReq{
		c.PostForm("name"),
		c.PostForm("passwd"),
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
