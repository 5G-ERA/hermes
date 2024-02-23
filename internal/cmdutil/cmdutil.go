package cmdutil

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func CreateAwsSession() (*s3.S3, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	if accessKey == "" || secretKey == "" {
		return nil, errors.New("AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set")
	}
	if isValidRegion(region) == false {
		return nil, errors.New("AWS_REGION must be set to a valid AWS region")
	}

	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: &region,
		}))
	sss := s3.New(sess)
	return sss, nil
}

func isValidRegion(region string) bool {

	allRegions := endpoints.AwsPartition().Regions()

	// Check if the provided region is in the list of valid regions
	for _, r := range allRegions {
		if r.ID() == region {
			return true
		}
	}

	return false
}
