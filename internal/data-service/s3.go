package data_service

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service struct {
	s3 *s3.Client
}

func NewS3Service(client s3.Client, bucket string) (*S3Service, error) {

	_, err := client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		return nil, err
	}

	return &S3Service{s3: &client}, nil
}

func (*S3Service) Fetch(dir string) error {
	fmt.Println("fetch called")
	return nil
}
