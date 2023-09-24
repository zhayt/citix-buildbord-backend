package connection

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	cloude "github.com/cloudinary/cloudinary-go/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"innovatex-app/internal/config"
	"innovatex-app/internal/connection/cloudinary"
	"innovatex-app/internal/connection/postgres"
	"innovatex-app/internal/connection/wasabi"
)

type Connection struct {
	PostgresClient   *sqlx.DB
	WasabiDownloader *s3manager.Downloader
	WasabiUploader   *s3manager.Uploader
	CloudinaryClient *cloude.Cloudinary
}

func NewConnection(config *config.Config) (*Connection, error) {
	postgresClient, err := postgres.Dial(config.Postgres)
	if err != nil {
		return nil, err
	}
	wasabiDownloader := wasabi.NewDownloader(config.Wasabi)
	wasabiUploader := wasabi.NewUploader(config.Wasabi)
	cloudinaryClient, err := cloudinary.NewCloudinary(config.Cloudinary)
	if err != nil {
		return nil, err
	}

	return &Connection{
		PostgresClient:   postgresClient,
		WasabiDownloader: wasabiDownloader,
		WasabiUploader:   wasabiUploader,
		CloudinaryClient: cloudinaryClient,
	}, nil
}

func (c *Connection) Close() error {
	if err := c.PostgresClient.Close(); err != nil {
		zap.S().Errorf("Closing postges connection error: %s", err.Error())
		return err
	}
	return nil
}
