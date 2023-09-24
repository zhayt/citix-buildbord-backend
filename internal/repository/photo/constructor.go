package photo

import (
	"innovatex-app/internal/config"
	"innovatex-app/internal/connection"
)

type viaRepository struct {
	*viaCloudinary
	*viaPostgres
}

func NewRepository(config *config.Config, connections *connection.Connection) Repository {
	return &viaRepository{
		viaCloudinary: newViaCloudinary(connections),
		viaPostgres:   newViaPostgres(config.Postgres, connections),
	}
}
