package gotwi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

type IGotwiClient interface {
	Exec(req *http.Request, i util.Response) (*resources.Non2XXError, error)
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
	Error      *resources.Non2XXError
	Body       []byte
	Response   util.Response
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
	}

	return true
}

func (c *GotwiClient) CallAPI(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
	req, err := c.prepare(ctx, endpoint, method, p)
	if err != nil {
		return wrapErr(err)
	}

	non200err, err := c.Exec(req, i)
	if err != nil {
		return wrapErr(err)
	}

	if non200err != nil {
		return wrapWithAPIErr(non200err)
	}

	return nil
}

var okCodes map[int]struct{} = map[int]struct{}{
	http.StatusOK:      {},
	http.StatusCreated: {},
}

func (c *GotwiClient) Exec(req *http.Request, i util.Response) (*resources.Non2XXError, error) {
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if _, ok := okCodes[res.StatusCode]; !ok {
		non200err, err := resolveNon2XXResponse(res)
		if err != nil {
			return nil, err
		}
		return non200err, nil
	}

	if err := json.NewDecoder(res.Body).Decode(i); err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *GotwiClient) prepare(ctx context.Context, endpointBase, method string, p util.Parameters) (*http.Request, error) {
	if p == nil {
		return nil, fmt.Errorf(gotwierrors.ErrorParametersNil, endpointBase)
	}

	if !c.IsReady() {
		return nil, fmt.Errorf(gotwierrors.ErrorClientNotReady)
	}

	endpoint := p.ResolveEndpoint(endpointBase)
	p.SetAccessToken(c.AccessToken)
	req, err := newRequest(ctx, endpoint, method, p)
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
	}

	return req, nil
}

const oauth1header = `OAuth oauth_consumer_key="%s",oauth_nonce="%s",oauth_signature="%s",oauth_signature_method="%s",oauth_timestamp="%s",oauth_token="%s",oauth_version="%s"`

// setOAuth1Header returns http.Request with the header information required for OAuth1.0a authentication.
func (c *GotwiClient) setOAuth1Header(r *http.Request, paramsMap map[string]string) (*http.Request, error) {
	in := &CreateOAuthSignatureInput{
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

func newRequest(ctx context.Context, endpoint, method string, p util.Parameters) (*http.Request, error) {
	body, err := p.Body()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	return req, nil
}

func resolveNon2XXResponse(res *http.Response) (*resources.Non2XXError, error) {
	non200err := &resources.Non2XXError{
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}

	cts := util.HeaderValues("Content-Type", res.Header)
	if len(cts) == 0 {
		non200err.APIErrors = []resources.ErrorInformation{
			{Message: "Content-Type is undefined."},
		}
		return non200err, nil
	}

	if !strings.Contains(cts[0], "application/json") {
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		non200err.APIErrors = []resources.ErrorInformation{
			{Message: strings.TrimRight(string(bytes), "\n")},
		}
	} else {
		if err := json.NewDecoder(res.Body).Decode(non200err); err != nil {
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

	return non200err, nil
}
