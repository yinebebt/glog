package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
)

// initLogger setup and return zap logger with set  of configuration
func initLogger(isDebug bool) *zap.Logger {
	logDir := os.Getenv("LOG_DIR") // e.g. /var/log/zaplog
	if logDir == "" {
		logDir = "./"
	}
	err := os.MkdirAll(logDir, 0755) // read/write/execute permissions
	if err != nil {
		log.Fatalf("failed to create log directory: %v", err)
	}

	logFilePath := filepath.Join(logDir, "zap.log")
	file, err := os.Create(logFilePath)
	if err != nil {
		log.Printf("failed to create log file: %v", err)
	}
	err = file.Close()
	if err != nil {
		log.Printf("failed to close log file: %v", err)
	}

	cfg := zap.NewProductionConfig()
	if isDebug {
		cfg = zap.NewDevelopmentConfig()
	}
	// time encoding format, a more readable, similar to ISO8601 format
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	// stacktrace level default to warn in development and error in production
	// AddCallerSkip(1) removes the call from the zaplog package itself.
	cfg.OutputPaths = []string{"stderr", logFilePath}
	logger, err := cfg.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	return logger
}

// PlainError remove some verbosity and make the error more concise and useful. The default error formatter is a bit complicated.
// Check this issue on GitHub: https://github.com/uber-go/zap/issues/650
type plainError struct {
	e error
}

func (pe plainError) Error() string {
	return pe.e.Error()
}

func PlainError(err error) zap.Field {
	return zap.Error(plainError{err})
}
