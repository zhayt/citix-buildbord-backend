package service

import (
	"context"
	"go.uber.org/zap"
	"innovatex-app/internal/models"
	"innovatex-app/internal/repository"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, user *models.User) (*models.User, error)
}

type viaUserService struct {
	repository *repository.Repository
}

func newViaUserService(repository *repository.Repository) *viaUserService {
	return &viaUserService{repository: repository}
}

func (s *viaUserService) CreateUser(ctx context.Context, user *models.User) error {
	// TODO: validate user data

	userID, err := s.repository.User.Create(ctx, user)
	if err != nil {
		zap.S().Errorf("Creating user error: %s", err.Error())
		return err
	}

	zap.S().Infof("User with %d id created", userID)
	return nil
}

func (s *viaUserService) GetUser(ctx context.Context, user *models.User) (*models.User, error) {
	userFromDB, err := s.repository.User.Get(ctx, user.Number)
	if err != nil {
		zap.S().Errorf("Getting user error: %s", err.Error())
		return nil, err
	}

	// TODO: check user password etc

	return userFromDB, nil
}
