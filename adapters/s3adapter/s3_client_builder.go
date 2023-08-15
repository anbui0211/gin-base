package s3adapter

import (
	"gin-base/adapters/awsconfig"
	"gin-base/internal/util/env"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type S3ClientBuilder struct {
	logger *zerolog.Logger
}

func NewS3ClientBuilder(logger *zerolog.Logger) *S3ClientBuilder {
	return &S3ClientBuilder{logger: logger}
}

func (s *S3ClientBuilder) Build() (*s3.Client, error) {
	cfg, err := awsconfig.GetConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if env.AWSEndPointURL() != "" {
		return s3.NewFromConfig(*cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		}), nil
	}

	return s3.NewFromConfig(*cfg), nil
}
