package cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"innovatex-app/internal/config"
)

func NewCloudinary(config *config.Cloudinary) (*cloudinary.Cloudinary, error) {
	cloudinaryClient, err := cloudinary.NewFromParams(config.CloudName, config.APIKey, config.APISecret)
	if err != nil {
		return nil, err
	}

	return cloudinaryClient, err
}
