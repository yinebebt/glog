package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func main() {
	log := initLogger(false)
	log.Info("info message", zap.String("user", "alice"))
	log.Warn("warn message", zap.String("city", "Addis Ababa"))
	log.Error("error message", zap.String("name", "Alice B."))

	// create an error with stack
	err := errors.WithStack(sql.ErrNoRows)
	log.Error("error message", zap.String("id", "12345"), zap.Error(err))

	// use PlainError
	log.Error("plain error", PlainError(err))

	log.Info("all done")
}
