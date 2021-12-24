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
