//@program: superlion
//@author: yanjl
//@create: 2023-10-08 10:39
package constants

type PostTypeCode string

type PostContentCode string

// 业务类型
const (
	// 文章类型
	POST_DRAFT   = "00" // 文章草稿
	POST_PUBLISH = "01" // 已发布文章

	// 文章内容类型
	POST_CONTENT_MARKDOWN = "0"
	POST_CONTENT_HTML     = "1"

	// 审核状态
	POST_AUDIT_STATE_UNKONW = "01"
	POST_AUDIT_STATE_NO     = "00"
	POST_AUDIT_STATE_YES    = "02"
)

// GetContentEnumValue 根据枚举常量获取值
func (code PostContentCode) GetContentEnumValue() string {

	switch code {
	case POST_CONTENT_HTML:
		return "html"
	case POST_CONTENT_MARKDOWN:
		return "markdown"
	default:
		return ""
	}
}

// GetPostTypeEnumValue 获取文章类型
func (code PostTypeCode) GetPostTypeEnumValue() string {
	switch code {
	case POST_PUBLISH:
		return "published"
	case POST_DRAFT:
		return "draft"
	default:
		return ""
	}
}
