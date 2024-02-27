package data_service

import (
	"context"
	"fmt"
	"github.com/Artonus/hermes/internal/util"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
	"path/filepath"
)

type S3Service struct {
	s3     *s3.Client
	bucket string
	region string
}

func NewS3Service(client s3.Client) (*S3Service, error) {
	bucket := os.Getenv("AWS_BUCKET")
	region := os.Getenv("AWS_REGION")

	return &S3Service{
		s3:     &client,
		bucket: bucket,
		region: region,
	}, nil
}

func (s3Service *S3Service) Fetch(netAppKey, targetDir string) error {
	fmt.Println("fetching the data")

	downloader := manager.NewDownloader(s3Service.s3)

	// List all objects in the bucket
	resp, err := s3Service.s3.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(s3Service.bucket),
	}, func(o *s3.Options) {
		o.Region = s3Service.region
	})

	if err != nil {
		return err
	}

	for _, obj := range resp.Contents {
		// Specify the local file path to save the object
		localFilePath := filepath.Join(targetDir, netAppKey, *obj.Key)
		errCreateDir := util.EnsurePathToFileExists(localFilePath)
		if errCreateDir != nil {
			return errCreateDir
		}
		fmt.Printf("Saving file to: %s", localFilePath)
		file, errFile := os.Create(localFilePath)
		if errFile != nil {
			return errFile
		}

		// Download the object
		_, errDownload := downloader.Download(context.TODO(), file, &s3.GetObjectInput{
			Bucket: aws.String(s3Service.bucket),
			Key:    obj.Key,
		})
		if errDownload != nil {
			fmt.Printf("Error downloading object %s: %v\n", *obj.Key, errDownload)
			return errDownload
		} else {
			fmt.Printf("Downloaded object: %s\n", localFilePath)
		}
		errClose := file.Close()
		if errClose != nil {
			return errClose
		}
	}

	return nil
}

func (s3Service *S3Service) Post(netAppKey, sourceDir string) error {
	uploader := manager.NewUploader(s3Service.s3)

	files, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(sourceDir, netAppKey, file.Name())
		fileData, openFileErr := os.Open(filePath)
		if openFileErr != nil {
			return openFileErr
		}
		fmt.Printf("Uploading file: %s\n", file.Name())
		_, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(s3Service.bucket),
			Key:    aws.String(file.Name()),
			Body:   fileData,
		})

		if uploadErr != nil {
			return uploadErr
		}
		errClose := fileData.Close()
		if errClose != nil {
			return errClose
		}
	}
	return nil
}
