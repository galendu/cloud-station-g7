package minio

import (
	"context"
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Options struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
	useSSL       bool
}

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
		Secure: opts.useSSL,
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
		Endpoint:     os.Getenv("ALI_OSS_ENDPOINT"),
		AccessKey:    os.Getenv("ALI_AK"),
		AccessSecret: os.Getenv("ALI_SK"),
	})
}

func (m *MinOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	// 1.获取buckets
	bucket, err := m.client.ListBuckets(m.ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := range bucket {
		fmt.Println(i)
	}
	// // 2.获取bucket对象
	// m.client.BucketExists(m.ctx,bucketName)
	// 3.上传文件到该bucket
	uploadInfo, err := m.client.FPutObject(m.ctx, bucketName, objectKey, fileName, minio.PutObjectOptions{
		ContentType: "application/txt",
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Successfully uploaded object: ", uploadInfo)

	// 4.打印下载链接
	// downloadURL,err:=
	return nil
}
