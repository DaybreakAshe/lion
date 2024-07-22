//@program: superlion
//@author: yanjl
//@create: 2023-10-08 17:05
package bean

import "superlion/model"

// 保存文章 or 草稿
type SavePostReq struct {
	Category        string  `json:"category"` // 文章分类
	Title           string  `json:"title"`
	ContentType     string  `json:"contentType"`
	MarkdownContent string  `json:"markdownContent"`
	HtmlContent     string  `json:"htmlContent"`
	Comments        int64   `json:"comments"`
	TypeId          int64   `json:"typeId"` // 文章类型ID
	HeadImg         string  `json:"headImg"`
	Top             int32   `json:"top"`     // 置顶
	Marrow          int32   `json:"marrow"`  // 精华1-是，0-否
	Tags            []int64 `json:"tags"`    // 文章标签id
	IsDraft         bool    `json:"isDraft"` // 是否草稿
}

// 更新文章请求
type UpdatePostParams struct {
	Id         int64  `json:"id"`
	Official   int32  `json:"official"` // 官方1-是，0-否
	AuditState string `json:"auditState"`
	Approvals  int64  `json:"approvals"`  // 点赞量
	Collection int64  `json:"collection"` // 收藏量
	Sort       int32  `json:"sort"`
	AuthorId   string `json:"authorId"`
}

// 文章列表查询返回bean
type PostBeanRsp struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	HeadImg    string `form:"headImg" json:"headImg"`
	Official   int32  `json:"official"` // 官方1-是，0-否
	AuditState string `json:"auditState"`
	Views      int64  `json:"views"`
	Approvals  int64  `json:"approvals"`  // 点赞量
	Collection int64  `json:"collection"` // 收藏量
	Sort       int32  `json:"sort"`
	// AuthorId   string `json:"authorId"`
	Preview string      `json:"preview"` // 预览内容
	Tags    []model.Tag `json:"tags"`    //gorm:"foreignKey:tagId;"
}

// 文章列表查询参数
type PostListParams struct {

	// 标题
	Title string `json:"title"`
	// 文章分类
	Category string `json:"category"`
	// 文章类型
	TypeId int64 `json:"typeId"`
	// 官方1-是，0-否
	Official int32 `json:"official"`
	// 精华1-是，0-否
	Marrow int32 `json:"marrow"`
	// 是否草稿1-是，0-否
	IsDraft int32 `json:"isDraft"`
	// 作者ID
	AuthorId string `json:"authorId"`
	// 标签
	TagId []int64 `json:"tagId"`
}