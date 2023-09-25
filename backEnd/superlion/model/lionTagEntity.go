//@program: superlion
//@author: yanjl
//@create: 2023-09-25 14:17
package model

// 标签表
type LionTag struct {
	Id          int64  `gorm:"column:id" json:"id"`                   //主键
	AuditState  string `gorm:"column:audit_state" json:"audit_state"` //审核状态
	GroupName   string `gorm:"column:group_name" json:"group_name"`   //所属分组
	Name        string `gorm:"column:name" json:"name"`               //名称
	Description string `gorm:"column:description" json:"description"` //描述
	RefCount    int64  `gorm:"column:ref_count" json:"ref_count"`     //引用统计
	CreatorId   string `gorm:"column:creator_id" json:"creator_id"`   //创建人,go_id
	IsDelete    uint8  `gorm:"column:is_delete" json:"is_delete"`     //删除标识（0:未删除、1:已删除）
}

func (*LionTag) TableName() string {
	return "lion_tag"
}
