package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/u2takey/go-utils/json"
	"github.com/u2takey/go-utils/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"superlion/bean"
	"superlion/config"
	"superlion/config/webConfig"
	"superlion/model"
	"superlion/repository"
	"time"
)

var (
	ctx          = context.Background()
	userNotFound = errors.New("user not found")
)

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

	var errMsg string

	rsp := &bean.CommonResponse{}
	// 请求谷歌api，获取用户信息
	url := "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

	url = url + params.AccessToken

	resp, eor := http.Get(url)

	// fmt.Printf("get google url:[%s] rsp code:%d\n", url, resp.StatusCode)
	// fmt.Printf("get google response info:%s\n" + resp.Body.Close().Error())
	if eor != nil {

		rsp.Code = 604
		rsp.Msg = "请求google出错了"
		fmt.Printf("get google info error:%s\n", eor)
		fmt.Printf("get google url:%s, rsp code:%d\n", url, resp.StatusCode)
		return rsp, ""
	} else {
		// 200 => 请求成功
		if http.StatusOK == resp.StatusCode {

			result, err := ParseResponse(resp)
			if err != nil {
				rsp.Code = 601
				rsp.Msg = "json转换出错"
				fmt.Printf("json prase error :%s\n", err.Error())
			}

			// 定义返回结构体
			goUserInfo := &GoUserInfo{}
			// map转结构体
			perr := mapstructure.Decode(&result, goUserInfo)

			if perr != nil {
				rsp.Code = 601
				rsp.Msg = "json格式化出错"
				fmt.Printf("json prase error :%s\n", perr.Error())
			} else {
				jsonData, err := json.Marshal(goUserInfo)
				goUserInfo.GoToken = params.AccessToken
				if err != nil {
					rsp.Code = 601
					rsp.Msg = "json格式化出错!"
					fmt.Printf("json format error:%s\n", err.Error)
				} else {
					fmt.Printf("get google body info :%s\n", jsonData)
					// 判断缓存是否存在：
					if len(goUserInfo.LionToken) == 0 {
						// 未带token登录，认为第一次登录？
						ltoken := uuid.NewUUID()
						goUserInfo.LionToken = ltoken
						SaveTokenToCache(goUserInfo)
					}
					rsp.Code = 200
					rsp.Data = *goUserInfo
				}

				// 异步插入数据库 todo
				log.Println("sync request over,start save info to DB")
				// 判断数据是否已存在
				user, _ := repository.NewUserDaoInstance().GetUserInfoByGId(goUserInfo.Id)

				// 不存在则插入
				if user == nil {
					SaveUserInfoToDB(goUserInfo)
				}

			}
		} else {
			rsp.Code = 608
			rsp.Msg = "谷歌token无效"
			errMsg = "谷歌token无效"
			return rsp, errMsg
		}
	}
	fmt.Printf("ready to return :%s\n", *rsp)

	return rsp, errMsg
}

/**
保存登录信息到redis，设置3天过期 todo
*/
func SaveTokenToCache(user *GoUserInfo) {

	redisP := config.GetRedisHelper()

	// 缓存3天
	name, err := redisP.Set(ctx, webConfig.RedisPre+user.LionToken, user, 24*3*time.Hour).Result()
	if err != nil {
		log.Fatal(err)
		log.Panicf("缓存用户失败:%s\n", err.Error())
		return
	}
	log.Printf("cache user to redis over,%s\n", name)
}

func SaveUserInfoToDB(user *GoUserInfo) (int, string) {

	userEntity := &model.UserEntity{
		GoId:            user.Id,
		GoEmail:         user.Email,
		GoLocale:        user.Locale,
		GoName:          user.Name,
		GoPicture:       user.Picture,
		GoToken:         user.LionToken,
		GoVerifiedEmail: user.VerifiedEmail,

		LoginName:  user.Name,
		Avatar:     user.Picture,
		Status:     "00",
		CreateTime: time.Now(),
	}

	rows, eor := repository.NewUserDaoInstance().SaveUerInfoToDB(userEntity)
	if len(eor) != 0 {
		return 0, eor
	}

	return rows, ""
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
根据谷歌id获取用户信息
*/
func GetUserInfoByGoId(gid string) (*LionUserInfo, error) {

	if len(gid) == 0 {
		return nil, errors.New("gid不可以为空")
	}

	user, _ := repository.NewUserDaoInstance().GetUserInfoByGId(gid)

	if user != nil {
		lionUser := &LionUserInfo{
			GoId:     user.GoId,
			GoName:   user.GoName,
			GoEmail:  user.GoEmail,
			UserName: user.LoginName,
			Avatar:   user.Avatar,
		}
		return lionUser, nil
	} else {
		return nil, userNotFound
	}

}

// 实现序列化？,保存redis必须实现
func (u *GoUserInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func UpdateUserInfo(user *GoUserInfo, req *bean.UpdateUserInfoBean) string {

	nickName := req.NickName
	avatar := req.Avatar

	// 名称和链接均不为空
	if len(nickName) != 0 && len(avatar) != 0 {
		userEntity := &model.UserEntity{
			Avatar:    avatar,
			LoginName: nickName,
			Signature: req.Signature,
		}

		rows, err := repository.NewUserDaoInstance().UpdateUerInfo(user.Id, userEntity)
		if rows != 0 {
			return ""
		} else {
			return err
		}
	}
	return "failed"
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
	AuthUser    string `json:"authuser"`
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
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	GoToken       string `json:"goToken"`
	// 系统token
	LionToken string `json:"lionToken"`
}

// LionUserInfo 前端可见用户bean
type LionUserInfo struct {
	GoId     string
	UserName string
	Avatar   string
	GoName   string
	GoEmail  string
}
