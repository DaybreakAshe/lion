//@program: superlion
//@author: yanjl
//@create: 2023-10-08 15:32
package repository

import (
	"log"
	"superlion/model"
	"sync"
	"time"
)

type ArticleTypeDao struct {
}

//func (model.UserEntity) TableName() string {
//	return "lion_user"
//}

var articleTypeDao *ArticleTypeDao

/**在 Do 方法被调用后，该函数将被执行，而且只会执行一次，即使在多个协程同时调用的情况下也是如此*/
var articleTypeDaoOnce sync.Once

// NewUserDaoInstance 单例构建Dao
func NewArticleTypeDaoInstance() *ArticleTypeDao {
	articleTypeDaoOnce.Do(
		func() {
			articleTypeDao = &ArticleTypeDao{}
		})
	return articleTypeDao
}

// 新增文章类别
func (*ArticleTypeDao) AddNewArticleType(typeEntity *model.LionArticleType) string {

	if len(typeEntity.Name) == 0 {
		log.Panicf("类别名不能为空")
		return "类别名不能为空"
	}

	typeEntity.UpdateAt = time.Now()

	err := db.Create(typeEntity).Error
	if err != nil {
		log.Print("类别名不能为空", err.Error())
		return "新增失败..."
	}
	return ""
}

// 查询文章类别
func (*ArticleTypeDao) FindUserArtTypes(typeEntity *model.LionArticleType) (*[]model.LionArticleType, string) {

	rlt := &[]model.LionArticleType{}

	err := db.Where(typeEntity).Find(rlt).Error

	if err != nil {
		log.Print("查询类别名失败...", err.Error())
		return nil, "查询失败..."
	}
	return rlt, ""
}
