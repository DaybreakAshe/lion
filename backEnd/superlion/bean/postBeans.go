//@program: superlion
//@author: yanjl
//@create: 2023-10-08 17:05
package bean

// 保存文章 or 草稿
type SavePostReq struct {
	Category        string  `json:"category"`
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
type UpdatePostParams struct {
	Id         int64  `json:"id"`
	Official   int32  `json:"official"` // 官方1-是，0-否
	AuditState string `json:"auditState"`
	Approvals  int64  `json:"approvals"`  // 点赞量
	Collection int64  `json:"collection"` // 收藏量
	Sort       int32  `json:"sort"`
	AuthorId   string `json:"authorId"`
}
