package aliyun_test

import (
	"os"
	"testing"

	"gitee.com/os4top_admin/cloud-station-g7/store"
	"gitee.com/os4top_admin/cloud-station-g7/store/aliyun"
	"github.com/stretchr/testify/assert"
)

var (
	uploader store.Uploader
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

func TestUpload(t *testing.T) {

	should := assert.New(t)

	err := uploader.Upload(BucketName, "test.txt", "store_test.go")

	if should.NoError(err) {
		t.Log("upload ok")
	}
}

func TestUploadError(t *testing.T) {
	should := assert.New(t)

	err := uploader.Upload(BucketName, "test.txt", "store_testxxx.go")
	should.Error(err, "open store_testxxx.go: The system cannot find the file specified.")
}

func init() {
	ali, err := aliyun.NewDefaultAliOssStore()
	if err != nil {
		panic(err)
	}
	uploader = ali
}
