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
)

type AuthenticationMethod string

const (
	AuthenMethodOAuth1UserContext = "OAuth 1.0a User context"
	AuthenMethodOAuth2BearerToken = "OAuth 2.0 Bearer token"
)

func (a AuthenticationMethod) Valid() bool {
	return a == AuthenMethodOAuth1UserContext || a == AuthenMethodOAuth2BearerToken
}

type NewGotwiClientInput struct {
	HTTPClient           *http.Client
	AuthenticationMethod AuthenticationMethod
	OAuthToken           string
	OAuthTokenSecret     string
}

type GotwiClient struct {
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

var defaultHTTPClient = &http.Client{
	Timeout: time.Duration(30) * time.Second,
}

func NewGotwiClient(in *NewGotwiClientInput) (*GotwiClient, error) {
	if in == nil {
		return nil, fmt.Errorf("NewGotwiClientInput is nil.")
	}

	if !in.AuthenticationMethod.Valid() {
		return nil, fmt.Errorf("AuthenticationMethod is invalid.")
	}

	c := GotwiClient{
		Client:               defaultHTTPClient,
		AuthenticationMethod: in.AuthenticationMethod,
	}

	if in.HTTPClient != nil {
		c.Client = in.HTTPClient
	}

	if err := c.authorize(in.OAuthToken, in.OAuthTokenSecret); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *GotwiClient) authorize(oauthToken, oauthTokenSecret string) error {
	apiKey := os.Getenv(APIKeyEnvName)
	apiKeySecret := os.Getenv(APIKeySecretEnvName)
	if apiKey == "" || apiKeySecret == "" {
		return fmt.Errorf("env '%s' and '%s' is required.", APIKeyEnvName, APIKeySecretEnvName)
	}
	c.OAuthConsumerKey = apiKey

	switch c.AuthenticationMethod {
	case AuthenMethodOAuth1UserContext:
		if oauthToken == "" || oauthTokenSecret == "" {
			return fmt.Errorf("OAuthToken and OAuthTokenSecret is required for using %s.", AuthenMethodOAuth1UserContext)
		}

		c.OAuthToken = oauthToken
		c.SigningKey = fmt.Sprintf("%s&%s",
			url.QueryEscape(apiKeySecret),
			url.QueryEscape(oauthTokenSecret))
	case AuthenMethodOAuth2BearerToken:
		accessToken, err := GenerateBearerToken(c, apiKey, apiKeySecret)
		if err != nil {
			return err
		}

		c.AccessToken = accessToken
	default:
		// noop
	}

	return nil
}

func (c *GotwiClient) IsReady() bool {
	if c == nil {
		return false
	}

	if !c.AuthenticationMethod.Valid() {
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

func (c *GotwiClient) CallAPI(endpoint, method string, p util.Parameters, i util.Response) error {
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

func (c *GotwiClient) Exec(req *http.Request) (*ClientResponse, *resources.Non200Error, error) {
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

func (c *GotwiClient) prepare(endpointBase, method string, p util.Parameters) (*http.Request, error) {
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
		req, err = c.setOAuth1Header(req, pm)
		if err != nil {
			return nil, err
		}
	case AuthenMethodOAuth2BearerToken:
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.AccessToken()))
	default:
		// noop
	}

	return req, nil
}

const oauth1header = `OAuth oauth_consumer_key="%s",oauth_nonce="%s",oauth_signature="%s",oauth_signature_method="%s",oauth_timestamp="%s",oauth_token="%s",oauth_version="%s"`

// setOAuth1Header returns http.Request with the header information required for OAuth1.0a authentication.
func (c *GotwiClient) setOAuth1Header(r *http.Request, paramsMap map[string]string) (*http.Request, error) {
	in := &CreateOAthSignatureInput{
		HTTPMethod:       r.Method,
		RawEndpoint:      r.URL.String(),
		OAuthConsumerKey: c.OAuthConsumerKey,
		OAuthToken:       c.OAuthToken,
		SigningKey:       c.SigningKey,
		ParameterMap:     paramsMap,
	}

	out, err := CreateOAuthSignature(in)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", fmt.Sprintf(oauth1header,
		url.QueryEscape(c.OAuthConsumerKey),
		url.QueryEscape(out.OAuthNonce),
		url.QueryEscape(out.OAuthSignature),
		url.QueryEscape(out.OAuthSignatureMethod),
		url.QueryEscape(out.OAuthTimestamp),
		url.QueryEscape(c.OAuthToken),
		url.QueryEscape(out.OAuthVersion),
	))

	return r, nil
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
