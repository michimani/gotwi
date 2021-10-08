package params_test

import (
	"testing"

	"github.com/michimani/gotwi/params"
	"github.com/stretchr/testify/assert"
)

func Test_UsersParams_SetAccessToken(t *testing.T) {
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
			p := &params.UsersParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UsersParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *params.UsersParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &params.UsersParams{IDs: []string{"test-id1", "test-id2"}},
			expect: endpointBase + "?ids=test-id1%2Ctest-id2",
		},
		{
			name: "normal: with expansions",
			params: &params.UsersParams{
				IDs:        []string{"test-id"},
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=test-id",
		},
		{
			name: "normal: with tweets.fields",
			params: &params.UsersParams{
				IDs:         []string{"test-id"},
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &params.UsersParams{
				IDs:        []string{"test-id"},
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=test-id&user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &params.UsersParams{
				IDs:         []string{"test-id"},
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointBase + "?expansions=ex&ids=test-id&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &params.UsersParams{
				IDs:         []string{},
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

func Test_UsersParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *params.UsersParams
	}{
		{
			name:   "empty params",
			params: &params.UsersParams{},
		},
		{
			name:   "some params",
			params: &params.UsersParams{IDs: []string{"id"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}
