package repository

import (
	"innovatex-app/internal/config"
	"innovatex-app/internal/connection"
	"innovatex-app/internal/repository/news"
	"innovatex-app/internal/repository/photo"
	"innovatex-app/internal/repository/survey"
	"innovatex-app/internal/repository/user"
)

type Repository struct {
	User   user.Repository
	News   news.Repository
	Survey survey.Repository
	Photo  photo.Repository
}

func NewRepository(config *config.Config, connection *connection.Connection) *Repository {
	return &Repository{
		User:   user.NewRepository(config.Postgres, connection),
		News:   news.NewRepository(config.Source),
		Survey: survey.NewRepository(config, connection),
		Photo:  photo.NewRepository(config, connection),
	}
}
