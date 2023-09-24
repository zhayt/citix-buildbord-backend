package wasabi

import (
	"innovatex-app/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func NewDownloader(config *config.Wasabi) *s3manager.Downloader {
	sessions := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(config.AccessKeyID, config.SecretKey, ""),
		Endpoint:    aws.String(config.Endpoint),
		Region:      aws.String(config.Region),
	}))

	return s3manager.NewDownloader(sessions)
}

func NewUploader(config *config.Wasabi) *s3manager.Uploader {
	sessions := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(config.AccessKeyID, config.SecretKey, ""),
		Endpoint:    aws.String(config.Endpoint),
		Region:      aws.String(config.Region),
	}))

	return s3manager.NewUploader(sessions)
}
