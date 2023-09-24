package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"innovatex-app/internal/config"
	"innovatex-app/internal/service"
	"innovatex-app/internal/transport/http/handler"
	"time"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	App             *echo.Echo
	cfg             *config.Config
	handler         *handler.Handler
	notify          chan error
	shutdownTimeout time.Duration
}

func NewServer(cfg *config.Config, service *service.Service) *Server {
	handler := handler.NewHandler(cfg.App, service)

	srv := &Server{
		cfg:             cfg,
		handler:         handler,
		shutdownTimeout: _defaultShutdownTimeout,
		notify:          make(chan error, 1),
	}

	return srv
}

func (s *Server) BuildingEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	return e
}

func (s *Server) Start() {
	s.App = s.BuildingEngine()
	s.SetUpRoute()
	go func() {
		s.notify <- s.App.Start(fmt.Sprintf("%s:%s", s.cfg.App.Host, s.cfg.App.Port))
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.App.Shutdown(ctx)
}
