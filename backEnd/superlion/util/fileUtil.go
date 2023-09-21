//@program: superlion
//@author: yanjl
//@create: 2023-09-15 09:52
package util

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
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

// sm.ms图床：
var (
	smToken = "YYLX7IFbJQBwbHoHnctnK5Fe2iOV8U4n"

	uploadApi = "https://sm.ms/api/v2/upload"
)

// https://imgse.com/json，
// 路过图床
var (
	seUploadApi = "https://imgse.com/json"
	seAuthToken = "472518322270abe96d244b5b9aff9dc571c09d9b"
	seCookie    = "_ga=GA1.1.253708145.1695000390; PHPSESSID=4os147rn9gqhirtnr509fpmqnl; Hm_lvt_d26a1af7fee9e628d7d351346aa2f13b=1695000390,1695178823; KEEP_LOGIN=7Gres%3A3c4c6bb058872d18fc8350a7ada89665af100b60bb8c3ca4e474a9db87600396863e4fecd1e7b42dac5827edb5a5c2e461dc9f8413ed67796d19a7e7e56ee33bc180e372cd86fd77919c3c655b8e421ac6283dd3%3A1695179047; __gads=ID=ed9e864bd0a0e7f7-2216bf88eae30077:T=1695000391:RT=1695189883:S=ALNI_MbRvn333Bk2qk-KpvCvxvVqN5XYEA; __gpi=UID=00000c498d101dda:T=1695000391:RT=1695189883:S=ALNI_MYukOy2W0oKowktLyeHdqJ7_TiQ4g; Hm_lpvt_d26a1af7fee9e628d7d351346aa2f13b=1695189916; _ga_CZP2J5CMLW=GS1.1.1695189881.5.1.1695189916.0.0.0"
)

// nginx服务器
var (
	nginxApi = "http://148.100.77.194:8999/upload"
)

func UploadPicToImagse(file *multipart.File, fileName string) {

	// 异步：
	// go uploadFileToNginx(file, fileName)
	// path := "F:\\Xayah_37.png"
	// smFile, _ := os.Create(path)

	body := map[string]string{
		// "source":     "C:\\Users\\Yanjilong\\Desktop\\LinkinPark\\lol\\Thresh_3.jpg",
		"type":       "file",
		"action":     "upload",
		"auth_token": seAuthToken,
		"nsfw":       "0",
	}

	// bytesF, _ := json.Marshal(body)

	req, eor := newfileUploadRequest(seUploadApi, body, "source", file, fileName)

	req.Header.Set("Cookie", seCookie)

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

}

// 通过ftp 传输文件到nginx服务器 todo
func UploadFileToNginx(file *multipart.File, fileName string) string {
	req, eor := newfileUploadRequest(nginxApi, nil, "nFile", file, fileName)

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
		fmt.Println("###[ERROR]:do post nginx request error:", eor.Error())
	}
	mapStr, _ := ParseResponse(resp)

	fmt.Println("[INFO] post req to nginx save pic over :", resp.StatusCode, mapStr)

	str, _ := json.Marshal(mapStr)
	return string(str)
}

//定义一个创建文件目录的方法
func Mkdir(basePath string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, folderName)
	//使用mkdirall会创建多层级目录
	os.MkdirAll(folderPath, os.ModePerm)
	return folderPath
}

// 上传sm.ms
/*
func UploadPicsToSMMS() {

	path := "F:\\Xayah_37.png"
	// smFile, _ := os.Create(path)

	body := map[string]string{
		"format": "json",
	}

	// bytesF, _ := json.Marshal(body)

	req, eor := newfileUploadRequest(uploadApi, body, "smfile", path)
	// req, eor := RequestPost(uploadApi, body, "smfile", path)

	//req, eor := http.NewRequest("POST", uploadApi, nil)
	//body := &bytes.Buffer{}
	//writer := multipart.NewWriter(body)
	//part, _ := writer.CreateFormFile("smfile", filepath.Base(path))
	//file, eor := os.Open(path)
	//io.Copy(part, file)

	req.Header.Set("Authorization", smToken)

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
}*/

/*
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
}*/

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
