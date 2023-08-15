package awsconfig

import (
	"context"

	"gin-base/internal/util/env"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/pkg/errors"
)

type AWSConfig struct {
	awsConfig *aws.Config
}

func GetConfig() (*aws.Config, error) {
	a := &AWSConfig{}

	// When had config
	if a.awsConfig != nil {
		return a.awsConfig, nil
	}

	// When no have config
	err := a.buildConfig()
	if err != nil {
		return nil, err
	}

	return a.awsConfig, nil
}

func (a *AWSConfig) buildConfig() error {
	awsRegion := env.AWSRegion()
	awsEndpoint := env.AWSEndPointURL()

	// Custom recover
	customRecover := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...any) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{PartitionID: "aws", URL: awsEndpoint, SigningRegion: awsRegion}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	// Get Config
	awsConfig, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolverWithOptions(customRecover),
	)
	if err != nil {
		return errors.WithStack(err)
	}

	a.awsConfig = &awsConfig
	return nil
}
