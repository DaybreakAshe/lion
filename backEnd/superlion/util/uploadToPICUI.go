//@program: superlion
//@author: yanjl
//@create: 2023-09-22 17:21
package util

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
)

var (
	picToken     = "6|NgtXSlzZa83i1Wh3Oqq9J8naLv7ZGWsELILPmv8C"
	picUploadApi = "https://picui.cn/api/v1/upload"
)

// 上传到picui图床，返回0为成功，-1为失败，todo：1为上传到nginx
func UploadToPIC(file *multipart.File, fileName string) (int, string) {

	// 请求参数，*为必填
	/*
		*file	File	图片文件
		token	String	临时上传 Token
		permission	Integer	权限，1=公开，0=私有
		strategy_id	Integer	储存策略ID
		album_id	Integer	相册ID
		expired_at	String	图片过期时间(yyyy-MM-dd HH:mm:ss)
	*/
	params := &map[string]string{
		"permission": "1",
		"album_id":   "lion",
	}

	req, err := newfileUploadRequest(picUploadApi, *params, "file", file, fileName)

	req.Header.Set("Authorization", "Bearer "+picToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "multipart/form-data")
	// http返回的response的body必须close,否则就会有内存泄露

	resp, err := httpClient.Do(req)
	defer func() { // 函数执行结束前才会调用 defer
		resp.Body.Close()
		fmt.Println("finish")
	}()
	if err != nil {
		fmt.Println("###[ERROR]:do post request error:", err.Error())
	}
	mapStr, _ := ParseResponse(resp)

	// 上传失败，todo：上传nginx
	if false == mapStr["status"] {
		fmt.Println("upload to PICUI failed:", mapStr["message"])
		str, _ := json.Marshal(mapStr["message"])
		return -1, string(str)
	}

	dataStr, _ := json.Marshal(mapStr["data"])
	dataMap := &map[string]any{}
	json.Unmarshal(dataStr, dataMap)
	linkMap := map[string]string{}
	dataStr, _ = json.Marshal(mapStr["links"])
	json.Unmarshal(dataStr, &linkMap)

	return 0, linkMap["url"]

}
