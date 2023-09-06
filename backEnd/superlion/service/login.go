package service

import (
	"superlion/bean"
)

func Login(req *bean.LoginReq) (string, string) {

	if len(req.Name) == 0 || len(req.Passwd) == 0 {

		return "", "登录信息不能为空"
	}

	return "token", ""
}
