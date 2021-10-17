package gotwi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/michimani/gotwi/internal/gotwierrors"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/resources"
)

const (
	APIKeyEnvName       = "GOTWI_API_KEY"
	APIKeySecretEnvName = "GOTWI_API_KEY_SECRET"
	OAuthToken          = "GOTWI_ACCESS_TOKEN"
	OAuthTokenSecret    = "GOTWI_ACCESS_TOKEN_SECRET"
)

type AuthenticationMethod string

const (
	AuthenMethodOAuth1UserContext = "OAuth 1.0a User context"
	AuthenMethodOAuth2BearerToken = "OAuth 2.0 Bearer token"
)

func (a AuthenticationMethod) Valid() bool {
	return a == AuthenMethodOAuth1UserContext || a == AuthenMethodOAuth2BearerToken
}

type TwitterClient struct {
	Client               *http.Client
	AuthenticationMethod AuthenticationMethod
	AccessToken          string
	OAuthToken           string
	SigningKey           string
	OAuthConsumerKey     string
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

func NewAuthorizedClient(authnMethod AuthenticationMethod) (*TwitterClient, error) {
	c := NewClient()
	c.AuthenticationMethod = authnMethod
	return Authorize(c, authnMethod)
}

func Authorize(c *TwitterClient, authnMethod AuthenticationMethod) (*TwitterClient, error) {
	if !authnMethod.Valid() {
		return nil, fmt.Errorf("invalid authnMethod")
	}

	apiKey := os.Getenv(APIKeyEnvName)
	apiKeySecret := os.Getenv(APIKeySecretEnvName)
	c.OAuthConsumerKey = apiKey

	switch authnMethod {
	case AuthenMethodOAuth1UserContext:
		oauthToken := os.Getenv(OAuthToken)
		oauthTokenSecret := os.Getenv(OAuthTokenSecret)
		if oauthToken != "" && oauthTokenSecret != "" {
			c.OAuthToken = oauthToken
			c.SigningKey = fmt.Sprintf("%s&%s",
				url.QueryEscape(apiKeySecret),
				url.QueryEscape(oauthTokenSecret))
		}
	case AuthenMethodOAuth2BearerToken:
		accessToken, err := GenerateBearerToken(c, apiKey, apiKeySecret)
		if err != nil {
			return nil, err
		}

		c.AccessToken = accessToken
	default:
		// noop
	}

	return c, nil
}

func (c *TwitterClient) IsReady() bool {
	if c == nil {
		return false
	}
	switch c.AuthenticationMethod {
	case AuthenMethodOAuth1UserContext:
		if c.OAuthToken == "" || c.SigningKey == "" {
			return false
		}
	case AuthenMethodOAuth2BearerToken:
		if c.AccessToken == "" {
			return false
		}
	default:
		// noop
	}

	return true
}

func (c *TwitterClient) CallAPI(endpoint, method string, p util.Parameters, i util.Response) error {
	req, err := c.prepare(endpoint, method, p)
	if err != nil {
		return err
	}

	res, not200err, err := c.Exec(req)
	if err != nil {
		return err
	}

	if not200err != nil {
		return fmt.Errorf(gotwierrors.ErrorNon200Status, not200err.Summary())
	}

	if err := json.Unmarshal(res.Body, &i); err != nil {
		return err
	}

	return nil
}

func (c *TwitterClient) Exec(req *http.Request) (*ClientResponse, *resources.Non200Error, error) {
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != http.StatusOK {
		non200err, err := resolveNon200Response(res, bytes)
		if err != nil {
			return nil, nil, err
		}
		return nil, non200err, nil
	}

	return &ClientResponse{
		StatusCode: res.StatusCode,
		Status:     res.Status,
		Body:       bytes,
	}, nil, nil
}

func (c *TwitterClient) prepare(endpointBase, method string, p util.Parameters) (*http.Request, error) {
	if p == nil {
		return nil, fmt.Errorf(gotwierrors.ErrorParametersNil, endpointBase)
	}

	if !c.IsReady() {
		return nil, fmt.Errorf(gotwierrors.ErrorClientNotReady)
	}

	endpoint := p.ResolveEndpoint(endpointBase)
	p.SetAccessToken(c.AccessToken)
	req, err := newRequest(endpoint, method, p)
	if err != nil {
		return nil, err
	}

	switch c.AuthenticationMethod {
	case AuthenMethodOAuth1UserContext:
		pm := p.ParameterMap()
		req, err = c.SetOAuth1Header(req, pm)
		if err != nil {
			return nil, err
		}

		fmt.Printf("%#+v\n", req)
	case AuthenMethodOAuth2BearerToken:
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.AccessToken()))
	default:
		// noop
	}

	return req, nil
}

func newRequest(endpoint, method string, p util.Parameters) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, p.Body())
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	return req, nil
}

func resolveNon200Response(res *http.Response, bodyBytes []byte) (*resources.Non200Error, error) {
	non200err := resources.Non200Error{
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}

	cts := util.HeaderValues("Content-Type", res.Header)
	if len(cts) == 0 {
		non200err.Errors = []resources.ErrorInformation{
			{Message: "Content-Type is undefined."},
		}
		return &non200err, nil
	}

	if !strings.Contains(cts[0], "application/json") {
		non200err.Errors = []resources.ErrorInformation{
			{Message: strings.TrimRight(string(bodyBytes), "\n")},
		}
	} else {
		if err := json.Unmarshal(bodyBytes, &non200err); err != nil {
			return nil, err
		}
	}

	// additional information for Rate Limit
	if res.StatusCode == http.StatusTooManyRequests {
		rri, err := util.GetRateLimitInformation(res)
		if err != nil {
			return nil, err
		}

		non200err.RateLimitInfo = rri
	}

	return &non200err, nil
}
