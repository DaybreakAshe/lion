// @program: superlion
// @author: yanjl
// @create: 2024-08-15 09:23
package util

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
	"mime/multipart"
)

const (
	// r2-s3 相关配置
	accessKey  = "12dbc3fa9a395358d77b3fbafbb1d6cf"
	secretKey  = "92f4c7c8817058dcf562f26406cd451b9c90f7bc06068dca55c9cb76af016cb7"
	region     = "auto"
	endpoint   = "https://4dfc4b4e437f18f3aa0e99a3900425c8.r2.cloudflarestorage.com/lion"
	bucket     = "lion"
	tokenValue = "p5cYtAaCFgeQhM2iIpoVURlwX7b4sWJs-IbI8pdN"
	// 路径相关配置
	uploadDir    = "tmp"
	fileKeyPre   = uploadDir + "/lion/"
	accessUrlPre = "https://lion.yanxue.eu.org/"
)

// S3Client wraps the AWS S3 client
type S3Client struct {
	client *s3.Client
	bucket string
}

// NewS3Client creates a new S3Client instance
func NewS3Client() *S3Client {

	// 配置 区域、key
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := s3.NewFromConfig(cfg)
	return &S3Client{
		client: client,
		bucket: bucket,
	}
}

// UploadFile uploads a file to the specified S3 bucket
func (s *S3Client) UploadFile(file multipart.File, filename string, contentType string) string {
	// s.client.

	point := endpoint

	// 函数选项模式 ，第三个参数
	output, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(fileKeyPre + filename),
		Body:        file,
		ACL:         types.ObjectCannedACLPrivate,
		ContentType: &contentType,
		//  : s3.EndpointResolverFromURL(endpoint),
	}, func(o *s3.Options) {
		o.BaseEndpoint = &point
	})
	if err != nil {
		log.Printf("upload to s3 error:%v", err.Error())

		return ""
	}

	// 打印上传结果
	log.Printf("upload to s3 result:%v", output.VersionId)

	return *output.VersionId
}
