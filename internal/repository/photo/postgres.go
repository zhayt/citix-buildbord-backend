package photo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"innovatex-app/internal/config"
	"innovatex-app/internal/connection"
	"innovatex-app/internal/models"
)

type viaPostgres struct {
	table string
	db    *sqlx.DB
}

func newViaPostgres(config *config.Postgres, connection *connection.Connection) *viaPostgres {
	return &viaPostgres{
		table: config.PhotoTable,
		db:    connection.PostgresClient,
	}
}

func (r *viaPostgres) SaveInfo(ctx context.Context, file *models.PhotoInfo) error {
	qr := fmt.Sprint(`INSERT INTO photo_info (image_url) VALUES ($1)`)

	if _, err := r.db.ExecContext(ctx, qr, file.ImageURL); err != nil {
		zap.S().Errorf("Saving photo info error: %s", err.Error())
		return err
	}

	return nil
}

func (r *viaPostgres) GetAll(ctx context.Context) ([]*models.PhotoInfo, error) {
	qr := fmt.Sprint(`SELECT image_url, created_at FROM photo_info`)

	var images []*models.PhotoInfo
	if err := r.db.SelectContext(ctx, &images, qr); err != nil {
		zap.S().Errorf("Getting all images error: %s", err.Error())
		return nil, err
	}

	return images, nil
}
