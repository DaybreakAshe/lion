package bean

// LoginReq 登录请求bean
type LoginReq struct {
	Name   string
	Passwd string
}

// CommonResponse 登录请求bean
type CommonResponse struct {
	// go的每一种数据类型都实现了该接口，因此，其他数据类型都可以赋值给interface{}
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

// FileRspBean 图片资源返回bean
type FileRspBean struct {
	FileId    string `json:"fileId"`
	FileName  string `json:"fileName"`
	FileS3Url string `json:"fileS3Url"`
	FileUrl   any    `json:"fileUrl"`
}

// 查询时分页参数
type PageParams struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

// 分页查询返回结果属性
type PageResult struct {
	Total int64   `json:"total"`
	Data  [][]any `json:"data"`
}
