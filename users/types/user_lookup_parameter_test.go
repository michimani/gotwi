package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/users/types"
	"github.com/stretchr/testify/assert"
)

func Test_UserLookupParams_SetAccessToken(t *testing.T) {
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
			p := &types.UserLookupParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UserLookupParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.UserLookupParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.UserLookupParams{IDs: []string{"test-id1", "test-id2"}},
			expect: endpointBase + "?ids=test-id1%2Ctest-id2",
		},
		{
			name: "normal: with expansions",
			params: &types.UserLookupParams{
				IDs:        []string{"test-id"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=test-id",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.UserLookupParams{
				IDs:         []string{"test-id"},
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.UserLookupParams{
				IDs:        []string{"test-id"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=test-id&user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.UserLookupParams{
				IDs:         []string{"test-id"},
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointBase + "?expansions=ex&ids=test-id&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.UserLookupParams{
				IDs:         []string{},
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
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

func Test_UserLookupParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UserLookupParams
	}{
		{
			name:   "empty params",
			params: &types.UserLookupParams{},
		},
		{
			name:   "some params",
			params: &types.UserLookupParams{IDs: []string{"id"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}

func Test_UserLookupIDParams_SetAccessToken(t *testing.T) {
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
			p := &types.UserLookupIDParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UserLookupIDParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.UserLookupIDParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.UserLookupIDParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: with expansions",
			params: &types.UserLookupIDParams{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.UserLookupIDParams{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.UserLookupIDParams{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.UserLookupIDParams{
				ID:          "test-id",
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "test-id?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.UserLookupIDParams{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
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

func Test_UserLookupIDParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UserLookupIDParams
	}{
		{
			name:   "empty params",
			params: &types.UserLookupIDParams{},
		},
		{
			name:   "some params",
			params: &types.UserLookupIDParams{ID: "id"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}

func Test_UserLookupByParams_SetAccessToken(t *testing.T) {
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
			p := &types.UserLookupByParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UserLookupByParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.UserLookupByParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.UserLookupByParams{Usernames: []string{"test-un1", "test-un2"}},
			expect: endpointBase + "?usernames=test-un1%2Ctest-un2",
		},
		{
			name: "normal: with expansions",
			params: &types.UserLookupByParams{
				Usernames:  []string{"test-un"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&usernames=test-un",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.UserLookupByParams{
				Usernames:   []string{"test-un"},
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?tweet.fields=tf1%2Ctf2&usernames=test-un",
		},
		{
			name: "normal: with users.fields",
			params: &types.UserLookupByParams{
				Usernames:  []string{"test-un"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?user.fields=uf1%2Cuf2&usernames=test-un",
		},
		{
			name: "normal: all query parameters",
			params: &types.UserLookupByParams{
				Usernames:   []string{"test-un"},
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointBase + "?expansions=ex&tweet.fields=tf&user.fields=uf&usernames=test-un",
		},
		{
			name: "normal: has no required parameter",
			params: &types.UserLookupByParams{
				Usernames:   []string{},
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
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

func Test_UserLookupByParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UserLookupByParams
	}{
		{
			name:   "empty params",
			params: &types.UserLookupByParams{},
		},
		{
			name:   "some params",
			params: &types.UserLookupByParams{Usernames: []string{"id"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}

func Test_UserLookupByUsernameParams_SetAccessToken(t *testing.T) {
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
			p := &types.UserLookupByUsernameParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UserLookupByUsernameParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:username"
	cases := []struct {
		name   string
		params *types.UserLookupByUsernameParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.UserLookupByUsernameParams{Username: "test-un"},
			expect: endpointRoot + "test-un",
		},
		{
			name: "normal: with expansions",
			params: &types.UserLookupByUsernameParams{
				Username:   "test-un",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-un?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.UserLookupByUsernameParams{
				Username:    "test-un",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-un?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.UserLookupByUsernameParams{
				Username:   "test-un",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-un?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.UserLookupByUsernameParams{
				Username:    "test-un",
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "test-un?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.UserLookupByUsernameParams{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
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

func Test_UserLookupByUsernameParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UserLookupByUsernameParams
	}{
		{
			name:   "empty params",
			params: &types.UserLookupByUsernameParams{},
		},
		{
			name:   "some params",
			params: &types.UserLookupByUsernameParams{Username: "id"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}

func Test_UserLookupMeParams_SetAccessToken(t *testing.T) {
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
			p := &types.UserLookupMeParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UserLookupMeParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.UserLookupMeParams
		expect string
	}{
		{
			name:   "normal: has no parameter",
			params: &types.UserLookupMeParams{},
			expect: endpointBase,
		},
		{
			name: "normal: with expansions",
			params: &types.UserLookupMeParams{
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.UserLookupMeParams{
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.UserLookupMeParams{
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.UserLookupMeParams{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointBase + "?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_UserLookupMeParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UserLookupMeParams
	}{
		{
			name:   "empty params",
			params: &types.UserLookupMeParams{},
		},
		{
			name:   "some params",
			params: &types.UserLookupMeParams{Expansions: fields.ExpansionList{"ex"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}
