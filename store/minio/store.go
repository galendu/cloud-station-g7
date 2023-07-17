package minio

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/galendu/cloud-station-g7/store"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
	useSSL       bool
}

var (
	_ store.Uploader = (*MinOssStore)(nil)
)

func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint, access_key access_secret has one empty")
	}

	return nil
}

type MinOssStore struct {
	client *minio.Client
	ctx    context.Context
}

func NewMinOssStore(opts *Options) (*MinOssStore, error) {
	c, err := minio.New(opts.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(opts.AccessKey, opts.AccessSecret, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	return &MinOssStore{
		client: c,
	}, nil
}
func NewDefaultMinOssStore() (*MinOssStore, error) {
	return NewMinOssStore(&Options{
		Endpoint:     "192.168.3.5:32098",
		AccessKey:    "HqfF0AioMtHh8pzrofXe",
		AccessSecret: "WJYCxSi4CkTRc8KRtTlDVew3MW4RZlfADPUsqIlQ",
	})
}

func (m *MinOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	// 1.获取buckets
	fmt.Println("开始获取buckets")
	bucket, err := m.client.ListBuckets(context.Background())
	if err != nil {
		fmt.Println("获取buckets失败")
		fmt.Println(err)
		return err
	}
	for _, i := range bucket {
		fmt.Println(i)
	}
	// // 2.获取bucket对象
	// m.client.BucketExists(m.ctx,bucketName)
	// 3.上传文件到该bucket
	uploadInfo, err := m.client.FPutObject(context.Background(), bucketName, objectKey, fileName, minio.PutObjectOptions{
		ContentType: "application/txt",
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\"main.go\"")

	presignedURL, err := m.client.PresignedGetObject(context.Background(), bucketName, fileName, time.Second*24*60*60, reqParams)
	if err != nil {
		fmt.Println("获取下载链接失败")
		return err
	}

	fmt.Println("Successfully generated presigned URL", presignedURL)
	// 4.打印下载链接
	// downloadURL,err:=
	return nil
}
