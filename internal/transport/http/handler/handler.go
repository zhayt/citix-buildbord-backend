package handler

import (
	"innovatex-app/internal/config"
	"innovatex-app/internal/service"
	"time"
)

type Handler struct {
	service *service.Service
	timeout time.Duration
}

func NewHandler(app *config.App, service *service.Service) *Handler {
	return &Handler{
		service: service,
		timeout: app.Timeout,
	}
}
