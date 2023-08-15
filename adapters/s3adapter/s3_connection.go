package s3adapter

import (
	"time"

	"gin-base/internal/util/plogger"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type S3 struct {
	s3        *s3.Client
	presigner *s3.PresignClient
	logger    *zerolog.Logger
}

func NewS3() *S3 {
	return &S3{}
}

const S3PresignedURLExpires = 10 * time.Second

func (c *S3) connect() (*s3.Client, error) {
	if c.logger == nil {
		var err error
		c.logger, _, err = plogger.GetLogger()
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if c.s3 == nil {
		builder := NewS3ClientBuilder(c.logger)
		client, err := builder.Build()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		c.s3 = client
	}

	if c.presigner != nil {
		c.presigner = s3.NewPresignClient(c.s3)
	}

	return c.s3, nil
}
