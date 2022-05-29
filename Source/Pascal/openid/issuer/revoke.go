package issuer

import (
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"strings"
)

func revokeToken(token, revokeURL, clientID, clientSecret string, authStyle oauth2.AuthStyle) error {
	req, err := newTokenRevokeRequest(token, revokeURL, clientID, clientSecret, authStyle)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return ErrTokenRevokeReturnedNotOK
	}

	return nil
}

func newTokenRevokeRequest(token, revokeURL, clientID, clientSecret string, authStyle oauth2.AuthStyle) (*http.Request, error) {
	v := url.Values{}
	v.Set("token", token)

	if authStyle == oauth2.AuthStyleInParams {
		if clientID != "" {
			v.Set("client_id", clientID)
		}
		if clientSecret != "" {
			v.Set("client_secret", clientSecret)
		}
	}

	req, err := http.NewRequest("POST", revokeURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if authStyle == oauth2.AuthStyleInHeader {
		req.SetBasicAuth(url.QueryEscape(clientID), url.QueryEscape(clientSecret))
	}

	return req, nil
}
