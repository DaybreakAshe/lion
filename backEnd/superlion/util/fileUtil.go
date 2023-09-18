//@program: superlion
//@author: yanjl
//@create: 2023-09-15 09:52
package util

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/u2takey/go-utils/json"
	"io/ioutil"
	"net/http"
)

var (
	// app_key_id = "0051893aeb584100000000002"
	app_key_id = "1893aeb58410"
	app_key    = "005ac72a2796c83df00a03c24582eacba7e434e9d5"
	// app_key    = "K005qxtrUJKLMnbh9k0k8DcU7fsh5Hg"
	buket_name = "superlion"
	bucketId   = "d1d8f9234a8e9b7588a40110"

	ctx = context.Background()
)

var (
	api     = "https://api.backblazeb2.com/"
	account = "b2api/v2/b2_authorize_account"

	api_api = "https://api005.backblazeb2.com/"

	uploadUrl = "https://pod-050-1007-13.backblaze.com/b2api/v2/b2_upload_file/d1d8f9234a8e9b7588a40110/c005_v0501007_t0052"
	//download_api = "https://f005.backblazeb2.com"
	authorizationToken = ""
)

// UploadToB2 上传文件到B2
func UploadToB2() {

	url := api + account

	auth := app_key_id + ":" + app_key

	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))

	fmt.Printf(basic)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// httpClient := request.Header.Set("Authorization", basic)
	request.Header.Set("Authorization", basic)

	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(request)
	// resp.Header.Set("Content-Type", " application/json;charset=UTF-8")
	if err != nil {
		fmt.Println("rq:" + err.Error())
		return
	}
	rspMap, err := ParseResponse(resp)
	if err != nil {
		fmt.Println("j1:" + err.Error())
		return
	}

	jsonstr, err := json.Marshal(rspMap)
	if err != nil {
		fmt.Println("j2:" + err.Error())
		return
	}
	fmt.Printf("receive body data:\nmap:%s\n", jsonstr)
	authorizationToken = rspMap["authorizationToken"].(string)
	//getFileUploadUrl()

	uploadFileToB2(rspMap["authorizationToken"].(string))
}

func getFileUploadUrl() {

	url := api_api + "b2api/v2/b2_get_upload_url?bucketId=" + bucketId

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("Authorization", authorizationToken)

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("###[ERROR]request failed:%s\n\n", err.Error())
	}
	rspMap, err := ParseResponse(resp)
	jsonstr, err := json.Marshal(rspMap)
	if err != nil {
		fmt.Println("j2:" + err.Error())
		return
	}
	fmt.Printf("get upload file url:\nmap:%s\n", jsonstr)
	//return rspMap[""]
	//= rspMap["authorizationToken"].(string)
}

func uploadFileToB2(authToken string) {

	// newFile, err := os.Create("C:\\Users\\Yanjilong\\Desktop\\LinkinPark\\Lucian_25.jpg")defer newFile.Close()
	byteData, err := ioutil.ReadFile("C:\\Users\\Yanjilong\\Desktop\\LinkinPark\\Lucian_25.jpg")

	req, _ := http.NewRequest(http.MethodPost, uploadUrl, bytes.NewReader(byteData))

	req.Header.Set("Authorization", authToken)
	req.Header.Set("X-Bz-File-Name", "test1.jpg")
	req.Header.Set("Content-Type", "b2/x-auto")
	req.Header.Set("X-Bz-Content-Sha1", getFileSHA1(byteData))
	req.Header.Set("X-Bz-Info-Author", "lion")

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("###[ERROR]request failed:%s\n\n", err.Error())
	}
	rspMap, err := ParseResponse(resp)
	jsonstr, err := json.Marshal(rspMap)
	if err != nil {
		fmt.Println("j2:", err.Error())
		return
	}
	fmt.Printf("get upload file rsp:\nmap:%s\n", jsonstr)
}

func getFileSHA1(f []byte) string {
	h := sha1.New() // md5加密类似md5.New()
	h.Write([]byte(f))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来对现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串。
	fmt.Printf("%x\n", bs)
	//如果需要对另一个字符串加密，要么重新生成一个新的散列，要么一定要调用h.Reset()方法，不然生成的加密字符串会是拼接第一个字符串之后进行加密
	h.Reset() //重要！！！

	fmt.Printf("%s\n", fmt.Sprintf("%x", h.Sum(nil)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

/*
func CopyFileToB2(src, dst string) error {

	// b2_authorize_account
	client, err := b2.NewClient(ctx, app_key_id, app_key)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(ctx, buket_name)
	if err != nil {
		return err
	}

	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	obj := bucket.Object(dst)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, f); err != nil {
		w.Close()
		return err
	}
	return w.Close()
}

func downloadFile(ctx context.Context, bucket *b2.Bucket, downloads int, src, dst string) error {
	r := bucket.Object(src).NewReader(ctx)
	defer r.Close()

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	r.ConcurrentDownloads = downloads
	if _, err := io.Copy(f, r); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}*/
