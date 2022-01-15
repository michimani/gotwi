package gotwi_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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

type MockGotwiClientForOAuth2 struct {
	MockExec func(req *http.Request, i util.Response) (*resources.Non2XXError, error)
}

func (m MockGotwiClientForOAuth2) Exec(req *http.Request, i util.Response) (*resources.Non2XXError, error) {
	return m.MockExec(req, i)
}

func newMockGotwiClientForOAuth2(returnedToken string, execHasError, hasNot200Error bool) *MockGotwiClientForOAuth2 {
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

	return &MockGotwiClientForOAuth2{
		MockExec: fn,
	}
}
