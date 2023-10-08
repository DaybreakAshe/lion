//@program: superlion
//@author: yanjl
//@create: 2023-09-25 14:31
package bean

import "time"

// 标签
type TagReqBean struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	GroupName   string `json:"group_name"` //所属分组
}

type TagRspBean struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	RefCnt      int64  `json:"refCnt"`
}

// 帖子缓存请求结构体
type PostCacheReq struct {
	Id    int64  `json:"id"`    // 传入主键则为修改
	Key   string `json:"key"`   // 缓存键
	Value string `json:"value"` // 缓存值
	Type  string `json:"type" ` // 业务类型
}

// 帖子缓存响应结构体
type PostCacheRsp struct {
	Id       int64     `json:"id"`        // 传入主键则为修改
	Key      string    `json:"key"`       // 缓存键 暂存作者id？todo
	Value    string    `json:"value"`     // 缓存值
	Type     string    `json:"type" `     // 业务类型
	CreateAt time.Time `json:"createAt" ` // 记录创建时间
	UpdateAt time.Time `json:"updateAt"`  // 记录修改时间
}

// 文章类别相关
// 新增文章类别req
type ArticleTypeReq struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// 查询文章类别请求参数
type FindArtTypeReq struct {
	Page   *PageParams `json:"page"`
	UserId int64       `json:"userId"`
	Name   string      `json:"name"`
}

// 查询文章类别响应参数
type UserArtTypeRsp struct {
	ID          int64     `json:"id" gorm:"id"`                   // 主键
	AuditState  string    `json:"audit_state" gorm:"audit_state"` // 审核状态
	Name        string    `json:"name" gorm:"name"`               // 名称
	Description string    `json:"description" gorm:"description"` // 描述
	RefCount    int64     `json:"ref_count" gorm:"ref_count"`     // 引用统计
	Scope       string    `json:"scope" gorm:"scope"`             // 作用域
	CreateAt    time.Time `json:"create_at" gorm:"create_at"`     // 记录创建时间
	UpdateAt    time.Time `json:"update_at" gorm:"update_at"`     // 记录修改时间
}
