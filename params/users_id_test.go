package params_test

import (
	"testing"

	"github.com/michimani/gotwi/params"
	"github.com/stretchr/testify/assert"
)

func Test_UsersIDParams_SetAccessToken(t *testing.T) {
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
			p := &params.UsersIDParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UsersIDParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *params.UsersIDParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &params.UsersIDParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: with expansions",
			params: &params.UsersIDParams{
				ID:         "test-id",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &params.UsersIDParams{
				ID:          "test-id",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &params.UsersIDParams{
				ID:         "test-id",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &params.UsersIDParams{
				ID:          "test-id",
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointRoot + "test-id?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &params.UsersIDParams{
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

func Test_UsersIDParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *params.UsersIDParams
	}{
		{
			name:   "empty params",
			params: &params.UsersIDParams{},
		},
		{
			name:   "some params",
			params: &params.UsersIDParams{ID: "id"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}
