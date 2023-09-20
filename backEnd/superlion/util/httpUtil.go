//@program: superlion
//@author: yanjl
//@create: 2023-09-15 15:17
package util

import (
	"bytes"
	"fmt"
	"github.com/u2takey/go-utils/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// ParseResponse 响应体转map
func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}

	if response.StatusCode != http.StatusOK {
		fmt.Println("@[WARNING]response not success:", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}

// 发送from-data 请求
// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, partName string, file *multipart.File, fileName string) (*http.Request, error) {
	// file, err := os.Open(path)
	//if err != nil {
	//	return nil, err
	//}
	// defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fmt.Printf("请求参数：%+v", params)
	// 写入文件

	part, err := writer.CreateFormFile(partName, fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, *file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func RequestPost(formFile os.File, filename string, postURL string) (rqe *http.Request, err error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	//if fw, err := w.CreateFormField("data"); err == nil {
	//	fw.Write(data)
	//}
	//if createFormFile, err := w.CreateFormFile("smfile", filename); err == nil {
	//	readAll, _ := ioutil.ReadAll()
	//	createFormFile.Write(filepath.)
	//}
	w.Close()
	req, err := http.NewRequest(http.MethodPost, postURL, buf)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	return req, err
}
