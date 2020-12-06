package consent

import "net/url"

type Accepter interface {
	AcceptConsentFlow(flow *Flow) (*url.URL, error)
}
