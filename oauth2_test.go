package gotwi_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/resources"
	"github.com/stretchr/testify/assert"
)

func Test_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		o2r    gotwi.OAuth2TokenResponse
		expect bool
	}{
		{
			name:   "normal: initial struct",
			o2r:    gotwi.OAuth2TokenResponse{},
			expect: false,
		},
		{
			name: "normal: has values",
			o2r: gotwi.OAuth2TokenResponse{
				TokenType:   "token-type",
				AccessToken: "access-token",
			},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			b := c.o2r.HasPartialError()
			assert.Equal(tt, c.expect, b)
		})
	}
}

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

func Test_GenerateBearerToken(t *testing.T) {
	wantAccessToken := "access_token"

	cases := []struct {
		name    string
		client  *MockGotwiClientForOAuth2
		wantErr bool
		expect  string
	}{
		{
			name:    "normal",
			client:  newMockGotwiClientForOAuth2(wantAccessToken, false, false),
			wantErr: false,
			expect:  wantAccessToken,
		},
		{
			name:    "error: error",
			client:  newMockGotwiClientForOAuth2(wantAccessToken, true, false),
			wantErr: true,
			expect:  "",
		},
		{
			name:    "error: not 200 error",
			client:  newMockGotwiClientForOAuth2(wantAccessToken, false, true),
			wantErr: true,
			expect:  "",
		},
		{
			name:    "error: token is empty",
			client:  newMockGotwiClientForOAuth2("", false, false),
			wantErr: true,
			expect:  "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			a, err := gotwi.GenerateBearerToken(c.client, "key", "sec")
			if c.wantErr {
				assert.Error(tt, err)
				assert.Empty(tt, a)
				return
			}

			assert.Nil(tt, err)
			assert.Equal(tt, c.expect, a)
		})
	}
}
