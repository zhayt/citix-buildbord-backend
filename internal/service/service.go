package service

import (
	"innovatex-app/internal/repository"
)

type Service struct {
	User   IUserService
	News   INewsService
	Survey ISurveyService
	Photo  IPhotoService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		User:   newViaUserService(repository),
		News:   newViaNewsService(repository),
		Survey: newViaSurveyService(repository),
		Photo:  newViaPhotoService(repository),
	}
}
