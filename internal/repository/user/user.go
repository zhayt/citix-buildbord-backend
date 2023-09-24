package user

import (
	"context"
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

func newViaPostgres(config *config.Postgres, connection *connection.Connection) *viaPostgres {
	return &viaPostgres{
		table:   config.UserTable,
		timeout: config.Timeout,
		db:      connection.PostgresClient,
	}
}

func (r *viaPostgres) Create(ctx context.Context, user *models.User) (int, error) {
	qr := fmt.Sprintf(`INSERT INTO %s (phone_number, password) VALUES ($1, $2) RETURNING id`, r.table)

	var userID int
	if err := r.db.GetContext(ctx, &userID, qr, user.Number, user.Password); err != nil {
		zap.S().Errorf("Creating user error: %s", err.Error())
		return 0, err
	}

	return userID, nil
}

func (r *viaPostgres) Get(ctx context.Context, phoneNumber string) (*models.User, error) {
	qr := fmt.Sprintf(`SELECT phone_number, password FROM %s WHERE phone_number = $1`, r.table)

	user := &models.User{}
	if err := r.db.GetContext(ctx, user, qr, phoneNumber); err != nil {
		zap.S().Errorf("Getting user error: %s", err.Error())
		return nil, err
	}

	return user, nil
}
