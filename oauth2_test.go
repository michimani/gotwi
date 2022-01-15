package gotwi_test

import (
	"testing"

	"github.com/michimani/gotwi"
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
