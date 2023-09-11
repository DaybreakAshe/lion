package service

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/u2takey/go-utils/json"
	"github.com/u2takey/go-utils/uuid"
	"io/ioutil"
	"net/http"
	"superlion/bean"
)

func PrintHello() {
	fmt.Println("hello !!!!!!")
}

/**
获取google授权后信息,
*/
func GetGoogleAuthBody(params LoginParmas) (*bean.CommonResponse, string) {

	// 打印json
	jsonstr, err := json.Marshal(params)
	if err != nil {
		fmt.Printf("json format error:%s\n", err.Error)
		return &bean.CommonResponse{}, err.Error()
	}
	fmt.Printf("recevice auth body :%s\n", string(jsonstr))

	rsp := bean.CommonResponse{
		Data: string(jsonstr),
		Code: "200",
		Msg:  "ok",
	}

	var errMsg string

	// 请求谷歌api，获取用户信息
	url := "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

	url = url + params.AccessToken

	resp, eor := http.Get(url)

	// fmt.Printf("get google response info:%s\n" + resp.Body.Close().Error())
	if eor != nil {
		rsp.Data = ""
		rsp.Code = "450"
		rsp.Msg = "请求google出错了"
		errMsg = rsp.Msg
		fmt.Printf("get google info error:%s\n", err.Error())
	} else {
		// 200 => 请求成功
		if http.StatusOK == resp.StatusCode {

			result, err := ParseResponse(resp)
			if err != nil {
				rsp.Code = "601"
				rsp.Msg = "json转换出错"
				fmt.Printf("json prase error :%s\n", err.Error())
			}

			var goUserInfo GoUserInfo
			perr := mapstructure.Decode(result, &goUserInfo)

			if perr != nil {
				rsp.Code = "601"
				rsp.Msg = "json格式化出错"
				fmt.Printf("json prase error :%s\n", perr.Error())
			} else {
				jsonData, err := json.Marshal(params)
				if err != nil {
					rsp.Code = "601"
					rsp.Msg = "json格式化出错!"
					fmt.Printf("json format error:%s\n", err.Error)
				}
				fmt.Printf("get google body info :%s\n", jsonData)
				rsp.Data = goUserInfo
			}
		}
	}

	return &rsp, errMsg
}

/**
登录接口 todo：未完成，未使用
*/
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

// ParseResponse 响应体转map
func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}

/**
登录信息结构体：（google返回的参数json）：
	示例数据：
	"state": "3EAB37D9D5310BFE",
	"access_token": "ya29.a0AfB_byCtVB2voZknHPiip_S8SBjWqVGx_Wf3uHYizmJm
	"token_type": "Bearer",
	"expires_in": "3599",
	"scope": "email https://www.googleapis.com/auth/userinfo.email openid",
	"authuser": "0",
	"prompt": "none"
*/
type LoginParmas struct {
	State       string `json:"state"`
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   string `json:"expiresIn"`
	Scope       string `json:"scope"`
	Authuser    string `json:"authuser"`
	Prompt      string `json:"prompt"`
}

/**
	谷歌api返回用户信息
url = https://www.googleapis.com/oauth2/v2/userinfo?access_token=ya29.a0AfB_byCk_X
    "id": "106256997442594399678",
    "email": "dravenxue@gmail.com",
    "verified_email": true,
    "name": "Draven XUE",
    "given_name": "Draven",
    "family_name": "XUE",
    "picture": "https://lh3.googleusercontent.com/a/ACg8ocKRSkY1TrhbRJEos2-LBYb6fzHAZa7rcR6vWjZZfizxcA=s96-c",
    "locale": "zh-CN"
*/
type GoUserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verifiedEmail"`
	Name          string `json:"name"`
	GivenName     string `json:"givenName"`
	FamilyName    string `json:"familyName"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}
