package public

import (
	"context"
	"net/http"
	"net/url"

	"dolittle.io/cookie-oidc-client/completion"
	"dolittle.io/cookie-oidc-client/cookies"
	"dolittle.io/cookie-oidc-client/server/handling"
	"dolittle.io/cookie-oidc-client/sessions"
)

type CompleteHandler handling.Handler

func NewCompleteHandler(parser completion.Parser, reader sessions.Reader, destroyer sessions.Destroyer, completer completion.Completer, writer cookies.Writer) CompleteHandler {
	return &complete{
		parser:    parser,
		reader:    reader,
		destroyer: destroyer,
		completer: completer,
		writer:    writer,
	}
}

type complete struct {
	parser    completion.Parser
	reader    sessions.Reader
	destroyer sessions.Destroyer
	completer completion.Completer
	writer    cookies.Writer
}

func (c *complete) Handle(w http.ResponseWriter, r *http.Request, ctx context.Context) error {
	response, err := c.parser.ParseFrom(r)
	if err != nil {
		return err
	}

	session, err := c.reader.ReadFrom(r)
	if err != nil {
		return err
	}

	err = c.destroyer.Destroy(r, w)
	if err != nil {
		// TODO: Log
	}

	token, err := c.completer.Complete(response, session)
	if err != nil {
		return err
	}

	if err := c.writer.WriteTokenCookie(token, w); err != nil {
		return err
	}

	redirect := url.URL(*session.ReturnTo)
	http.Redirect(w, r, redirect.String(), http.StatusFound)
	return nil
}
