package proxy

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func newZapLogWriter(logger *zap.Logger) *log.Logger {
	return log.New(
		&zapLogWriter{logger},
		"",
		log.LstdFlags)
}

type zapLogWriter struct {
	logger *zap.Logger
}

func (w *zapLogWriter) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}
