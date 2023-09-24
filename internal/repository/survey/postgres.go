package survey

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"innovatex-app/internal/config"
	"innovatex-app/internal/connection"
	"innovatex-app/internal/models"
	"time"
)

type viaPostgres struct {
	table   string
	timeout time.Duration
	db      *sqlx.DB
}

func newViaPostgres(postgres *config.Postgres, connection *connection.Connection) *viaPostgres {
	return &viaPostgres{
		table:   postgres.SurveyTable,
		timeout: postgres.Timeout,
		db:      connection.PostgresClient,
	}
}

func (r *viaPostgres) GetList(ctx context.Context) ([]*models.SurveyInfo, error) {
	qr := fmt.Sprintf("SELECT survey_id, survey_title, company_logo FROM %s GROUP BY survey_id, survey_title, company_logo", r.table)

	var surveyInfo []*models.SurveyInfo
	if err := r.db.SelectContext(ctx, &surveyInfo, qr); err != nil {
		zap.S().Errorf("Getting survey list error: %s", err.Error())
		return nil, err
	}

	return surveyInfo, nil
}

func (r *viaPostgres) Get(ctx context.Context, surveyID string) (*models.Survey, error) {
	qr := fmt.Sprintf(`SELECT survey_id, survey_title, company_logo, questions FROM %s WHERE survey_id = $1`, r.table)

	survey := &models.Survey{}
	var questionsJson []byte
	if err := r.db.QueryRowContext(ctx, qr, surveyID).Scan(&survey.SurveyID, &survey.SurveyTitle, &survey.CompanyLogo, &questionsJson); err != nil {
		zap.S().Errorf("Getting survey error: %s", err.Error())
		return nil, err
	}

	var questions []*models.Question
	if err := json.Unmarshal(questionsJson, &questions); err != nil {
		zap.S().Errorf("Parsing questions error: %s", err.Error())
		return nil, err
	}

	survey.Questions = questions
	return survey, nil
}
