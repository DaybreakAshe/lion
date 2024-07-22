//@program: superlion
//@author: yanjl
//@create: 2023-09-27 09:23
package model

import "time"

// 标签-帖子关系表
type LionTagPostsMapping struct {
	ID       int64     `json:"id" gorm:"id"`              // 主键
	TagId    int64     `json:"tagId" gorm:"tag_id"`       // 标签ID
	PostId   int64     `json:"postId" gorm:"post_id"`     // 帖子ID
	IsDelete int8      `json:"isDelete" gorm:"is_delete"` // 删除标识（0:未删除、1:已删除）
	CreateAt time.Time `json:"createAt" gorm:"create_at"` // 记录创建时间
	UpdateAt time.Time `json:"updateAt" gorm:"update_at"` // 记录修改时间
}

func (*LionTagPostsMapping) TableName() string {
	return "lion_tag_posts_mapping"
}
