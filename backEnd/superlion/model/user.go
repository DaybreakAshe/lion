//@program: superlion
//@author: yanjl
//@create: 2023-09-07 15:52
package model

import "time"

// 数据库映射实体结构体
type UserEntity struct {
	GoName    string `json:"GoName,omitempty" gorm:"column:go_name"`
	LoginName string `json:"LoginName,omitempty" gorm:"column:login_name"`
	Avatar    string `json:"Avatar,omitempty" gorm:"column:avatar"`
	Status    string `json:"Status,omitempty" gorm:"status"`
	// 主键 todo
	GoId            string    `json:"GoId" gorm:"primaryKey column:go_id"`
	GoEmail         string    `json:"GoEmail" json:"GoEmail,omitempty"`
	GoToken         string    `json:"GoToken,omitempty" gorm:"column:go_token"`
	GoVerifiedEmail bool      `json:"GoVerified_Email,omitempty" gorm:"column:go_verified_email"`
	UserId          string    `json:"UserId,omitempty" gorm:"column:user_id"`
	GoPicture       string    `json:"GoPicture,omitempty" gorm:"column:go_picture"`
	GoLocale        string    `json:"GoLocale,omitempty" gorm:"column:go_locale"`
	CreateTime      time.Time `json:"CreateTime" gorm:"column:create_time"`
	Signature       string    `json:"signature" gorm:"column:signature"`
}

func (UserEntity) TableName() string {
	return "lion_user"
}
