package bean

// 登录请求bean
type LoginReq struct {
	Name   string
	Passwd string
}

// 登录请求bean
type CommonResponse struct {
	Data string
	Code string
	Msg  string
}
