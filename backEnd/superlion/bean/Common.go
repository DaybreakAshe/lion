package bean

// LoginReq 登录请求bean
type LoginReq struct {
	Name   string
	Passwd string
}

// CommonResponse 登录请求bean
type CommonResponse struct {
	Data interface{} `json:"data"`
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
}
