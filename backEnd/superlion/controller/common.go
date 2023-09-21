//@program: superlion
//@author: yanjl
//@create: 2023-09-13 13:49
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/u2takey/go-utils/json"
	"log"
	"mime/multipart"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"superlion/bean"
	"superlion/constants"
	"superlion/service"
	"superlion/util"
	"time"
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
	file, eor := c.FormFile("picture")
	// file := c.Request.MultipartForm.Value["picture"]
	// file, fheader, eor := c.Request.FormFile("picture")
	user := GetLoginInfoByC(c)
	// c = getFile(file, c)
	if eor != nil {
		fmt.Println("获取数据失败:\n", eor.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    608,
			"message": "获取数据失败",
		})
		return
	} else {
		fmt.Println("file:", file.Size)
		data, err := service.PictureUpload(nil, file, busiType, user)
		if len(err) != 0 {
			c.JSON(http.StatusOK, bean.CommonResponse{
				Data: data,
				Msg:  err,
				Code: 608,
			})
		} else {
			c.JSON(http.StatusOK, bean.CommonResponse{
				Data: data,
				Msg:  err,
				Code: 0,
			})
		}
		return
	}

}

func getFile(file *multipart.FileHeader, ctx *gin.Context) *gin.Context {
	//获取文件名称
	fmt.Println(file.Filename)
	//文件大小
	fmt.Println(file.Size)
	//获取文件的后缀名
	extstring := path.Ext(file.Filename)
	fmt.Println(extstring)
	//根据当前时间鹾生成一个新的文件名
	fileNameInt := time.Now().Unix()
	fileNameStr := strconv.FormatInt(fileNameInt, 10)
	//新的文件名
	fileName := fileNameStr + extstring
	//保存上传文件
	filePath := filepath.Join(util.Mkdir("upload"), "/", fileName)
	err := ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		fmt.Println("save file failed:", err.Error())
		return ctx
	}
	return ctx
}
