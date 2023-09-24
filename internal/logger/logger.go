package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"innovatex-app/internal/config"
	"log"
	"time"
)

func InitZapLogger(app *config.App) error {
	var loggerConfig zap.Config
	switch app.Mode {
	case "prod":
		loggerConfig = zap.NewProductionConfig()
		loggerConfig.DisableStacktrace = true
	case "dev":
		loggerConfig = zap.NewDevelopmentConfig()
	default:
		loggerConfig = zap.NewDevelopmentConfig()
	}

	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)

	return nil
}

func Sync() {
	err := zap.S().Sync()
	if err != nil {
		log.Printf("Error while sync logger: %s", err.Error())
	}
}
