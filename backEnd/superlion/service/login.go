package service

import (
	"fmt"
	"github.com/u2takey/go-utils/uuid"
	"superlion/bean"
)

func PrintHello() {
	fmt.Println("hello !!!!!!")
}

func Login(req *bean.LoginReq) (string, string) {

	fmt.Printf("login request params:name=%s,pwd=%s\n", req.Name, req.Passwd)

	if len(req.Name) == 0 || len(req.Passwd) == 0 {

		return "", "登录信息不能为空"
	}

	// 校验数据库，检查登录信息

	// 签发一个token
	token := uuid.NewUUID()

	// todo
	// 保存redis, 返回登录信息

	return token, ""
}

/**
登录信息结构体：（google返回的参数json）：
	示例数据：
	"state": "3EAB37D9D5310BFE",
	"access_token": "ya29.a0AfB_byCtVB2voZknHPiip_S8SBjWqVGx_Wf3uHYizmJmNNubIwwrDgK_juTmvz5U86lV17W54IIPjgXZuHgmUUKGe8sKa-ZurtNS0wo5RiPKRQaCgYKAbASARESFQHsvYlsoOkFLgS-kaFRpSTvA18sCw0173",
	"token_type": "Bearer",
	"expires_in": "3599",
	"scope": "email https://www.googleapis.com/auth/userinfo.email openid",
	"authuser": "0",
	"prompt": "none"
*/
type LoginParmas struct {
	State       string `json:"state"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
	Authuser    string `json:"authuser"`
	Prompt      string `json:"prompt"`
}
