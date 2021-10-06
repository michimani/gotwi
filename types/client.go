package types

import (
	"io/ioutil"
	"net/http"

	"github.com/michimani/gotwi/types/response"
)

const (
	APIKeyEnvName       = "GOTWI_API_KEY"
	APIKeySecretEnvName = "GOTWI_API_KEY_SECRET"
)

type TwitterClient struct {
	Client      *http.Client
	AccessToken string
}

func (c *TwitterClient) IsReady() bool {
	if c == nil {
		return false
	}

	if c.AccessToken == "" {
		return false
	}

	return true
}

func (c *TwitterClient) Exec(req *http.Request) (*response.ClientResponse, error) {
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &response.ClientResponse{
		StatusCode: res.StatusCode,
		Status:     res.Status,
		Body:       bytes,
	}, nil
}
