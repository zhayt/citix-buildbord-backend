package news

import (
	"context"
	"innovatex-app/internal/models"
)

type Repository interface {
	GetAll(ctx context.Context) (*models.News, error)
}
