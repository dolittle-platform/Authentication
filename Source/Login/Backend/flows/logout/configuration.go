package logout

import "net/url"

type Configuration interface {
	FlowIDQueryParameter() string
	LoggedOutRedirect() *url.URL
}
