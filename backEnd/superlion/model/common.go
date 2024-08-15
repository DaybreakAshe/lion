//@program: superlion
//@author: yanjl
//@create: 2023-09-15 09:51
package model

import "time"

// 文件-上传管理表 过时，请勿使用
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

// LionFileUpload 文件-上传管理表
type LionFileUpload struct {
	FileID         string    `gorm:"column:file_id;primaryKey;size:32" json:"fileId"`       // 文件编号
	ReferID        string    `gorm:"column:refer_id;size:50" json:"referId"`                // 对象所属编号
	BusinessID     string    `gorm:"column:business_id;size:50" json:"businessId"`          // 业务编号
	BusinessType   string    `gorm:"column:business_type;size:4" json:"businessType"`       // 业务类型
	StoreName      string    `gorm:"column:store_name;size:100" json:"storeName"`           // 保存文件名
	RealName       string    `gorm:"column:real_name;size:100" json:"realName"`             // 源文件名
	StorePath      string    `gorm:"column:store_path;size:100" json:"storePath"`           // 保存路径
	FileSize       string    `gorm:"column:file_size;size:20" json:"fileSize"`              // 文件大小
	FileSizeByte   int64     `gorm:"column:file_size_byte" json:"fileSizeByte"`             // 文件大小(字节)
	FileURL        string    `gorm:"column:file_url;size:128" json:"fileUrl"`               // 文件地址
	ContentType    string    `gorm:"column:content_type;size:100" json:"contentType"`       // 内容类型
	PreviewFileID  string    `gorm:"column:preview_file_id;size:32" json:"previewFileId"`   // 预览文件ID
	Tags           string    `gorm:"column:tags;size:200" json:"tags"`                      // 有效时间
	ValidDate      time.Time `gorm:"column:valid_date" json:"validDate"`                    // 创建时间
	RowStat        string    `gorm:"column:row_stat;size:2" json:"rowStat"`                 // 业务状态
	Stat           string    `gorm:"column:stat;size:2" json:"stat"`                        // 记录状态
	Version        int       `gorm:"column:version" json:"version"`                         // 版本号
	CreateUserID   string    `gorm:"column:create_user_id;size:32" json:"createUserId"`     // 创建者 用户Id
	CreateUserName string    `gorm:"column:create_user_name;size:30" json:"createUserName"` // 创建者名称 真实名称
	CreateDate     time.Time `gorm:"column:create_date" json:"createDate"`                  // 创建时间
	UpdateUserID   string    `gorm:"column:update_user_id;size:32" json:"updateUserId"`     // 更新者编号 用户Id
	UpdateUserName string    `gorm:"column:update_user_name;size:30" json:"updateUserName"` // 更新者名称 真实名称
	UpdateDate     time.Time `gorm:"column:update_date" json:"updateDate"`                  // 更新时间
}

// TableName sets the insert table name for this struct type
func (LionFileUpload) TableName() string {
	return "lion_file_upload"
}
