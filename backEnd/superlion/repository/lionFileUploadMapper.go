// @program: superlion
// @author: yanjl
// @description: superlion
// @create: 2024-08-15 14:54
package repository

import (
	"superlion/model"
	"sync"
)

type FileUploadDao struct {
}

var (
	fileUploadDao     *FileUploadDao
	fileUploadDaoOnce sync.Once
)

// GetFileUploadDao 获取文件上传Dao
func NewFileUploadDao() *FileUploadDao {
	fileUploadDaoOnce.Do(func() {
		fileUploadDao = &FileUploadDao{}
	})
	return fileUploadDao
}

// 保存文件信息
func (*FileUploadDao) SaveFileInfo(fileInfo *model.LionFileUpload) error {

	// 保存文件信息到数据库
	err := db.Create(fileInfo).Error

	return err
}
