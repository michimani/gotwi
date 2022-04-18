package gotwi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/resources"
)

type TypedClient[T util.Response] struct {
	Client               *http.Client
	AuthenticationMethod AuthenticationMethod
	AccessToken          string
	OAuthToken           string
	SigningKey           string
	OAuthConsumerKey     string
}

func NewTypedClient[T util.Response](c *Client) *TypedClient[T] {
	return &TypedClient[T]{
		Client:               c.Client,
		AuthenticationMethod: c.AuthenticationMethod,
		AccessToken:          c.AccessToken,
		OAuthToken:           c.OAuthToken,
		SigningKey:           c.SigningKey,
		OAuthConsumerKey:     c.OAuthConsumerKey,
	}
}

func (c *TypedClient[T]) IsReady() bool {
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

func (c *TypedClient[T]) Exec(req *http.Request, i util.Response) (*resources.Non2XXError, error) {
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

func (c *TypedClient[T]) accessToken() string {
	return c.AccessToken
}

func (c *TypedClient[T]) authenticationMethod() AuthenticationMethod {
	return c.AuthenticationMethod
}

// setOAuth1Header returns http.Request with the header information required for OAuth1.0a authentication.
func (c *TypedClient[T]) setOAuth1Header(r *http.Request, paramsMap map[string]string) (*http.Request, error) {
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

func (c *TypedClient[T]) CallStreamAPI(ctx context.Context, endpoint, method string, p util.Parameters) (*StreamClient[T], error) {
	req, err := c.prepare(ctx, endpoint, method, p)
	if err != nil {
		return nil, wrapErr(err)
	}

	res, non200err, err := c.ExecStream(req)
	if err != nil {
		return nil, wrapErr(err)
	}

	if non200err != nil {
		return nil, wrapWithAPIErr(non200err)
	}

	s, err := newStreamClient[T](res)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (c *TypedClient[T]) ExecStream(req *http.Request) (*http.Response, *resources.Non2XXError, error) {
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if _, ok := okCodes[res.StatusCode]; !ok {
		non200err, err := resolveNon2XXResponse(res)
		if err != nil {
			return nil, nil, err
		}
		return nil, non200err, nil
	}

	return res, nil, nil
}

func (c *TypedClient[T]) prepare(ctx context.Context, endpointBase, method string, p util.Parameters) (*http.Request, error) {
	return prepare(ctx, endpointBase, method, p, c)
}
