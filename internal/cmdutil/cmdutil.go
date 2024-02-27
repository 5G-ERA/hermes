package cmdutil

import (
	"context"
	"errors"
	"fmt"
	"github.com/Artonus/hermes/internal/data-service"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

func CreateFetchClient() (data_service.DataFetcher, error) {
	return create()
}

func CreatePostClient() (data_service.DataPoster, error) {
	return create()
}

func create() (*data_service.S3Service, error) {
	s3Client, err := createS3Client()
	if err != nil {
		return nil, err
	}
	s3Service, err := data_service.NewS3Service(*s3Client)
	if err != nil {
		return nil, err
	}
	return s3Service, nil
}

func createS3Client() (*s3.Client, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")

	if accessKey == "" || secretKey == "" {
		return nil, errors.New("AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set")
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")))

	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return nil, err
	}
	if isValidRegionAndBucket(cfg, region, bucket) == false {
		return nil, errors.New("AWS_REGION must be set to a valid AWS region")
	}

	sss := s3.NewFromConfig(cfg)
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
