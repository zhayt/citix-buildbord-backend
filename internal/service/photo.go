package service

import (
	"context"
	"go.uber.org/zap"
	"innovatex-app/internal/models"
	"innovatex-app/internal/repository"
)

const timestampLayout = "2006-01-02"

type IPhotoService interface {
	SavePhoto(ctx context.Context, file string) error
	GetAllPhoto(ctx context.Context) ([]*models.PhotoInfo, error)
}

type viaPhotoService struct {
	repository *repository.Repository
}

func newViaPhotoService(repository *repository.Repository) *viaPhotoService {
	return &viaPhotoService{repository: repository}
}

func (s *viaPhotoService) SavePhoto(ctx context.Context, file string) error {
	zap.S().Info("Saving photo in s3")
	imageURL, err := s.repository.Photo.Save(ctx, file)
	if err != nil {
		zap.S().Errorf("Saving photo error: %s", err.Error())
		return err
	}

	zap.S().Info("Saving photo info")
	photoInfo := &models.PhotoInfo{ImageURL: imageURL}
	if err = s.repository.Photo.SaveInfo(ctx, photoInfo); err != nil {
		zap.S().Errorf("Saving photo info error: %s", err.Error())
		return err
	}

	return nil
}

func (s *viaPhotoService) GetAllPhoto(ctx context.Context) ([]*models.PhotoInfo, error) {
	return s.repository.Photo.GetAll(ctx)
}
