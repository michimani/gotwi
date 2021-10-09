package gotwi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/michimani/gotwi/internal/gotwierrors"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/resources"
)

const (
	APIKeyEnvName       = "GOTWI_API_KEY"
	APIKeySecretEnvName = "GOTWI_API_KEY_SECRET"
)

type TwitterClient struct {
	Client      *http.Client
	AccessToken string
}

type ClientResponse struct {
	StatusCode int
	Status     string
	Error      *resources.Non200Error
	Body       []byte
}

func NewClient() *TwitterClient {
	return &TwitterClient{
		Client: &http.Client{
			Timeout: time.Duration(30) * time.Second,
		},
	}
}

func NewAuthorizedClient() (*TwitterClient, error) {
	c := NewClient()
	return Authorize(c)
}

func Authorize(c *TwitterClient) (*TwitterClient, error) {
	apiKey := os.Getenv(APIKeyEnvName)
	apiKeySecret := os.Getenv(APIKeySecretEnvName)
	accessToken, err := GenerateBearerToken(c, apiKey, apiKeySecret)
	if err != nil {
		return nil, err
	}

	c.AccessToken = accessToken
	return c, nil
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

func (c *TwitterClient) Exec(req *http.Request) (*ClientResponse, error) {
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		non200err := resources.Non200Error{}
		if err := json.Unmarshal(bytes, &non200err); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(gotwierrors.ErrorNon200Status, res.StatusCode, non200err.Title, non200err.Detail)
	}

	return &ClientResponse{
		StatusCode: res.StatusCode,
		Status:     res.Status,
		Body:       bytes,
	}, nil
}

func (c *TwitterClient) Prepare(endpointBase, method string, p util.Parameters) (*http.Request, error) {
	if p == nil {
		return nil, fmt.Errorf(gotwierrors.ErrorParametersNil, endpointBase)
	}

	if !c.IsReady() {
		return nil, fmt.Errorf(gotwierrors.ErrorClientNotReady)
	}

	endpoint := p.ResolveEndpoint(endpointBase)
	p.SetAccessToken(c.AccessToken)
	return newRequest(endpoint, method, p)
}

func newRequest(endpoint, method string, p util.Parameters) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, p.Body())
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.AccessToken()))

	return req, nil
}
