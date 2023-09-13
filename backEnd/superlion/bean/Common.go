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
