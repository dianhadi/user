package log

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	logPath = "/var/log"
)

var (
	errorLogger zerolog.Logger
	infoLogger  zerolog.Logger
)

type Fields = map[string]interface{}

func New(serviceName string) {
	// Initialize zerolog global settings
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Initialize custom loggers
	errorLogger = initLogger(serviceName, "error.log", zerolog.ErrorLevel)
	infoLogger = initLogger(serviceName, "info.log", zerolog.InfoLevel)
}

func initLogger(serviceName, fileName string, level zerolog.Level) zerolog.Logger {
	path := fmt.Sprintf("%s/%s-%s", logPath, serviceName, fileName)

	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}

	var logger zerolog.Logger
	env := os.Getenv("ENV")
	if env == "" {
		logger = zerolog.New(io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr}, logFile)).Level(level)
	} else {
		logger = zerolog.New(logFile).Level(level)
	}

	logger = logger.With().Str("service", serviceName).Logger()

	return logger
}
