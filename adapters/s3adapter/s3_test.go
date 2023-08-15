package s3adapter_test

import (
	"gin-base/internal/util/env"
	"os"
	"testing"

	"gin-base/adapters/s3adapter"
	"github.com/stretchr/testify/assert"
)

func buildTestS3Client() *s3adapter.S3 {
	// LocalStackを指定する
	os.Setenv("AWS_ENDPOINT_URL", "http://localhost:4566")
	os.Setenv("AWS_REGION", "ap-southeast-1")
	os.Setenv("S3_BUCKET_NAME", "kozo-dev")
	// Need Static Credentials
	//	os.Setenv("AWS_ACCESS_KEY_ID", "dummy")
	//	os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy")
	return s3adapter.NewS3()
}

func TestS3_GetDownloadURL(t *testing.T) {
	client := buildTestS3Client()

	bucket := env.S3BucketName()
	key := "exampleKey"
	url, err := client.GetDownloadURL(bucket, key)
	println(url)
	assert.NoError(t, err)
}

func TestS3_GetUploadURL(t *testing.T) {
	client := buildTestS3Client()

	bucket := env.S3BucketName()
	key := "exampleKey"
	url, err := client.GetUploadURL(bucket, key)
	println(url)
	assert.NoError(t, err)
}

func TestS3_Upload(t *testing.T) {
	client := buildTestS3Client()

	bucket := env.S3BucketName()
	key := "exampleKey"
	file, err := os.Open("../../assets/testdata/product.csv")
	assert.NoError(t, err)

	err = client.Upload(bucket, key, file)
	assert.NoError(t, err)
}

func TestS3_Download(t *testing.T) {
	client := buildTestS3Client()

	bucket := env.S3BucketName()
	key := "exampleKey"
	tmpFileName, err := client.Download(bucket, key)
	println(tmpFileName)
	assert.NoError(t, err)
}

func TestS3_Delete(t *testing.T) {
	client := buildTestS3Client()
	bucket := env.S3BucketName()
	key := "exampleKey"

	err := client.Delete(bucket, key)
	assert.NoError(t, err)
}
