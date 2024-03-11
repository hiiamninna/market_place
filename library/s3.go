package library

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	ID        string
	SecretKey string
	BaseUrl   string
}

func NewS3Config(config S3) (*s3manager.Uploader, error) {
	session, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(config.ID, config.SecretKey, ""),
	})
	if err != nil {
		return &s3manager.Uploader{}, fmt.Errorf("open con : %w", err)
	}

	return s3manager.NewUploader(session), nil
}

func (storage *S3) UploadFile(bucketName string, data io.Reader, FileName string, uploader *s3manager.Uploader) (string, error) {
	resp, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(FileName),
		Body:   data,
	})
	if err != nil {
		return "", fmt.Errorf("fail upload file : %w", err)
	}

	return resp.Location, nil
}
