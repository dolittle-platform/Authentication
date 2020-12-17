package proxy

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"dolittle.io/impersonator/audit"
	"dolittle.io/impersonator/identities"
	"dolittle.io/impersonator/proxy/context"
	"go.uber.org/zap"
)

var (
	impersonateHeaderRegex = regexp.MustCompile("(?i)^impersonate.*")
)

type Modifier interface {
	ModifyRequest(ctx *context.Context, req *http.Request) error
	ModifyResponse(ctx *context.Context, req *http.Request, res *http.Response) error
}

func NewModifier(configuration Configuration, audit audit.Logger, writer identities.Writer, logger *zap.Logger) Modifier {
	return &modifier{
		configuration: configuration,
		audit:         audit,
		writer:        writer,
		logger:        logger,
	}
}

type modifier struct {
	configuration Configuration
	audit         audit.Logger
	writer        identities.Writer
	logger        *zap.Logger
}

func (m *modifier) ModifyRequest(ctx *context.Context, req *http.Request) error {
	m.audit.LogBefore(ctx.Identity, req)

	m.removeImpersonateHeaders(req)
	m.writeRequestTarget(req)
	if err := m.writer.WriteIdentityTo(ctx.Identity, req); err != nil {
		return err
	}

	return nil
}

func (m *modifier) removeImpersonateHeaders(req *http.Request) {
	for header, values := range req.Header {
		if impersonateHeaderRegex.MatchString(header) {
			m.logger.Warn("Illegal impersonate header values removed from request", zap.String("header", header), zap.Strings("values", values))
			delete(req.Header, header)
		}
	}
}

func (m *modifier) writeRequestTarget(req *http.Request) {
	req.URL.Scheme = m.configuration.APIServerURL().Scheme
	req.URL.Host = m.configuration.APIServerURL().Host
}

func (m *modifier) ModifyResponse(ctx *context.Context, req *http.Request, res *http.Response) error {
	m.audit.LogAfter(ctx.Identity, req, res)

	requestHandlingTime.Observe(time.Now().Sub(ctx.RequestStarted).Seconds())
	tenantRequestTotal.WithLabelValues(req.Method, fmt.Sprint(res.StatusCode), ctx.Identity.TenantID).Inc()

	return nil
}
