package s3adapter

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/pkg/errors"
)

// Upload file to s3
func (c *S3) Upload(bucketName, key string, file io.ReadSeeker) error {
	svc, err := c.connect()
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = svc.PutObject(context.Background(), &s3.PutObjectInput{
		Body:   file,
		Bucket: &bucketName,
		Key:    &key,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Download file from s3
func (c *S3) Download(bucketName, key string) (string, error) {
	svc, err := c.connect()
	if err != nil {
		return "", errors.WithStack(err)
	}

	f, err := os.CreateTemp("", key)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer f.Close()

	c.logger.Info().Msgf("Download from S3,  Bucket: %v, Key: %v", bucketName, key)

	getObj, err := svc.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", errors.WithStack(err)
	}

	// Read value (getObj) and write to file (f)
	_, err = io.Copy(f, getObj.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}

	c.logger.Info().Msgf("tmpFileName: %v", f.Name())
	return f.Name(), nil
}

// GetUploadURL get url upload file to s3
func (c *S3) GetUploadURL(bucketName, key string) (string, error) {
	_, err := c.connect()
	if err != nil {
		return "", errors.WithStack(err)
	}

	req, err := c.presigner.PresignPutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = S3PresignedURLExpires
	})
	if err != nil {
		return "", errors.WithStack(err)
	}

	return req.URL, nil
}

// GetDownloadURL get url download file from s3
func (c *S3) GetDownloadURL(bucketName, key string) (string, error) {
	_, err := c.connect()
	if err != nil {
		return "", errors.WithStack(err)
	}
	req, err := c.presigner.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = S3PresignedURLExpires
	})
	if err != nil {
		return "", errors.WithStack(err)
	}
	return req.URL, nil
}

// Delete file in s3
func (c *S3) Delete(bucketName, key string) error {
	svc, err := c.connect()
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = svc.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
