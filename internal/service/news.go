package service

import (
	"context"
	"innovatex-app/internal/models"
	"innovatex-app/internal/repository"
)

type INewsService interface {
	GetNews(ctx context.Context) (*models.News, error)
}

type viaNewsService struct {
	repository *repository.Repository
}

func newViaNewsService(repository *repository.Repository) *viaNewsService {
	return &viaNewsService{repository: repository}
}

func (s *viaNewsService) GetNews(ctx context.Context) (*models.News, error) {
	return s.repository.News.GetAll(ctx)
}
