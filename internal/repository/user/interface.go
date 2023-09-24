package user

import (
	"context"
	"innovatex-app/internal/models"
)

type Repository interface {
	Create(ctx context.Context, user *models.User) (int, error)
	Get(ctx context.Context, phoneNumber string) (*models.User, error)
}
