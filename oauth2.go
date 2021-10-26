package gotwi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/internal/gotwierrors"
)

const OAuth2TokenEndpoint = "https://api.twitter.com/oauth2/token"

type OAuth2TokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func GenerateBearerToken(c *GotwiClient, apiKey, apiKeySecret string) (string, error) {
	uv := url.Values{}
	uv.Add("grant_type", "client_credentials")
	body := strings.NewReader(uv.Encode())

	req, err := http.NewRequest("POST", OAuth2TokenEndpoint, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.SetBasicAuth(apiKey, apiKeySecret)

	res, not200err, err := c.Exec(req)
	if err != nil {
		return "", err
	}

	if not200err != nil {
		return "", fmt.Errorf(gotwierrors.ErrorNon2XXStatus, not200err.Summary())
	}

	var o2r OAuth2TokenResponse
	if err = json.Unmarshal(res.Body, &o2r); err != nil {
		return "", err
	}

	if o2r.AccessToken == "" {
		return "", fmt.Errorf("access_token is empty")
	}

	return o2r.AccessToken, nil
}
