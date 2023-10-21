//@program: superlion
//@author: yanjl
//@create: 2023-10-13 15:56
package repository

import (
	"superlion/model"
	"sync"
)

type TagPostDao struct {
}

var (
	tagPostDao *TagPostDao
	tagPostMap sync.Once
)

// 构建实例
func NewTagPostMapInstance() *TagPostDao {
	tagPostMap.Do(
		func() {
			tagPostDao = &TagPostDao{}
		})
	return tagPostDao
}

// SaveTagPostMapList 保存标签-文章关系表
func (*TagPostDao) SaveTagPostMapList(dataList []model.LionTagPostsMapping) string {
	if len(dataList) == 0 {
		return ""
	}

	err := db.Create(dataList).Error.Error()
	return err
}
