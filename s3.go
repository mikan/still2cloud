package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func putS3Object(config Config, path string, content []byte) error {
	awsConfig, err := newAWSConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("[DEST] failed to load AWS configuration: %w", err)
	}
	_, err = newS3Client(config, awsConfig).PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &config.Destination.Bucket,
		Key:    &path,
		Body:   bytes.NewReader(content),
	})
	if err != nil {
		return fmt.Errorf("[DEST] failed to put object: %w", err)
	}
	return nil
}

func newAWSConfig(ctx context.Context, config Config) (aws.Config, error) {
	var awsConfig aws.Config
	var err error
	if config.Destination.AccessKeyID != "" && config.Destination.SecretAccessKey != "" {
		cp := credentials.NewStaticCredentialsProvider(config.Destination.AccessKeyID, config.Destination.SecretAccessKey, "")
		if config.Destination.Region != "" {
			awsConfig, err = awsconfig.LoadDefaultConfig(ctx, awsconfig.WithCredentialsProvider(cp), awsconfig.WithRegion(config.Destination.Region))
		} else {
			awsConfig, err = awsconfig.LoadDefaultConfig(ctx, awsconfig.WithCredentialsProvider(cp))
		}
	} else if config.Destination.Region != "" {
		awsConfig, err = awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(config.Destination.Region))
	} else {
		awsConfig, err = awsconfig.LoadDefaultConfig(ctx)
	}
	return awsConfig, err
}

func newS3Client(config Config, awsConfig aws.Config) *s3.Client {
	if config.Destination.Endpoint != "" {
		return s3.NewFromConfig(awsConfig, func(o *s3.Options) { o.BaseEndpoint = &config.Destination.Endpoint })
	}
	return s3.NewFromConfig(awsConfig)
}
