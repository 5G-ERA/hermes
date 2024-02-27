package cmdutil

import (
	"context"
	"errors"
	"fmt"
	hermesConfig "github.com/Artonus/hermes/internal/config"
	"github.com/Artonus/hermes/internal/data-service"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func CreateFetchClient(cfg *hermesConfig.HermesConfig) (data_service.DataFetcher, error) {
	return create(cfg)
}

func CreatePostClient(cfg *hermesConfig.HermesConfig) (data_service.DataPoster, error) {
	return create(cfg)
}
func CreateDeleteClient(cfg *hermesConfig.HermesConfig) (data_service.DataCleaner, error) {
	return create(cfg)
}

func create(cfg *hermesConfig.HermesConfig) (*data_service.S3Service, error) {
	s3Client, err := createS3Client(cfg)
	if err != nil {
		return nil, err
	}
	s3Service, err := data_service.NewS3Service(*s3Client)
	if err != nil {
		return nil, err
	}
	return s3Service, nil
}

func createS3Client(cfg *hermesConfig.HermesConfig) (*s3.Client, error) {
	if cfg.AwsAccessKeyId == "" || cfg.AwsAccessSecretKey == "" {
		return nil, errors.New("AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set")
	}
	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(cfg.AwsRegion),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AwsAccessKeyId, cfg.AwsAccessSecretKey, "")))

	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return nil, err
	}
	if isValidRegionAndBucket(awsCfg, cfg.AwsRegion, cfg.AwsBucket) == false {
		return nil, errors.New("AWS_REGION must be set to a valid AWS region")
	}

	sss := s3.NewFromConfig(awsCfg)
	return sss, nil
}

func isValidRegionAndBucket(cfg aws.Config, region, bucketName string) bool {
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = region
	})

	// Use the client to check if the region is valid
	buckets, err := s3Client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	bucketExists := false
	for _, bucket := range buckets.Buckets {
		if *bucket.Name == bucketName {
			bucketExists = true
			break
		}
		//fmt.Println(*bucket.Name)
	}
	return err == nil && bucketExists
}
