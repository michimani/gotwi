package gotwi_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/stretchr/testify/assert"
)

type testParameter struct {
	BodyResErr bool
}

func (tp testParameter) SetAccessToken(t string) {}

func (tp testParameter) AccessToken() string { return "token" }

func (tp testParameter) ResolveEndpoint(e string) string { return e }

func (tp testParameter) Body() (io.Reader, error) {
	if tp.BodyResErr {
		return nil, fmt.Errorf("body error")
	}
	return nil, nil
}

func (tp testParameter) ParameterMap() map[string]string { return nil }

func Test_NewGotwiClient(t *testing.T) {
	cases := []struct {
		name            string
		envAPIKey       string
		envAPIKeySecret string
		in              *gotwi.NewGotwiClientInput
		wantErr         bool
		expect          *gotwi.GotwiClient
	}{
		{
			name:            "normal: OAuth1.0",
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			in: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "oauth-token",
				OAuthTokenSecret:     "oauth-token-secret",
			},
			wantErr: false,
			expect: &gotwi.GotwiClient{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				AccessToken:          "",
				OAuthToken:           "oauth-token",
				OAuthConsumerKey:     "api-key",
				SigningKey:           "api-key-secret&oauth-token-secret",
			},
		},
		{
			name:            "error: input is nil",
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			in:              nil,
			wantErr:         true,
			expect:          nil,
		},
		{
			name:            "error: invalid authentication method",
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			in: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: "invalid method",
				OAuthToken:           "oauth-token",
				OAuthTokenSecret:     "oauth-token-secret",
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name:            "error: api key is empty",
			envAPIKey:       "",
			envAPIKeySecret: "api-key-secret",
			in: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "oauth-token",
				OAuthTokenSecret:     "oauth-token-secret",
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name:            "error: api key secret is empty",
			envAPIKey:       "api-key",
			envAPIKeySecret: "",
			in: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "oauth-token",
				OAuthTokenSecret:     "oauth-token-secret",
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name:            "error: OAuth1.0: oauth token is empty",
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			in: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "",
				OAuthTokenSecret:     "oauth-token-secret",
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name:            "error: OAuth1.0: oauth token secret is empty",
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			in: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "oauth-token",
				OAuthTokenSecret:     "",
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Setenv("GOTWI_API_KEY", c.envAPIKey)
			tt.Setenv("GOTWI_API_KEY_SECRET", c.envAPIKeySecret)

			gc, err := gotwi.NewGotwiClient(c.in)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, gc)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.expect.AuthenticationMethod, gc.AuthenticationMethod)
			assert.Equal(tt, c.expect.AccessToken, gc.AccessToken)
			assert.Equal(tt, c.expect.OAuthToken, gc.OAuthToken)
			assert.Equal(tt, c.expect.OAuthConsumerKey, gc.OAuthConsumerKey)
			assert.Equal(tt, c.expect.SigningKey, gc.SigningKey)
		})
	}
}

func Test_IsReady(t *testing.T) {
	cases := []struct {
		name   string
		client *gotwi.GotwiClient
		expect bool
	}{
		{
			name: "true: OAuth 1.0",
			client: &gotwi.GotwiClient{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "oauth-token",
				SigningKey:           "signing-key",
			},
			expect: true,
		},
		{
			name: "true: OAuth 2.0",
			client: &gotwi.GotwiClient{
				AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
				AccessToken:          "access-token",
			},
			expect: true,
		},
		{
			name:   "false: client is nil",
			client: nil,
			expect: false,
		},
		{
			name: "false: invalid authentication method",
			client: &gotwi.GotwiClient{
				AuthenticationMethod: "invalid method",
				AccessToken:          "access-token",
			},
			expect: false,
		},
		{
			name: "false: OAuth 1.0: oauth token is empty",
			client: &gotwi.GotwiClient{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "",
				SigningKey:           "signing-key",
			},
			expect: false,
		},
		{
			name: "false: OAuth 1.0: signing key is empty",
			client: &gotwi.GotwiClient{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "oauth-token",
				SigningKey:           "",
			},
			expect: false,
		},
		{
			name: "false: OAuth 2.0: access token is empty",
			client: &gotwi.GotwiClient{
				AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
				AccessToken:          "",
			},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			b := c.client.IsReady()
			assert.Equal(tt, c.expect, b)
		})
	}
}

func Test_newRequest(t *testing.T) {
	cases := []struct {
		name     string
		method   string
		endpoint string
		p        testParameter
		wantErr  bool
		expect   *http.Request
	}{
		{
			name:     "normal: GET",
			method:   "GET",
			endpoint: "endpoint",
			p:        testParameter{},
			wantErr:  false,
			expect: &http.Request{
				Method: "GET",
				URL:    &url.URL{Path: "endpoint"},
				Header: http.Header{
					"Content-Type": []string{"application/json;charset=UTF-8"},
				},
			},
		},
		{
			name:     "normal: POST",
			method:   "POST",
			endpoint: "endpoint",
			p:        testParameter{},
			wantErr:  false,
			expect: &http.Request{
				Method: "POST",
				URL:    &url.URL{Path: "endpoint"},
				Header: http.Header{
					"Content-Type": []string{"application/json;charset=UTF-8"},
				},
			},
		},
		{
			name:     "error: Body() returns error",
			method:   "GET",
			endpoint: "endpoint",
			p:        testParameter{BodyResErr: true},
			wantErr:  true,
			expect:   nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := gotwi.ExportNewRequest(context.Background(), c.endpoint, c.method, c.p)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, r)
				return
			}

			assert.Equal(tt, c.expect.Method, r.Method)
			assert.Equal(tt, c.expect.URL, r.URL)
			assert.Equal(tt, c.expect.Header, r.Header)
		})
	}
}
