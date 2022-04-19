package gotwi_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/stretchr/testify/assert"
)

func Test_NewTypedClient(t *testing.T) {
	cases := []struct {
		name    string
		client  *gotwi.Client
		wantNil bool
	}{
		{
			name:   "ok",
			client: &gotwi.Client{},
		},
		{
			name:    "ok (nil)",
			client:  nil,
			wantNil: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			tc := gotwi.NewTypedClient[*gotwi.MockResponse](c.client)
			if c.wantNil {
				asst.Nil(tc)
				return
			}

			asst.NotNil(tc)
		})
	}
}

func Test_TypedClient_IsReady(t *testing.T) {
	cases := []struct {
		name   string
		client *gotwi.Client
		expect bool
	}{
		{
			name: "true: OAuth 1.0",
			client: gotwiClientField{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "oauth-token",
				SigningKey:           "signing-key",
			}.build(),
			expect: true,
		},
		{
			name: "true: OAuth 2.0",
			client: gotwiClientField{
				AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
				AccessToken:          "access-token",
			}.build(),
			expect: true,
		},
		{
			name:   "false: client is nil",
			client: nil,
			expect: false,
		},
		{
			name: "false: invalid authentication method",
			client: gotwiClientField{
				AuthenticationMethod: "invalid method",
				AccessToken:          "access-token",
			}.build(),
			expect: false,
		},
		{
			name: "false: OAuth 1.0: oauth token is empty",
			client: gotwiClientField{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "",
				SigningKey:           "signing-key",
			}.build(),
			expect: false,
		},
		{
			name: "false: OAuth 1.0: signing key is empty",
			client: gotwiClientField{
				AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
				OAuthToken:           "oauth-token",
				SigningKey:           "",
			}.build(),
			expect: false,
		},
		{
			name: "false: OAuth 2.0: access token is empty",
			client: gotwiClientField{
				AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
				AccessToken:          "",
			}.build(),
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			tc := gotwi.NewTypedClient[*gotwi.MockResponse](c.client)

			b := tc.IsReady()
			asst.Equal(c.expect, b)
		})
	}
}

func Test_TypedClient_Exec(t *testing.T) {
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
			asst := assert.New(tt)
			mockClient := newMockHTTPClient(c.mockInput)
			tc := gotwi.NewTypedClient[*gotwi.MockResponse](&gotwi.Client{
				Client: mockClient,
			})

			not200err, err := tc.Exec(c.req, &mockAPIResponse{})

			if c.wantErr {
				asst.Nil(not200err)
				asst.Error(err)
				return
			}

			if c.wantNot200Err {
				asst.Nil(err)
				asst.NotNil(not200err)
				return
			}

			asst.Nil(err)
			asst.Nil(not200err)
		})
	}
}

func Test_TypedClient_ExecStream(t *testing.T) {
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
			asst := assert.New(tt)
			mockClient := newMockHTTPClient(c.mockInput)
			tc := gotwi.NewTypedClient[*gotwi.MockResponse](&gotwi.Client{
				Client: mockClient,
			})

			res, not200err, err := tc.ExecStream(c.req)

			if c.wantErr {
				asst.Nil(res)
				asst.Nil(not200err)
				asst.Error(err)
				return
			}

			if c.wantNot200Err {
				asst.Nil(res)
				asst.Nil(err)
				asst.NotNil(not200err)
				return
			}

			asst.NotNil(res)
			asst.Nil(err)
			asst.Nil(not200err)
		})
	}
}

func Test_CallStreamAPI(t *testing.T) {
	cases := []struct {
		name            string
		mockInput       *mockInput
		clientInput     *gotwi.NewClientInput
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
			clientInput: &gotwi.NewClientInput{
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
			clientInput: &gotwi.NewClientInput{
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
			clientInput: &gotwi.NewClientInput{
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
			clientInput: &gotwi.NewClientInput{
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
			clientInput: &gotwi.NewClientInput{
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
			asst := assert.New(tt)
			tt.Setenv("GOTWI_API_KEY", c.envAPIKey)
			tt.Setenv("GOTWI_API_KEY_SECRET", c.envAPIKeySecret)

			mockClient := newMockHTTPClient(c.mockInput)
			in := c.clientInput
			in.HTTPClient = mockClient
			client, _ := gotwi.NewClient(in)
			tc := gotwi.NewTypedClient[*gotwi.MockResponse](client)

			s, err := tc.CallStreamAPI(context.Background(), c.endpoint, c.method, c.params)
			if c.wantErr {
				asst.Nil(s)
				asst.Error(err)
				return
			}

			asst.NotNil(s)
			asst.Nil(err)
		})
	}
}
