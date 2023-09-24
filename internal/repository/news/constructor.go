package news

import "innovatex-app/internal/config"

func NewRepository(source *config.Source) Repository {
	return newViaHTTP(source)
}
