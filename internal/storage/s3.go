package storage

import (
	"context"
	"file-sharing-backend/pkg/config"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
	Bucket string
}

func NewS3Client(cfg *config.Config) *S3Client {
	awsCfg, _ := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{URL: cfg.AWSAccessKey}, nil
		})),
		config.WithCredentialsProvider(aws.Credentials{
			AccessKeyID: cfg.AWSAccessKey, SecretAccessKey: cfg.AWSSecretKey,
		}),
	)
	return &S3Client{
		Client: s3.NewFromConfig(awsCfg),
		Bucket: cfg.AWSBucket,
	}
}

func (s *S3Client) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fileHeader.Filename),
		Body:   file,
	})
	return fmt.Sprintf("%s/%s", s.Bucket, fileHeader.Filename), err
}
