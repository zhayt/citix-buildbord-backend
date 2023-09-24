package user

import (
	"innovatex-app/internal/config"
	"innovatex-app/internal/connection"
)

func NewRepository(config *config.Postgres, connection *connection.Connection) Repository {
	return newViaPostgres(config, connection)
}
