package service

import (
	"fmt"
	"superlion/bean"
)

func Login(req *bean.LoginReq) (string, string) {

	fmt.Printf("login request params:name=%s,pwd=%s\n", req.Name, req.Passwd)

	if len(req.Name) == 0 || len(req.Passwd) == 0 {

		return "", "登录信息不能为空"
	}

	return "token", ""
}
