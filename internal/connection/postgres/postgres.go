package postgres

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"innovatex-app/internal/config"
)

func Dial(postgres *config.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", postgres.MakeConnectionURL())
	if err != nil {
		return nil, fmt.Errorf("connot open db: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), postgres.Timeout)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("cannot connect to db: %w", err)
	}

	return db, nil
}
