//@program: superlion
//@author: yanjl
//@create: 2023-09-14 16:20
package model

import (
	"fmt"
	"time"
)

// 文章或帖子实体类
type LionPostEntity struct {
	Id              int64  `gorm:"primaryKey column:id;" form:"id" json:"id"`
	AuditState      string `gorm:"column:audit_state;" form:"auditState" json:"auditState"`
	Category        string `gorm:"column:category;" form:"category" json:"category"`
	AuthorId        string `gorm:"column:author_id;" form:"authorId" json:"authorId"`
	Title           string `gorm:"column:title;" form:"title" json:"title"`
	ContentType     string `gorm:"column:content_type;" form:"contentType" json:"contentType"`
	MarkdownContent string `gorm:"column:markdown_content;" form:"markdownContent" json:"markdownContent"`
	HtmlContent     string `gorm:"column:html_content;" form:"htmlContent" json:"htmlContent"`
	Views           int64  `gorm:"column:views;" form:"views" json:"views"`                // 浏览量
	Approvals       int64  `gorm:"column:approvals;" form:"approvals" json:"approvals"`    // 点赞量
	Collection      int64  `gorm:"column:collection;" form:"collection" json:"collection"` // 收藏量
	Comments        int64  `gorm:"column:comments;" form:"comments" json:"comments"`
	TypeId          int64  `gorm:"column:type_id;" form:"typeId" json:"typeId"`
	HeadImg         string `gorm:"column:head_img;" form:"headImg" json:"headImg"`
	Official        int32  `gorm:"column:official;" form:"official" json:"official"`
	Top             int32  `gorm:"column:top;" form:"top" json:"top"`
	Sort            int32  `gorm:"column:sort;" form:"sort" json:"sort"`
	Marrow          int32  `gorm:"column:marrow;" form:"marrow" json:"marrow"`
	CommentId       int64  `gorm:"column:comment_id;" form:"commentId" json:"commentId"`
	IsDelete        int32  `gorm:"column:is_delete;" form:"isDelete" json:"isDelete"`
	CreateAt        string `gorm:"column:create_at;" form:"createAt" json:"createAt"`
	UpdateAt        string `gorm:"column:update_at;" form:"updateAt" json:"updateAt"`
}

func (LionPostEntity) TableName() string {
	return "lion_post"
}

type jsonTime time.Time

//实现它的json序列化方法
func (this jsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// 文章或帖子实体类
type PostEntity struct {
	Id              int64    `gorm:"primaryKey column:id;" form:"id" json:"id"`
	AuditState      string   `gorm:"column:audit_state;" form:"auditState" json:"auditState"`
	Category        string   `gorm:"column:category;" form:"category" json:"category"`
	AuthorId        string   `gorm:"column:author_id;" form:"authorId" json:"authorId"`
	AuthorName      string   `gorm:"column:authorName;" form:"AuthorName" json:"AuthorName"`
	Title           string   `gorm:"column:title;" form:"title" json:"title"`
	ContentType     string   `gorm:"column:content_type;" form:"contentType" json:"contentType"`
	MarkdownContent string   `gorm:"column:markdown_content;" form:"markdownContent" json:"markdownContent"`
	HtmlContent     string   `gorm:"column:html_content;" form:"htmlContent" json:"htmlContent"`
	Views           int64    `gorm:"column:views;" form:"views" json:"views"`                // 浏览量
	Approvals       int64    `gorm:"column:approvals;" form:"approvals" json:"approvals"`    // 点赞量
	Collection      int64    `gorm:"column:collection;" form:"collection" json:"collection"` // 收藏量
	Comments        int64    `gorm:"column:comments;" form:"comments" json:"comments"`
	TypeId          int64    `gorm:"column:type_id;" form:"typeId" json:"typeId"`
	HeadImg         string   `gorm:"column:head_img;" form:"headImg" json:"headImg"`
	Official        int32    `gorm:"column:official;" form:"official" json:"official"`
	Top             int32    `gorm:"column:top;" form:"top" json:"top"`
	Sort            int32    `gorm:"column:sort;" form:"sort" json:"sort"`
	Marrow          int32    `gorm:"column:marrow;" form:"marrow" json:"marrow"` // 精华
	CommentId       int64    `gorm:"column:comment_id;" form:"commentId" json:"commentId"`
	IsDelete        int32    `gorm:"column:is_delete;" form:"isDelete" json:"isDelete"`
	CreateAt        jsonTime `gorm:"column:create_at;" form:"createAt" json:"createAt"`
	UpdateAt        jsonTime `gorm:"column:update_at;" form:"updateAt" json:"updateAt"`
	Tags            []Tag    `gorm:"-" json:"tags"` //gorm:"foreignKey:tagId;"
}

type Tag struct {
	TagId int64  `gorm:"column:tagId;" json:"tagId"`
	Tag   string `gorm:"column:tags;" json:"tag"`
}
