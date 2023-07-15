package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	client *oss.Client
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

// 测试阿里云OssSDK BucketList接口
func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

// 测试阿里云OssSDK PutObjectFromFile接口
func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("mydir/oss_test1.go", "oss_test.go")
	if err != nil {
		t.Log(err)
	}
}

// 初始化一个Oss Client,等下给其他所有测试用例使用
func init() {

	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	fmt.Println(OssEndpoint)
	if err != nil {
		fmt.Println("连接错误")
		panic(err)
	}
	client = c
}
