package audit

import (
	"net/http"

	"dolittle.io/impersonator/identities"
	"go.uber.org/zap"
)

type Logger interface {
	LogBefore(identity *identities.Identity, req *http.Request)
	LogAfter(identity *identities.Identity, req *http.Request, res *http.Response)
}

func NewLogger(zapLogger *zap.Logger) Logger {
	return &logger{
		logger: zapLogger,
	}
}

type logger struct {
	logger *zap.Logger
}

func (l *logger) LogBefore(identity *identities.Identity, req *http.Request) {
	l.logger.Info(
		"Audit request",
		zap.String("correlation", req.Header.Get("correlation")),
		zap.String("user-id", identity.UserID),
		zap.String("tenant-id", identity.TenantID),
		zap.String("method", req.Method),
		zap.String("path", req.URL.Path),
		zap.String("query", req.URL.RawQuery))
}

func (l *logger) LogAfter(identity *identities.Identity, req *http.Request, res *http.Response) {
	l.logger.Info(
		"Audit response",
		zap.String("correlation", req.Header.Get("correlation")),
		zap.String("user-id", identity.UserID),
		zap.String("tenant-id", identity.TenantID),
		zap.String("method", req.Method),
		zap.String("path", req.URL.Path),
		zap.String("query", req.URL.RawQuery),
		zap.Int("status", res.StatusCode))
}
