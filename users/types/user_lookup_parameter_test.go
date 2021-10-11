package types_test

import (
	"testing"

	"github.com/michimani/gotwi/users/types"
	"github.com/stretchr/testify/assert"
)

func Test_UserLookupUsersParams_SetAccessToken(t *testing.T) {
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
			p := &types.UserLookupUsersParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UserLookupUsersParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.UserLookupUsersParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.UserLookupUsersParams{IDs: []string{"test-id1", "test-id2"}},
			expect: endpointBase + "?ids=test-id1%2Ctest-id2",
		},
		{
			name: "normal: with expansions",
			params: &types.UserLookupUsersParams{
				IDs:        []string{"test-id"},
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=test-id",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.UserLookupUsersParams{
				IDs:         []string{"test-id"},
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.UserLookupUsersParams{
				IDs:        []string{"test-id"},
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=test-id&user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.UserLookupUsersParams{
				IDs:         []string{"test-id"},
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointBase + "?expansions=ex&ids=test-id&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.UserLookupUsersParams{
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

func Test_UserLookupUsersParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UserLookupUsersParams
	}{
		{
			name:   "empty params",
			params: &types.UserLookupUsersParams{},
		},
		{
			name:   "some params",
			params: &types.UserLookupUsersParams{IDs: []string{"id"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}

func Test_UserLookupUsersIDParams_SetAccessToken(t *testing.T) {
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
			p := &types.UserLookupUsersIDParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UserLookupUsersIDParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.UserLookupUsersIDParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.UserLookupUsersIDParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: with expansions",
			params: &types.UserLookupUsersIDParams{
				ID:         "test-id",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.UserLookupUsersIDParams{
				ID:          "test-id",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.UserLookupUsersIDParams{
				ID:         "test-id",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.UserLookupUsersIDParams{
				ID:          "test-id",
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointRoot + "test-id?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.UserLookupUsersIDParams{
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

func Test_UserLookupUsersIDParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UserLookupUsersIDParams
	}{
		{
			name:   "empty params",
			params: &types.UserLookupUsersIDParams{},
		},
		{
			name:   "some params",
			params: &types.UserLookupUsersIDParams{ID: "id"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}

func Test_UserLookupUsersByParams_SetAccessToken(t *testing.T) {
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
			p := &types.UserLookupUsersByParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UserLookupUsersByParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.UserLookupUsersByParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.UserLookupUsersByParams{Usernames: []string{"test-un1", "test-un2"}},
			expect: endpointBase + "?usernames=test-un1%2Ctest-un2",
		},
		{
			name: "normal: with expansions",
			params: &types.UserLookupUsersByParams{
				Usernames:  []string{"test-un"},
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&usernames=test-un",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.UserLookupUsersByParams{
				Usernames:   []string{"test-un"},
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?tweet.fields=tf1%2Ctf2&usernames=test-un",
		},
		{
			name: "normal: with users.fields",
			params: &types.UserLookupUsersByParams{
				Usernames:  []string{"test-un"},
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointBase + "?user.fields=uf1%2Cuf2&usernames=test-un",
		},
		{
			name: "normal: all query parameters",
			params: &types.UserLookupUsersByParams{
				Usernames:   []string{"test-un"},
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointBase + "?expansions=ex&tweet.fields=tf&user.fields=uf&usernames=test-un",
		},
		{
			name: "normal: has no required parameter",
			params: &types.UserLookupUsersByParams{
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

func Test_UserLookupUsersByParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UserLookupUsersByParams
	}{
		{
			name:   "empty params",
			params: &types.UserLookupUsersByParams{},
		},
		{
			name:   "some params",
			params: &types.UserLookupUsersByParams{Usernames: []string{"id"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}

func Test_UserLookupUsersByUsernameParams_SetAccessToken(t *testing.T) {
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
			p := &types.UserLookupUsersByUsernameParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UserLookupUsersByUsernameParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:username"
	cases := []struct {
		name   string
		params *types.UserLookupUsersByUsernameParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.UserLookupUsersByUsernameParams{Username: "test-un"},
			expect: endpointRoot + "test-un",
		},
		{
			name: "normal: with expansions",
			params: &types.UserLookupUsersByUsernameParams{
				Username:   "test-un",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-un?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.UserLookupUsersByUsernameParams{
				Username:    "test-un",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-un?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.UserLookupUsersByUsernameParams{
				Username:   "test-un",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-un?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.UserLookupUsersByUsernameParams{
				Username:    "test-un",
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointRoot + "test-un?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.UserLookupUsersByUsernameParams{
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

func Test_UserLookupUsersByUsernameParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UserLookupUsersByUsernameParams
	}{
		{
			name:   "empty params",
			params: &types.UserLookupUsersByUsernameParams{},
		},
		{
			name:   "some params",
			params: &types.UserLookupUsersByUsernameParams{Username: "id"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}
