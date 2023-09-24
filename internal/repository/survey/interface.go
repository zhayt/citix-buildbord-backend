package survey

import (
	"context"
	"innovatex-app/internal/models"
)

type Repository interface {
	GetList(ctx context.Context) ([]*models.SurveyInfo, error)
	Get(ctx context.Context, surveyID string) (*models.Survey, error)
	Save(ctx context.Context, survey *models.Survey) error
}
