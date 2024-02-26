package cmdutil

import (
	"context"
	"errors"
	"fmt"
	"github.com/Artonus/hermes/internal/data-service"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

func CreateFetchClient() (data_service.DataFetcher, error) {
	s3Client, err := createS3Client()
	if err != nil {
		return nil, err
	}
	return s3Client, nil
}
func createS3Client() (*s3.Client, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	if accessKey == "" || secretKey == "" {
		return nil, errors.New("AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set")
	}
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return nil, err
	}
	if isValidRegion(cfg, region) == false {
		return nil, errors.New("AWS_REGION must be set to a valid AWS region")
	}

	sss := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = region
	})
	return sss, nil
}

func isValidRegion(cfg aws.Config, region string) bool {
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = region
	})

	// Use the client to check if the region is valid
	_, err := s3Client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	return err == nil
}
