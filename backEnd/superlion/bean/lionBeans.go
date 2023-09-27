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
	Key      string    `json:"key"`       // 缓存键
	Value    string    `json:"value"`     // 缓存值
	Type     string    `json:"type" `     // 业务类型
	CreateAt time.Time `json:"createAt" ` // 记录创建时间
	UpdateAt time.Time `json:"updateAt"`  // 记录修改时间
}
