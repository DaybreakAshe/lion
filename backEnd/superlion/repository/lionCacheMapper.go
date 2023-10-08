//@program: superlion
//@author: yanjl
//@create: 2023-09-27 10:01
package repository

import (
	"superlion/constants"
	"superlion/model"
	"superlion/util"
	"sync"
)

type CacheDao struct {
}

var (
	cacheDao     *CacheDao
	cacheDaoOnce sync.Once

	// 文章草稿类型：
	cacheTypeDraft = constants.POST_DRAFT
)

// NewCacheDaoInstance 单例构建Dao
func NewCacheDaoInstance() *CacheDao {
	cacheDaoOnce.Do(
		func() {
			cacheDao = &CacheDao{}
		})
	return cacheDao
}

// 新增一个草稿（缓存表）
func (*CacheDao) SaveNewCache(entity *model.LionCache) error {

	err := db.Create(entity).Error

	return err
}

// 更新一个草稿（缓存表）
func (*CacheDao) UpdateCache(entity *model.LionCache) error {

	err := db.Where("id = ? and type = ?", entity.ID, cacheTypeDraft).Updates(entity).Error

	return err
}

// 查询一个草稿（缓存表）
func (*CacheDao) FindCacheById(id int64) (*model.LionCache, error) {

	cacheEntity := &model.LionCache{}

	err := db.Where("id = ? and type = ?", id, cacheTypeDraft).Find(cacheEntity).Error

	return cacheEntity, err
}

// 查询草稿列表（缓存表）
func (*CacheDao) CacheList(page *model.PageDto, cache *model.LionCache) ([]model.LionCache, int64, error) {

	var cacheEntity []model.LionCache

	var total int64
	err := db.Scopes(util.Paginate(page)).Where(cache).Find(&cacheEntity).Count(&total).Error

	return cacheEntity, total, err
}
