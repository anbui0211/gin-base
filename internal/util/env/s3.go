package env

import "os"

func AWSEndPointURL() string {
	return os.Getenv("AWS_ENDPOINT_URL")
}

func AWSRegion() string {
	return os.Getenv("AWS_REGION")
}

func S3BucketName() string {
	return os.Getenv("S3_BUCKET_NAME")
}
