package gotwi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/types/response"
)

const OAuth2TokenEndpoint = "https://api.twitter.com/oauth2/token"

func GenerateBearerToken(c *TwitterClient, apiKey, apiKeySecret string) (string, error) {
	uv := url.Values{}
	uv.Add("grant_type", "client_credentials")
	body := strings.NewReader(uv.Encode())

	req, err := http.NewRequest("POST", OAuth2TokenEndpoint, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.SetBasicAuth(apiKey, apiKeySecret)

	res, err := c.Exec(req)
	if err != nil {
		return "", err
	}

	var o2r response.Token
	if err = json.Unmarshal(res.Body, &o2r); err != nil {
		return "", err
	}

	if o2r.AccessToken == "" {
		return "", fmt.Errorf("access_token is empty")
	}

	return o2r.AccessToken, nil
}
