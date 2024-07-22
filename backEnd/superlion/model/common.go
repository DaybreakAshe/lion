//@program: superlion
//@author: yanjl
//@create: 2023-09-15 09:51
package model

import "time"

// 文件-上传管理表
type TFilePic struct {
	FileId       string    `gorm:"column:file_id" json:"file_id"`               //文件编号
	RealName     string    `gorm:"column:real_name" json:"real_name"`           //源文件名
	StorePath    string    `gorm:"column:store_path" json:"store_path"`         //保存路径
	FileSize     string    `gorm:"column:file_size" json:"file_size"`           //文件大小
	FileSizeByte int64     `gorm:"column:file_size_byte" json:"file_size_byte"` //文件大小(字节)
	FileUrl      string    `gorm:"column:file_url" json:"file_url"`             //文件地址
	B2FileId     string    `gorm:"column:b2_file_id" json:"b2_file_id"`         //b2文件ID
	RowStat      string    `gorm:"column:row_stat" json:"row_stat"`             //业务状态
	CreateUserId string    `gorm:"column:create_user_id" json:"create_user_id"` //创建者 用户Id
	CreateDate   time.Time `gorm:"column:create_date" json:"create_date"`       //创建时间
	ContentType  string    `gorm:"column:content_type" json:"content_type"`     //文件类型
	ContentSha1  string    `gorm:"column:content_sha1" json:"content_sha1"`     //sha1密文
}

func (TFilePic) TableName() string {
	return "t_file_pic"
}

type PageDto struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}
