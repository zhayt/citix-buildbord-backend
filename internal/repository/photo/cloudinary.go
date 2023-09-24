package photo

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.uber.org/zap"
	"innovatex-app/internal/connection"
)

type viaCloudinary struct {
	cld *cloudinary.Cloudinary
}

func newViaCloudinary(connection *connection.Connection) *viaCloudinary {
	return &viaCloudinary{
		cld: connection.CloudinaryClient,
	}
}

func (r *viaCloudinary) Save(ctx context.Context, file string) (string, error) {
	//imageBinary, err := base64.StdEncoding.DecodeString(file)
	//if err != nil {
	//	zap.S().Errorf("Converting base64 error: %s", err.Error())
	//	return "", err
	//}

	r.cld.Config.API.ChunkSize = 6000000
	resp, err := r.cld.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		zap.S().Errorf("Uploading error: %s", err.Error())
		return "", err
	}

	return resp.URL, nil
}
