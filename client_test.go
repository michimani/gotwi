package gotwi_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/resources"
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
		mockInput       *mockInput
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
			name:            "normal: OAuth2.0",
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{"token_type":"token_type","access_token":"access_token"}`)),
			},
			in: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
			},
			wantErr: false,
			expect: &gotwi.GotwiClient{
				AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
				AccessToken:          "access_token",
				OAuthConsumerKey:     "api-key",
			},
		},
		{
			name:            "error: OAuth2.0",
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusInternalServerError,
				ResponseBody:       io.NopCloser(strings.NewReader(``)),
			},
			in: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
			},
			wantErr: true,
			expect:  nil,
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

			mockClient := newMockHTTPClient(c.mockInput)
			if mockClient != nil {
				c.in.HTTPClient = mockClient
			}

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

func Test_CallAPI(t *testing.T) {
	cases := []struct {
		name            string
		mockInput       *mockInput
		clientInput     *gotwi.NewGotwiClientInput
		endpoint        string
		method          string
		envAPIKey       string
		envAPIKeySecret string
		params          util.Parameters
		response        util.Response
		wantErr         bool
	}{
		{
			name: "ok: OAuth 1.0",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{"message": "ok"}`)),
			},
			clientInput: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "token",
				OAuthTokenSecret:     "secret",
			},
			endpoint:        "test-endpoint",
			method:          http.MethodGet,
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			params:          &mockAPIParameter{},
			response:        &mockAPIResponse{},
			wantErr:         false,
		},
		{
			name: "error: parameter is nil",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{"message": "ok"}`)),
			},
			clientInput: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "token",
				OAuthTokenSecret:     "secret",
			},
			endpoint:        "test-endpoint",
			method:          http.MethodGet,
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			params:          nil,
			response:        &mockAPIResponse{},
			wantErr:         true,
		},
		{
			name: "error: client is not ready",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{"message": "ok"}`)),
			},
			clientInput: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "",
				OAuthTokenSecret:     "secret",
			},
			endpoint:        "test-endpoint",
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			method:          http.MethodGet,
			params:          &mockAPIParameter{},
			response:        &mockAPIResponse{},
			wantErr:         true,
		},
		{
			name: "error: not 200 response",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusInternalServerError,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			clientInput: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "token",
				OAuthTokenSecret:     "secret",
			},
			endpoint:        "test-endpoint",
			method:          http.MethodGet,
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			params:          &mockAPIParameter{},
			response:        &mockAPIResponse{},
			wantErr:         true,
		},
		{
			name: "error: failed to decode json",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`///`)),
			},
			clientInput: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "token",
				OAuthTokenSecret:     "secret",
			},
			endpoint:        "test-endpoint",
			method:          http.MethodGet,
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			params:          &mockAPIParameter{},
			response:        &mockAPIResponse{},
			wantErr:         true,
		},
		{
			name: "error: invalid method",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			clientInput: &gotwi.NewGotwiClientInput{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "token",
				OAuthTokenSecret:     "secret",
			},
			endpoint:        "test-endpoint",
			method:          "invalid method",
			envAPIKey:       "api-key",
			envAPIKeySecret: "api-key-secret",
			params:          &mockAPIParameter{},
			response:        &mockAPIResponse{},
			wantErr:         true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			tt.Setenv("GOTWI_API_KEY", c.envAPIKey)
			tt.Setenv("GOTWI_API_KEY_SECRET", c.envAPIKeySecret)

			mockClient := newMockHTTPClient(c.mockInput)
			in := c.clientInput
			in.HTTPClient = mockClient
			client, _ := gotwi.NewGotwiClient(in)

			err := client.CallAPI(context.Background(), c.endpoint, c.method, c.params, c.response)
			if c.wantErr {
				assert.Error(tt, err)
				return
			}

			assert.Nil(tt, err)
		})
	}
}

func Test_Exec(t *testing.T) {
	nonErrReq, _ := http.NewRequestWithContext(context.TODO(), "GET", "https://example.com", nil)
	errReq := &http.Request{Method: "invalid method"}

	cases := []struct {
		name          string
		mockInput     *mockInput
		req           *http.Request
		wantErr       bool
		wantNot200Err bool
	}{
		{
			name: "ok",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			req:           nonErrReq,
			wantErr:       false,
			wantNot200Err: false,
		},
		{
			name: "error: not 200 error",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusInternalServerError,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			req:           nonErrReq,
			wantErr:       false,
			wantNot200Err: true,
		},
		{
			name: "error: cannot resolve 200 error",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusInternalServerError,
				ResponseHeader: map[string][]string{
					"Content-Type": {"application/json;charset=UTF-8"},
				},
				ResponseBody: io.NopCloser(strings.NewReader(`///`)),
			},
			req:           nonErrReq,
			wantErr:       true,
			wantNot200Err: false,
		},
		{
			name: "error: http.Client.Do error",
			mockInput: &mockInput{
				ResponseStatusCode: http.StatusInternalServerError,
				ResponseBody:       io.NopCloser(strings.NewReader(`{}`)),
			},
			req:           errReq,
			wantErr:       true,
			wantNot200Err: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			mockClient := newMockHTTPClient(c.mockInput)
			client := gotwi.GotwiClient{
				Client: mockClient,
			}

			not200err, err := client.Exec(c.req, &mockAPIResponse{})

			if c.wantErr {
				assert.Nil(tt, not200err)
				assert.Error(tt, err)
				return
			}

			if c.wantNot200Err {
				assert.Nil(tt, err)
				assert.NotNil(tt, not200err)
				return
			}

			assert.Nil(tt, err)
			assert.Nil(tt, not200err)
		})
	}
}

func Test_resolveNon2XXResponse(t *testing.T) {
	resetTime := time.Unix(int64(100000000), 0)

	cases := []struct {
		name             string
		res              *http.Response
		hasRateLimitInfo bool
		wantErr          bool
		expect           resources.Non2XXError
	}{
		{
			name: "normal: no rate limit error",
			res: &http.Response{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Header: map[string][]string{
					"Content-Type": {"application/json;charset=UTF-8"},
				},
				Body: io.NopCloser(strings.NewReader(`{"message": "error"}`)),
			},
			wantErr: false,
			expect: resources.Non2XXError{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "normal: content-type is text/plain",
			res: &http.Response{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Header: map[string][]string{
					"Content-Type": {"text/plain"},
				},
				Body: io.NopCloser(strings.NewReader("text error message")),
			},
			wantErr: false,
			expect: resources.Non2XXError{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				APIErrors: []resources.ErrorInformation{
					{Message: "text error message"},
				},
			},
		},
		{
			name: "normal: content-type is empty",
			res: &http.Response{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Header:     map[string][]string{},
			},
			wantErr: false,
			expect: resources.Non2XXError{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				APIErrors: []resources.ErrorInformation{
					{Message: "Content-Type is undefined."},
				},
			},
		},
		{
			name: "normal: rate limit error",
			res: &http.Response{
				Status:     "Too Many Requests",
				StatusCode: http.StatusTooManyRequests,
				Header: map[string][]string{
					"Content-Type":           {"application/json;charset=UTF-8"},
					"X-Rate-Limit-Limit":     {"1"},
					"X-Rate-Limit-Remaining": {"2"},
					"X-Rate-Limit-Reset":     {"100000000"},
				},
				Body: io.NopCloser(strings.NewReader(`{"message": "error"}`)),
			},
			hasRateLimitInfo: true,
			wantErr:          false,
			expect: resources.Non2XXError{
				Status:     "Too Many Requests",
				StatusCode: http.StatusTooManyRequests,
				RateLimitInfo: &util.RateLimitInformation{
					Limit:     1,
					Remaining: 2,
					ResetAt:   &resetTime,
				},
			},
		},
		{
			name: "error: failed to decode json",
			res: &http.Response{
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Header: map[string][]string{
					"Content-Type": {"application/json;charset=UTF-8"},
				},
				Body: io.NopCloser(strings.NewReader(`///`)),
			},
			wantErr: true,
		},
		{
			name: "error: on getting rate limit information",
			res: &http.Response{
				Status:     "Too Many Requests",
				StatusCode: http.StatusTooManyRequests,
				Header: map[string][]string{
					"Content-Type":           {"application/json;charset=UTF-8"},
					"X-Rate-Limit-Limit":     {"1"},
					"X-Rate-Limit-Remaining": {"xxxx"},
					"X-Rate-Limit-Reset":     {"100000000"},
				},
				Body: io.NopCloser(strings.NewReader(`{"message": "error"}`)),
			},
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			e, err := gotwi.ExportResolveNon2XXResponse(c.res)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, e)
				return
			}

			assert.NoError(tt, err)

			assert.Equal(tt, c.expect.APIErrors, e.APIErrors)
			assert.Equal(tt, c.expect.Title, e.Title)
			assert.Equal(tt, c.expect.Detail, e.Detail)
			assert.Equal(tt, c.expect.Type, e.Type)
			assert.Equal(tt, c.expect.Status, e.Status)
			assert.Equal(tt, c.expect.StatusCode, e.StatusCode)
			if c.hasRateLimitInfo {
				assert.NotNil(tt, e.RateLimitInfo)
			}
		})
	}
}
