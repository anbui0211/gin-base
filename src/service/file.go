package service

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"os"
)

// FileDataInterface ...
type FileDataInterface interface {
	UploadProducts(ctx context.Context) error
}

// fileDataImpl ...
type fileDataImpl struct {
	awsService AWSService
}

// FileData ...
func FileData() FileDataInterface {
	return &fileDataImpl{}
}

// UploadProducts ...
func (s *fileDataImpl) UploadProducts(ctx context.Context) (err error) {
	// Config service
	config, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Println("Error while loading config: ", err)
		return
	}

	s.awsService.S3Client = s3.NewFromConfig(config)

	// Uploading
	fmt.Println("-- starting upload")
	var (
		bucketName = "kozo-dev"
		bucketKey  = "product"
		fileName   = "./assets/upload_file/product.txt"
	)

	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error while opening file: ", err)
		return
	}

	defer file.Close()

	_, err = s.awsService.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(bucketKey),
		Body:   file,
	})

	if err != nil {
		log.Println("Error while uploading file: ", err)
		return
	}

	fmt.Println("-- finish upload")
	return
}

type AWSService struct {
	S3Client *s3.Client
}
