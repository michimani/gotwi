package gotwi_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/resources"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newMockClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

type mockInput struct {
	ResponseStatusCode int
	ResponseHeader     map[string][]string
	ResponseBody       io.ReadCloser
}

func newMockHTTPClient(in *mockInput) *http.Client {
	if in == nil {
		return nil
	}

	return newMockClient(func(req *http.Request) *http.Response {
		return &http.Response{
			Status:     "mock response status",
			StatusCode: in.ResponseStatusCode,
			Body:       in.ResponseBody,
			Header:     in.ResponseHeader,
		}
	})
}

type mockAPIParameter struct{}

func (mp mockAPIParameter) SetAccessToken(token string)                {}
func (mp mockAPIParameter) AccessToken() string                        { return "" }
func (mp mockAPIParameter) ResolveEndpoint(endpointBase string) string { return "" }
func (mp mockAPIParameter) Body() (io.Reader, error)                   { return nil, nil }
func (mp mockAPIParameter) ParameterMap() map[string]string            { return map[string]string{} }

type mockAPIResponse struct{}

func (mr mockAPIResponse) HasPartialError() bool { return false }

type MockGotwiClient struct {
	Client                   *http.Client
	MockExec                 func(req *http.Request, i util.Response) (*resources.Non2XXError, error)
	MockIsReady              func() bool
	MockAccessToken          func() string
	MockAuthenticationMethod func() gotwi.AuthenticationMethod
	MockOAuthToken           func() string
	MockOAuthConsumerKey     func() string
	MockSigningKey           func() string
}

func (m *MockGotwiClient) Exec(req *http.Request, i util.Response) (*resources.Non2XXError, error) {
	return m.MockExec(req, i)
}

func (m *MockGotwiClient) IsReady() bool {
	if m == nil {
		return false
	}

	if !m.AuthenticationMethod().Valid() {
		return false
	}

	switch m.AuthenticationMethod() {
	case gotwi.AuthenMethodOAuth1UserContext:
		if m.OAuthToken() == "" || m.SigningKey() == "" {
			return false
		}
	case gotwi.AuthenMethodOAuth2BearerToken:
		if m.AccessToken() == "" {
			return false
		}
	}

	return true
}

func (m *MockGotwiClient) AccessToken() string {
	return m.MockAccessToken()
}

func (m *MockGotwiClient) AuthenticationMethod() gotwi.AuthenticationMethod {
	return m.MockAuthenticationMethod()
}

func (m *MockGotwiClient) OAuthToken() string {
	return m.MockAccessToken()
}

func (m *MockGotwiClient) OAuthConsumerKey() string {
	return m.MockAccessToken()
}

func (m *MockGotwiClient) SigningKey() string {
	return m.MockAccessToken()
}

func newMockGotwiClient(returnedToken string, execHasError, hasNot200Error bool) *MockGotwiClient {
	fn := func(req *http.Request, i util.Response) (*resources.Non2XXError, error) {
		if execHasError {
			return nil, fmt.Errorf("has error")
		}

		if hasNot200Error {
			return &resources.Non2XXError{}, nil
		}

		resBody := strings.NewReader(`{"token_type":"token_type","access_token":"` + returnedToken + `"}`)

		if err := json.NewDecoder(resBody).Decode(i); err != nil {
			return nil, err
		}

		return nil, nil
	}

	return &MockGotwiClient{
		MockExec: fn,
	}
}

type mockFuncInput struct {
	MockExec                 func(req *http.Request, i util.Response) (*resources.Non2XXError, error)
	MockIsReady              func() bool
	MockAccessToken          func() string
	MockAuthenticationMethod func() gotwi.AuthenticationMethod
	MockOAuthToken           func() string
	MockOAuthConsumerKey     func() string
	MockSigningKey           func() string
}

func newMockGotwiClientWithFunc(in mockFuncInput) *MockGotwiClient {
	m := MockGotwiClient{}

	if in.MockExec != nil {
		m.MockExec = in.MockExec
	} else {
		m.MockExec = func(req *http.Request, i util.Response) (*resources.Non2XXError, error) { return nil, nil }
	}

	if in.MockIsReady != nil {
		m.MockIsReady = in.MockIsReady
	} else {
		m.MockIsReady = func() bool { return false }
	}

	if in.MockAccessToken != nil {
		m.MockAccessToken = in.MockAccessToken
	} else {
		m.MockAccessToken = func() string { return "" }
	}

	if in.MockAuthenticationMethod != nil {
		m.MockAuthenticationMethod = in.MockAuthenticationMethod
	} else {
		m.MockAuthenticationMethod = func() gotwi.AuthenticationMethod { return gotwi.AuthenticationMethod("") }
	}

	if in.MockOAuthToken != nil {
		m.MockOAuthToken = in.MockOAuthToken
	} else {
		m.MockOAuthToken = func() string { return "" }
	}

	if in.MockOAuthConsumerKey != nil {
		m.MockOAuthConsumerKey = in.MockOAuthConsumerKey
	} else {
		m.MockOAuthConsumerKey = func() string { return "" }
	}

	if in.MockSigningKey != nil {
		m.MockSigningKey = in.MockSigningKey
	} else {
		m.MockSigningKey = func() string { return "" }
	}

	return &m
}
