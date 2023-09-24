package service

import (
	"context"
	"innovatex-app/internal/models"
	"innovatex-app/internal/repository"
)

type ISurveyService interface {
	GetSurveyList(ctx context.Context) ([]*models.SurveyInfo, error)
	GetSurvey(ctx context.Context, surveyID string) (*models.Survey, error)
	SaveSurvey(ctx context.Context, survey *models.Survey) error
}

type viaSurveyService struct {
	repository *repository.Repository
}

func newViaSurveyService(repository *repository.Repository) *viaSurveyService {
	return &viaSurveyService{repository: repository}
}

func (s *viaSurveyService) GetSurveyList(ctx context.Context) ([]*models.SurveyInfo, error) {
	return s.repository.Survey.GetList(ctx)
}

func (s *viaSurveyService) GetSurvey(ctx context.Context, surveyID string) (*models.Survey, error) {
	return s.repository.Survey.Get(ctx, surveyID)
}

func (s *viaSurveyService) SaveSurvey(ctx context.Context, survey *models.Survey) error {
	return s.repository.Survey.Save(ctx, survey)
}
