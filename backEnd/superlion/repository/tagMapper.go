//@program: superlion
//@author: yanjl
//@create: 2023-09-25 14:42
package repository

import (
	"fmt"
	"superlion/model"
	"sync"
)

type TagDao struct {
}

//func (model.UserEntity) TableName() string {
//	return "lion_user"
//}

var tagDao *TagDao

/**在 Do 方法被调用后，该函数将被执行，而且只会执行一次，即使在多个协程同时调用的情况下也是如此*/
var tagDaoOnce sync.Once

// NewUserDaoInstance 单例构建Dao
func NewTagDaoInstance() *TagDao {
	tagDaoOnce.Do(
		func() {
			tagDao = &TagDao{}
		})
	return tagDao
}

// 插入新标签
func (*TagDao) AddNewTag(tag *model.LionTag) (int, error) {

	err := db.Create(tag).Error
	if err != nil {
		fmt.Println("insert new tag failed:", err.Error())
		return 1, nil
	} else {
		return 0, err
	}
}

// 查询用户标签
func (*TagDao) FindUserTagsByGoId(goId string) ([]model.LionTag, error) {

	var tags []model.LionTag

	err := db.Where("creator_id = ?", goId).Find(&tags).Error

	if err != nil {
		return nil, err
	} else {
		return tags, nil
	}
}
