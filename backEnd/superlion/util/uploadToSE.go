//@program: superlion
//@author: yanjl
//@create: 2023-09-21 11:16
package util

import (
	"bufio"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	IMGTU_USER_NAME             = "draven.yjl@outlook.com"
	IMGTU_PASSWORD              = "YJLscanf&152919"
	IMGTU_INIT_URL              = "https://imgtu.com/init"
	IMGTU_LOGIN_URL             = "https://imgtu.com/login"
	IMGTU_OPERATE_URL           = "https://imgtu.com/json"
	SESSION_ID_PATTERN, _       = regexp.Compile("PHPSESSID=([^;]*); path=/; HttpOnly")
	AUTH_TOKEN_PATTERN, _       = regexp.Compile("PF\\.obj\\.config\\.auth_token = \"([0-9a-f]{40})\";")
	KEEP_LOGIN_PATTERN, _       = regexp.Compile("KEEP_LOGIN=([^;]*);")
	INIT_VALID_DURATION   int64 = 15 * 60 * 1000
	LOGIN_VALID_DURATION  int64 = 30 * 24 * 60 * 60 * 1000

	loginTimestamp int64 = 0
	initTimestamp  int64 = 0
	sessionId            = ""
	authToken            = ""
	keepLogin            = "7Gres%3Ad75e69bb1ccee4f4f3a8c39ff4c62dd82c39247e882908d074c7b9e31c3fcb42ac00852b27ed4cf15f92c82fa1d1ed5bb27ae144135b7763f7319353ac4544f102b5e3dce18b3e999850d6ca0abfddfa2e23f1c114d954935883e83d1b8f78d039b2c6d6c7513cc6909eef1%3A1695285196"
)

// 初始化认证信息
func InitSession(forceAction bool) (bool, string) {

	if !forceAction && !isSessionIdExpired() {
		fmt.Println("【初始化】成功：会话有效期内，无需重新初始化。")
		return true, ""
	}
	resp, err := http.Get(IMGTU_INIT_URL)
	if err != nil {
		fmt.Println("[WARNING]init session failed:", err.Error())
		return false, "init failed"
	}

	// 获取请求头信息：
	cookie := resp.Header["Set-Cookie"]

	// 获取sessionId
	fmt.Printf("Cookies info:%s\n", resp.Header)
	match := SESSION_ID_PATTERN.FindStringSubmatch(cookie[0])
	sessionId = match[len(match)-1]
	if len(sessionId) == 0 {
		fmt.Println("【初始化】失败：获取SessionId失败。")
		return false, "get sessionId failed"
	}

	// 获取auth_token

	// 1\创建文件
	f, err := os.Create("./doc/cache.txt")
	if err != nil {
		fmt.Println("error open file :", err.Error())
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		fmt.Println("error write file :", err.Error())
		return false, "open file error"
	}
	// 2\讀取文件543行獲取token
	rawToken := ReadLine(543, ".cache.txt")
	match = AUTH_TOKEN_PATTERN.FindStringSubmatch(rawToken)
	authToken = match[1]
	fmt.Println("token info :", authToken)
	initTimestamp = time.Now().UnixMilli()
	fmt.Println("auth info init over")
	return true, ""
}

func Login() bool {
	return LoginF(false)
}

func LoginF(forceAction bool) bool {
	if !forceAction && !isLoginExpired() {
		fmt.Println("【登录】成功：登录状态有效期内，无需重新登录。")
		return false
	}
	// 初始化会话
	if isSessionIdExpired() {
		b, err := InitSession(false)
		if !b {
			fmt.Println("【登录】失败：初始化会话受阻。", err)
		}
		// 设置请求头
		data := "login-subject=" + IMGTU_USER_NAME + "&password=" + IMGTU_PASSWORD + "&auth_token=" + authToken

		req, _ := http.NewRequest(http.MethodPost, IMGTU_LOGIN_URL, strings.NewReader(data))
		// req.Header.Set("cookie", "PHPSESSID="+sessionId+";")
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
		req.Header.Set("connection", "keep-alive")
		//req.PostForm.Add("login-subject", IMGTU_USER_NAME)
		//req.PostForm.Add("password", IMGTU_PASSWORD)
		//req.PostForm.Add("auth_token", authToken)

		myclient := &http.Client{}
		rsp, eor := myclient.Do(req)
		if eor != nil {
			fmt.Println("get login error:", eor.Error())
			return false
		}
		rspMap, _ := ParseResponse(rsp)
		fmt.Println("rsp info:", rspMap)
		cookie := rsp.Header["Set-Cookie"]

		// 获取sessionId
		fmt.Printf("login Cookies info:%s\n", rsp.Header)
		match := KEEP_LOGIN_PATTERN.FindStringSubmatch(cookie[0])
		// sessionId = match[0]
		fmt.Printf("login keep info:%s\n", match)
		if len(match) == 0 {
			fmt.Println("【初始化】失败：获取SessionId失败。")
			// return false
		}

	}
	return true
}

func Upload(out multipart.File) map[string]any {

	InitSession(false)

	body := map[string]string{
		// "source":     "C:\\Users\\Yanjilong\\Desktop\\LinkinPark\\lol\\Thresh_3.jpg",
		"type":       "file",
		"action":     "upload",
		"auth_token": authToken,
		"nsfw":       "0",
	}

	// bytesF, _ := json.Marshal(body)

	//file, _ := os.Open("C:\\Users\\Yanjilong\\Desktop\\LinkinPark\\lol\\Thresh_3.jpg")
	//out := multipart.File(file)

	req, eor := newfileUploadRequest(seUploadApi, body, "source", &out, "test1.jpg")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "PHPSESSID="+sessionId+"; KEEP_LOGIN="+keepLogin)

	if eor != nil {
		fmt.Println("###[ERROR]:new request error:", eor.Error())
	}

	// 获取client 指针！
	httpClient := &http.Client{}
	// http返回的response的body必须close,否则就会有内存泄露
	resp, eor := httpClient.Do(req)
	defer func() { // 函数执行结束前才会调用 defer
		resp.Body.Close()
		fmt.Println("finish")
	}()
	if eor != nil {
		fmt.Println("###[ERROR]:do post request error:", eor.Error())
	}
	mapStr, _ := ParseResponse(resp)

	fmt.Println("[INFO] post req to save pic over :", resp.StatusCode, mapStr)

	return mapStr
}

func isSessionIdExpired() bool {
	return initTimestamp+INIT_VALID_DURATION < time.Now().UnixMilli()
}

func isLoginExpired() bool {
	return loginTimestamp+LOGIN_VALID_DURATION < time.Now().UnixMilli()
}

func ReadLine(lineNumber int, path string) string {
	file, _ := os.Open(path)
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text()
		}
		lineCount++
	}
	defer file.Close()
	return ""
}
