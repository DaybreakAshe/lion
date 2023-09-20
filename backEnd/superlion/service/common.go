//@program: superlion
//@author: yanjl
//@create: 2023-09-15 09:22
package service

import (
	"fmt"
	imgtype "github.com/shamsher31/goimgtype"
	"github.com/u2takey/go-utils/uuid"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"superlion/bean"
	"superlion/constants"
	"superlion/model"
	"superlion/util"
	"time"
)

// 文件上传并保存
func PictureUpload(sourceFile *multipart.File, file *multipart.FileHeader, busiType string, user *LionUserInfo) (*bean.FileRspBean, string) {

	if user == nil {
		return nil, "请先登录"
	}

	//util.UploadToB2()

	fileId := uuid.NewUUID()

	fileEntity := &model.TFilePic{
		FileId:       fileId,
		FileSize:     strconv.FormatInt(file.Size, 10),
		FileSizeByte: file.Size,
		RealName:     file.Filename,
		RowStat:      constants.STATUS_OK,
		ContentType:  path.Ext(file.Filename), // 获取文件后缀

		StorePath: constants.PICS_FILE_PATH + busiType + "/" + file.Filename,
	}

	fileName := file.Filename

	mapStr := uploadFileToSM(sourceFile, fileName)
	//if len(err) != 0 {
	//	return nil, err
	//}

	rsp := &bean.FileRspBean{
		FileId:    fileId,
		FileName:  fileName,
		FileS3Url: "",
		FileUrl:   mapStr,
	}
	//str, _ := json.Marshal(fileInfo)
	fmt.Printf("save file to local path:%s\nfile:%T\n", "nil", *fileEntity)
	// SavePicInfoToDB(fileInfo)

	return rsp, ""
}

func SavePicInfoToDB() {

}

func getDatePath() string {
	dateStr := time.Now().Format("YYYYMMDD")
	return fmt.Sprintf("%s%s/", constants.PICS_FILE_PATH, dateStr)
}

// 上传文件到NGINX
func uploadFileToSM(part *multipart.File, fileName string) string {

	// 获取图片的类型
	datatype, err2 := imgtype.Get(fileName)
	if err2 != nil {
		println(`不是图片文件`)
	} else {
		// 根据文件类型执行响应的操作
		println(`文件类型是`, datatype)
	}

	return util.UploadFileToNginx(part, fileName)
}

// 上传文件到服务器
func UploadResource(dir string, file *multipart.File, sourceFile *multipart.FileHeader) (*FileLocalInfo, string) {

	if len(dir) == 0 {
		dir = getDatePath()
	}

	//完整路径
	var savePath = dir + sourceFile.Filename

	// 打开目录
	localFileInfo, fileStatErr := os.Stat(dir)
	//目录不存在
	if fileStatErr != nil || !localFileInfo.IsDir() {
		//创建目录
		errByMkdirAllDir := os.MkdirAll(dir, 0755)
		if errByMkdirAllDir != nil {
			fmt.Printf("###[error]%s mkdir error.....", dir, errByMkdirAllDir.Error())
			return nil, "创建目录失败"
		}
	}
	////上传文件前 先删除该资源之前上传过的资源文件
	////（编辑-重新选择文件-需要先删除该资源之前上传过的资源文件）
	////该代码执行的条件----上传的名称是唯一的，否则会出现误删
	////获取文件的前缀
	//fileNameOnly := fileNameParam
	//deleteFileWithName(fileNameOnly, saveDir)
	//deleteFileWithName(fileNameOnly, model.WebConfig.ResourcePath+"/"+
	//  model.WebConfig.WebConvertToPath)

	out, err := os.Create(savePath)
	if err != nil {
		return nil, err.Error()
	}
	defer out.Close()
	_, err = io.Copy(out, *file)
	if err != nil {
		return nil, err.Error()
	}

	//没有错误的情况下
	info := &FileLocalInfo{
		SaveDir:  dir,
		SaveName: sourceFile.Filename,
		SavePath: savePath,
	}
	return info, ""
}

type FileLocalInfo struct {
	SaveDir  string
	SaveName string
	SavePath string
}
