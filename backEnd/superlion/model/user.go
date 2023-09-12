//@program: superlion
//@author: yanjl
//@create: 2023-09-07 15:52
package model

import "time"

// 数据库映射实体结构体
type UserEntity struct {
	GoName    string
	LoginName string
	Avatar    string
	Status    string
	// 主键 todo
	GoId             string
	GoEmail          string
	GoToken          string
	GoVerified_email bool
	UserId           string
	GoPicture        string
	GoLocale         string
	CreateTime       time.Time
}
