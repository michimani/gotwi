package params_test

import (
	"testing"

	"github.com/michimani/gotwi/params"
	"github.com/stretchr/testify/assert"
)

func Test_UsersByParams_SetAccessToken(t *testing.T) {
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
			p := &params.UsersByParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UsersByParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *params.UsersByParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &params.UsersByParams{Usernames: []string{"test-un1", "test-un2"}},
			expect: endpointBase + "?usernames=test-un1%2Ctest-un2",
		},
		{
			name: "normal: with expansions",
			params: &params.UsersByParams{
				Usernames:  []string{"test-un"},
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&usernames=test-un",
		},
		{
			name: "normal: with tweets.fields",
			params: &params.UsersByParams{
				Usernames:   []string{"test-un"},
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?tweet.fields=tf1%2Ctf2&usernames=test-un",
		},
		{
			name: "normal: with users.fields",
			params: &params.UsersByParams{
				Usernames:  []string{"test-un"},
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointBase + "?user.fields=uf1%2Cuf2&usernames=test-un",
		},
		{
			name: "normal: all query parameters",
			params: &params.UsersByParams{
				Usernames:   []string{"test-un"},
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointBase + "?expansions=ex&tweet.fields=tf&user.fields=uf&usernames=test-un",
		},
		{
			name: "normal: has no required parameter",
			params: &params.UsersByParams{
				Usernames:   []string{},
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

func Test_UsersByParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *params.UsersByParams
	}{
		{
			name:   "empty params",
			params: &params.UsersByParams{},
		},
		{
			name:   "some params",
			params: &params.UsersByParams{Usernames: []string{"id"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}
