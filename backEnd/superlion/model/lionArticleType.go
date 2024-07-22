//@program: superlion
//@author: yanjl
//@create: 2023-10-08 15:31
package model

import "time"

// LionArticleType 文章类别表
type LionArticleType struct {
	ID          int64     `json:"id" gorm:"id"`                   // 主键
	AuditState  string    `json:"audit_state" gorm:"audit_state"` // 审核状态
	Name        string    `json:"name" gorm:"name"`               // 名称
	Description string    `json:"description" gorm:"description"` // 描述
	RefCount    int64     `json:"ref_count" gorm:"ref_count"`     // 引用统计
	Scope       string    `json:"scope" gorm:"scope"`             // 作用域
	CreatorId   string    `json:"creator_id" gorm:"creator_id"`   // 创建人goId
	IsDelete    int8      `json:"is_delete" gorm:"is_delete"`     // 删除标识（0:未删除、1:已删除）
	CreateAt    time.Time `json:"create_at" gorm:"create_at"`     // 记录创建时间
	UpdateAt    time.Time `json:"update_at" gorm:"update_at"`     // 记录修改时间
}

// TableName 表名称
func (*LionArticleType) TableName() string {
	return "lion_article_type"
}
