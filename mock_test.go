package gotwi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/resources"
	"github.com/stretchr/testify/assert"
)

func Test_RoundTripFunc(t *testing.T) {
	cases := []struct {
		name           string
		roundTripFunc  RoundTripFunc
		expectedStatus int
	}{
		{
			name: "success: returns status code 200",
			roundTripFunc: func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader("")),
				}
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			client := newMockClient(c.roundTripFunc)
			req, _ := http.NewRequest("GET", "http://example.com", nil)
			resp, err := client.Do(req)

			asst.NoError(err)
			asst.Equal(c.expectedStatus, resp.StatusCode)
		})
	}
}

func Test_NewMockHTTPClient(t *testing.T) {
	cases := []struct {
		name           string
		input          *MockInput
		expectedStatus int
	}{
		{
			name: "success: returns status code 200",
			input: &MockInput{
				ResponseStatusCode: http.StatusOK,
				ResponseBody:       io.NopCloser(strings.NewReader("")),
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "error: input is nil",
			input:          nil,
			expectedStatus: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			client := NewMockHTTPClient(c.input)
			if c.input == nil {
				asst.Nil(client)
				return
			}

			req, _ := http.NewRequest("GET", "http://example.com", nil)
			resp, err := client.Do(req)

			asst.NoError(err)
			asst.Equal(c.expectedStatus, resp.StatusCode)
		})
	}
}

func Test_MockGotwiClient_Exec(t *testing.T) {
	cases := []struct {
		name           string
		returnedToken  string
		execHasError   bool
		hasNot200Error bool
		expectedError  bool
	}{
		{
			name:           "success: returns token",
			returnedToken:  "test-token",
			execHasError:   false,
			hasNot200Error: false,
			expectedError:  false,
		},
		{
			name:           "error: returns error",
			returnedToken:  "",
			execHasError:   true,
			hasNot200Error: false,
			expectedError:  true,
		},
		{
			name:           "error: returns non-200 error",
			returnedToken:  "",
			execHasError:   false,
			hasNot200Error: true,
			expectedError:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			client := NewMockGotwiClient(c.returnedToken, c.execHasError, c.hasNot200Error)
			req, _ := http.NewRequest("GET", "http://example.com", nil)
			var response util.Response = &MockAPIResponse{}

			non2xxError, err := client.Exec(req, response)

			if c.expectedError {
				asst.Error(err)
			} else {
				asst.NoError(err)
			}

			if c.hasNot200Error {
				asst.NotNil(non2xxError)
			} else {
				asst.Nil(non2xxError)
			}
		})
	}
}

func Test_MockGotwiClient_IsReady(t *testing.T) {
	cases := []struct {
		name   string
		client *MockGotwiClient
		expect bool
	}{
		{
			name: "success: returns true with AuthenMethodOAuth1UserContext",
			client: NewMockGotwiClientWithFunc(MockFuncInput{
				MockAuthenticationMethod: func() AuthenticationMethod {
					return AuthenMethodOAuth1UserContext
				},
				MockAccessToken: func() string {
					return "test-token"
				},
				MockSigningKey: func() string {
					return "test-signing-key"
				},
			}),
			expect: true,
		},
		{
			name: "error: returns false with AuthenMethodOAuth1UserContext",
			client: NewMockGotwiClientWithFunc(MockFuncInput{
				MockAuthenticationMethod: func() AuthenticationMethod {
					return AuthenMethodOAuth1UserContext
				},
				MockSigningKey: func() string {
					return "test-signing-key"
				},
			}),
			expect: false,
		},
		{
			name: "success: returns true with AuthenMethodOAuth2BearerToken",
			client: NewMockGotwiClientWithFunc(MockFuncInput{
				MockAuthenticationMethod: func() AuthenticationMethod {
					return AuthenMethodOAuth2BearerToken
				},
				MockAccessToken: func() string {
					return "test-token"
				},
			}),
			expect: true,
		},
		{
			name: "error: returns false with AuthenMethodOAuth2BearerToken",
			client: NewMockGotwiClientWithFunc(MockFuncInput{
				MockAuthenticationMethod: func() AuthenticationMethod {
					return AuthenMethodOAuth2BearerToken
				},
			}),
			expect: false,
		},
		{
			name: "error: return false with invalid AuthenticationMethod",
			client: NewMockGotwiClientWithFunc(MockFuncInput{
				MockAuthenticationMethod: func() AuthenticationMethod {
					return AuthenticationMethod("invalid method")
				},
			}),
			expect: false,
		},
		{
			name:   "error: returns false with nil client",
			client: nil,
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			asst.Equal(c.expect, c.client.IsReady())
		})
	}
}

func Test_MockGotwiClientWithFunc(t *testing.T) {
	cases := []struct {
		name  string
		input MockFuncInput
	}{
		{
			name:  "success: uses default functions",
			input: MockFuncInput{},
		},
		{
			name: "success: uses custom functions",
			input: MockFuncInput{
				MockExec: func(req *http.Request, i util.Response) (*resources.Non2XXError, error) {
					return nil, nil
				},
				MockIsReady:     func() bool { return true },
				MockAccessToken: func() string { return "custom-token" },
				MockAuthenticationMethod: func() AuthenticationMethod {
					return AuthenMethodOAuth2BearerToken
				},
				MockOAuthToken: func() string {
					return "custom-oauth-token"
				},
				MockOAuthConsumerKey: func() string {
					return "custom-oauth-consumer-key"
				},
				MockSigningKey: func() string {
					return "custom-signing-key"
				},
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			client := NewMockGotwiClientWithFunc(c.input)

			// can call all methods
			asst.NotPanics(func() {
				client.Exec(nil, nil)
				client.IsReady()
				client.AccessToken()
				client.AuthenticationMethod()
				client.OAuthToken()
				client.OAuthConsumerKey()
				client.SigningKey()
				client.CallAPI(context.Background(), "endpoint", "method", nil, nil)
			})
		})
	}
}

func Test_MockAPIParameter(t *testing.T) {
	param := MockAPIParameter{}

	t.Run("test MockAPIParameter methods", func(tt *testing.T) {
		asst := assert.New(tt)
		param.SetAccessToken("test-token")
		asst.Equal("", param.AccessToken())
		asst.Equal("", param.ResolveEndpoint("base"))

		body, err := param.Body()
		asst.Nil(body)
		asst.NoError(err)

		params := param.ParameterMap()
		asst.Empty(params)
	})
}

func Test_MockAPIResponse(t *testing.T) {
	response := MockAPIResponse{}

	t.Run("test MockAPIResponse methods", func(tt *testing.T) {
		asst := assert.New(tt)
		asst.False(response.HasPartialError())
	})
}

func Test_MockGotwiClient_CallAPI(t *testing.T) {
	cases := []struct {
		name        string
		client      *MockGotwiClient
		expectError bool
	}{
		{
			name: "success: returns true",
			client: &MockGotwiClient{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return nil
				},
			},
			expectError: false,
		},
		{
			name: "error: returns error",
			client: &MockGotwiClient{
				MockCallAPI: func(ctx context.Context, endpoint, method string, p util.Parameters, i util.Response) error {
					return errors.New("error")
				},
			},
			expectError: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			res := c.client.CallAPI(context.Background(), "endpoint", "method", nil, nil)
			if c.expectError {
				asst.Error(res)
			} else {
				asst.NoError(res)
			}
		})
	}
}
