//@program: superlion
//@author: yanjl
//@create: 2023-09-13 13:49
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/u2takey/go-utils/json"
	"log"
	"net/http"
	"superlion/constants"
	"superlion/service"
)

/**
获取解析token 的用户信息
*/
func GetLoginInfoByC(c *gin.Context) *service.LionUserInfo {

	luser, eor := c.Get(constants.LoginUser)
	if !eor {
		return nil
	}
	luserBean := &service.LionUserInfo{}

	err := json.Unmarshal([]byte(fmt.Sprintln(luser)), luserBean)

	if err != nil {
		log.Printf("get login json format error %s\n:", err.Error())
		return nil
	}
	return luserBean
}

/**
图片上传公共函数
*/
func PictureUpload(c *gin.Context) {

	busiType := c.PostForm("busiType")
	// c.Request.FormFile("picture")
	user := GetLoginInfoByC(c)
	file, eor := c.FormFile("picture")
	if eor != nil {
		fmt.Println("获取数据失败:\n", eor.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    608,
			"message": "获取数据失败",
		})
	} else {
		service.PictureUpload(nil, file, busiType, user)
	}

}
