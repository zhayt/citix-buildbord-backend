package photo

import (
	"context"
	"innovatex-app/internal/models"
)

type Repository interface {
	SaveInfo(ctx context.Context, file *models.PhotoInfo) error
	Save(ctx context.Context, file string) (string, error)
	GetAll(ctx context.Context) ([]*models.PhotoInfo, error)
}
