//@program: superlion
//@author: yanjl
//@create: 2023-09-27 10:01
package repository

import (
	"superlion/model"
	"sync"
)

type CacheDao struct {
}

var (
	cacheDao     *CacheDao
	cacheDaoOnce sync.Once
)

// NewCacheDaoInstance 单例构建Dao
func NewCacheDaoInstance() *CacheDao {
	cacheDaoOnce.Do(
		func() {
			cacheDao = &CacheDao{}
		})
	return cacheDao
}

// 新增一个草稿（缓存）
func (*CacheDao) SaveNewCache(entity *model.LionCache) error {

	err := db.Create(entity).Error

	return err
}

// 更新一个草稿（缓存）
func (*CacheDao) UpdateCache(entity *model.LionCache) error {

	err := db.Where("id = ?", entity.ID).Updates(entity).Error

	return err
}

// 查询一个草稿（缓存）
func (*CacheDao) FindCacheById(id int64) (*model.LionCache, error) {

	cacheEntity := &model.LionCache{}

	err := db.Where("id = ?", id).Find(cacheEntity).Error

	return cacheEntity, err
}
