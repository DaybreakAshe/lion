// @program: superlion
// @author: yanjl
// @description: 文件上传服务
// @create: 2024-08-15 09:46
package service

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"superlion/model"
	"superlion/repository"
	"superlion/util"
	"time"
)

const (
	filePathPres = "https://lion.yanxue.eu.org/lion/tmp/lion"
	fileUrlProxy = "https://64d0bb8.webp.ee/lion/tmp/lion"
)

var (
	fileUploadDao = repository.NewFileUploadDao()
)

// UploadToR2Aws 上传文件到R2的S3
func UploadToR2Aws(file multipart.File, header *multipart.FileHeader, busType string) {
	// 初始化 S3Client
	s3Client := util.NewS3Client()

	// 上传文件的路由
	defer file.Close()

	filename := filepath.Base(header.Filename)

	versionId := s3Client.UploadFile(file, busType+"/"+filename, header.Header.Get("Content-Type"))
	log.Print("complete to upload file to Cloudflare R2:" + versionId)

	// 保存到数据库
	go saveFileInfoToDB(header, busType, versionId)
	fmt.Sprintf("File %s uploaded successfully", filename)

}

func saveFileInfoToDB(header *multipart.FileHeader, busType string, versionId string) {

	fileName := header.Filename
	// fileUrl := fmt.Sprintf("%s/%s/%s", filePathPres, busType, fileName)
	fileProxyUrl := fmt.Sprintf("%s/%s/%s", fileUrlProxy, busType, fileName)
	now := time.Now()

	fileEntity := &model.LionFileUpload{
		// 雪花id
		FileID: util.GenerateID(),
		// ReferID:        "",
		BusinessID:   versionId,
		BusinessType: busType,
		StoreName:    fileName,
		RealName:     fileName,
		StorePath:    fmt.Sprintf("%s/%s", busType, fileName),
		FileSize:     strconv.FormatInt(header.Size, 10),
		FileSizeByte: header.Size,
		FileURL:      fileProxyUrl,
		ContentType:  header.Header.Get("Content-Type"),
		// PreviewFileID:  "",
		// Tags:           "",
		ValidDate: now,
		RowStat:   "00",
		// Stat:           "",
		// CreateUserID:   "",
		// CreateUserName: ,
		CreateDate: now,
		// UpdateUserID:   "",
		// UpdateUserName: "",
		UpdateDate: now,
	}

	err := fileUploadDao.SaveFileInfo(fileEntity)

	if err != nil {
		log.Print("error: Failed to save file info to database" + err.Error())
	}
}
