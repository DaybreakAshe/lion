//@program: superlion
//@author: yanjl
//@create: 2023-09-25 14:31
package bean

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
