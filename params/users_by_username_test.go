package params_test

import (
	"testing"

	"github.com/michimani/gotwi/params"
	"github.com/stretchr/testify/assert"
)

func Test_UsersByUsernameParams_SetAccessToken(t *testing.T) {
	cases := []struct {
		name   string
		token  string
		expect string
	}{
		{
			name:   "normal",
			token:  "test-token",
			expect: "test-token",
		},
		{
			name:   "normal: empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &params.UsersByUsernameParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UsersByUsernameParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:username"
	cases := []struct {
		name   string
		params *params.UsersByUsernameParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &params.UsersByUsernameParams{Username: "test-un"},
			expect: endpointRoot + "test-un",
		},
		{
			name: "normal: with expansions",
			params: &params.UsersByUsernameParams{
				Username:   "test-un",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-un?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &params.UsersByUsernameParams{
				Username:    "test-un",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-un?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &params.UsersByUsernameParams{
				Username:   "test-un",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-un?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &params.UsersByUsernameParams{
				Username:    "test-un",
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointRoot + "test-un?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &params.UsersByUsernameParams{
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_UsersByUsernameParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *params.UsersByUsernameParams
	}{
		{
			name:   "empty params",
			params: &params.UsersByUsernameParams{},
		},
		{
			name:   "some params",
			params: &params.UsersByUsernameParams{Username: "id"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}
