package main

import (
	"fmt"
	"go.uber.org/zap"
	"innovatex-app/internal/config"
	"innovatex-app/internal/connection"
	"innovatex-app/internal/logger"
	"innovatex-app/internal/repository"
	"innovatex-app/internal/service"
	"innovatex-app/internal/transport/http"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	var once sync.Once
	once.Do(config.PrepareEnv)

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("getting config error: %s", err.Error())
	}

	if err = logger.InitZapLogger(cfg.App); err != nil {
		return fmt.Errorf("initing logger error: %s", err.Error())
	}
	defer logger.Sync()

	connections, err := connection.NewConnection(cfg)
	if err != nil {
		return fmt.Errorf("getting connection error: %s", err.Error())
	}
	defer connections.Close()

	repo := repository.NewRepository(cfg, connections)
	serv := service.NewService(repo)
	delivery := http.NewServer(cfg, serv)

	zap.S().Info("Start server")
	delivery.Start()

	// grace full shutdown
	osSignCh := make(chan os.Signal, 1)
	signal.Notify(osSignCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-osSignCh:
		zap.S().Infof("signal accepted: %s", s.String())
	case err = <-delivery.Notify():
		zap.S().Errorf("server closing %s", err.Error())
	}

	if err = delivery.Shutdown(); err != nil {
		return fmt.Errorf("error while shutting down server: %s", err)
	}

	return nil
}
