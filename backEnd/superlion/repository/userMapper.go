//@program: superlion
//@author: yanjl
//@create: 2023-09-07 15:48
package repository

import "superlion/model"

type UserDao struct {
}

func (UserDao) TableName() string {
	return "lion_user"
}

/**
通过googleId查询用户
*/
func (*UserDao) GetUserInfoByGId(gid string) (*model.UserEntity, string) {

	if len(gid) == 0 {
		return nil, "id不可以为空！"
	}
	return nil, ""
}
