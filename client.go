package gotwi

import (
	"net/http"
	"os"
	"time"

	"github.com/michimani/gotwi/api/oauth2"
	"github.com/michimani/gotwi/types"
)

func NewClient() *types.TwitterClient {
	return &types.TwitterClient{
		Client: &http.Client{
			Timeout: time.Duration(30) * time.Second,
		},
	}
}

func NewAuthorizedClient() (*types.TwitterClient, error) {
	c := NewClient()
	return Authorize(c)
}

func Authorize(c *types.TwitterClient) (*types.TwitterClient, error) {
	apiKey := os.Getenv(types.APIKeyEnvName)
	apiKeySecret := os.Getenv(types.APIKeySecretEnvName)
	accessToken, err := oauth2.GenerateBearerToken(c, apiKey, apiKeySecret)
	if err != nil {
		return nil, err
	}

	c.AccessToken = accessToken
	return c, nil
}
