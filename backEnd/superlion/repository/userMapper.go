//@program: superlion
//@author: yanjl
//@create: 2023-09-07 15:48
package repository

import (
	"fmt"
	"log"
	"superlion/model"
	"sync"
	"time"
)

// 数据库映射实体结构体
type UserEntityb struct {
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
}

type UserDao struct {
}

//func (model.UserEntity) TableName() string {
//	return "lion_user"
//}

var userDao *UserDao

/**在 Do 方法被调用后，该函数将被执行，而且只会执行一次，即使在多个协程同时调用的情况下也是如此*/
var userOnce sync.Once

// NewUserDaoInstance 单例构建Dao
func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

/**
通过googleId查询用户
*/
func (*UserDao) GetUserInfoByGId(gid string) (*model.UserEntity, string) {

	if len(gid) == 0 {
		return nil, "id不可以为空！"
	}

	user := &model.UserEntity{}

	err := db.Where("go_id = ?", gid).Take(user).Error

	if err != nil {
		fmt.Printf("query user from db failed:%s\n", err.Error())
		return nil, "query db error。"
	}
	return user, ""
}

/**
保存用户信息
*/
func (*UserDao) SaveUerInfoToDB(user *model.UserEntity) (int, string) {

	if len(user.GoId) == 0 {
		return 0, "gid不可以为空！"
	}

	// 插入数据
	err := db.Create(user).Error
	if err != nil {
		log.Printf("save user to db failed:%s\n", err.Error())
		return 0, err.Error()
	}
	return 1, ""
}

/**
更新用户昵称和头像信息
*/
func (*UserDao) UpdateUerInfo(gid string, user *model.UserEntity) (int, string) {

	if len(gid) == 0 {
		return 0, "gid不可以为空！"
	}

	// 插入数据
	err := db.Model(model.UserEntity{}).Where("go_id = ?", gid).Updates(user).Error
	if err != nil {
		log.Printf("update user info failed:%s\n", err.Error())
		return 0, err.Error()
	}
	return 1, ""
}
