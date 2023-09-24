package survey

import (
	"innovatex-app/internal/config"
	"innovatex-app/internal/connection"
)

type viaRepository struct {
	*viaPostgres
	*viaClickhouse
}

func NewRepository(config *config.Config, connection *connection.Connection) Repository {
	return &viaRepository{
		viaPostgres:   newViaPostgres(config.Postgres, connection),
		viaClickhouse: newViaClickhouse(),
	}
}
