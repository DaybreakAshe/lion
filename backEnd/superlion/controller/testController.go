package controller

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {

	/* 处理请求 */
	c.JSONP(200, gin.H{"msg": "Hello index 2!"})
}

func Hello(c *gin.Context) {

	/* 处理请求 */
	c.JSONP(200, gin.H{"msg": "Hello World 2!"})
}
