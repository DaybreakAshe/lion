//@program: superlion
//@author: yanjl
//@create: 2023-09-27 09:39
package model

import "time"

// LionCache 缓存表
type LionCache struct {
	ID       int64     `json:"id" gorm:"id"`              // 主键
	Key      string    `json:"key" gorm:"key"`            // 缓存键
	Value    string    `json:"value" gorm:"value"`        // 缓存值
	Type     string    `json:"type" gorm:"type"`          // 业务类型
	IsDelete int8      `json:"isDelete" gorm:"is_delete"` // 删除标识（0:未删除、1:已删除）
	CreateAt time.Time `json:"createAt" gorm:"create_at"` // 记录创建时间
	UpdateAt time.Time `json:"updateAt" gorm:"update_at"` // 记录修改时间
}

// TableName 表名称
func (*LionCache) TableName() string {
	return "lion_cache"
}
