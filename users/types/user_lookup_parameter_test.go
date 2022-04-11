package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/users/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListUsersInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListUsersInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListUsersInput_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.ListUsersInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListUsersInput{IDs: []string{"test-id1", "test-id2"}},
			expect: endpointBase + "?ids=test-id1%2Ctest-id2",
		},
		{
			name: "normal: with expansions",
			params: &types.ListUsersInput{
				IDs:        []string{"test-id"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=test-id",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.ListUsersInput{
				IDs:         []string{"test-id"},
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.ListUsersInput{
				IDs:        []string{"test-id"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=test-id&user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.ListUsersInput{
				IDs:         []string{"test-id"},
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointBase + "?expansions=ex&ids=test-id&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.ListUsersInput{
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

func Test_ListUsersInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListUsersInput
	}{
		{
			name:   "empty params",
			params: &types.ListUsersInput{},
		},
		{
			name:   "some params",
			params: &types.ListUsersInput{IDs: []string{"id"}},
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

func Test_GetUserInput_SetAccessToken(t *testing.T) {
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
			p := &types.GetUserInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_GetUserInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.GetUserInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.GetUserInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: with expansions",
			params: &types.GetUserInput{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.GetUserInput{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.GetUserInput{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.GetUserInput{
				ID:          "test-id",
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "test-id?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.GetUserInput{
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

func Test_GetUserInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.GetUserInput
	}{
		{
			name:   "empty params",
			params: &types.GetUserInput{},
		},
		{
			name:   "some params",
			params: &types.GetUserInput{ID: "id"},
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

func Test_ListUsersByUsernamesInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListUsersByUsernamesInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListUsersByUsernamesInput_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.ListUsersByUsernamesInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListUsersByUsernamesInput{Usernames: []string{"test-un1", "test-un2"}},
			expect: endpointBase + "?usernames=test-un1%2Ctest-un2",
		},
		{
			name: "normal: with expansions",
			params: &types.ListUsersByUsernamesInput{
				Usernames:  []string{"test-un"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&usernames=test-un",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.ListUsersByUsernamesInput{
				Usernames:   []string{"test-un"},
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?tweet.fields=tf1%2Ctf2&usernames=test-un",
		},
		{
			name: "normal: with users.fields",
			params: &types.ListUsersByUsernamesInput{
				Usernames:  []string{"test-un"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?user.fields=uf1%2Cuf2&usernames=test-un",
		},
		{
			name: "normal: all query parameters",
			params: &types.ListUsersByUsernamesInput{
				Usernames:   []string{"test-un"},
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointBase + "?expansions=ex&tweet.fields=tf&user.fields=uf&usernames=test-un",
		},
		{
			name: "normal: has no required parameter",
			params: &types.ListUsersByUsernamesInput{
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

func Test_ListUsersByUsernamesInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListUsersByUsernamesInput
	}{
		{
			name:   "empty params",
			params: &types.ListUsersByUsernamesInput{},
		},
		{
			name:   "some params",
			params: &types.ListUsersByUsernamesInput{Usernames: []string{"id"}},
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

func Test_GetUserByUsernameInput_SetAccessToken(t *testing.T) {
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
			p := &types.GetUserByUsernameInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_GetUserByUsernameInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:username"
	cases := []struct {
		name   string
		params *types.GetUserByUsernameInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.GetUserByUsernameInput{Username: "test-un"},
			expect: endpointRoot + "test-un",
		},
		{
			name: "normal: with expansions",
			params: &types.GetUserByUsernameInput{
				Username:   "test-un",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-un?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.GetUserByUsernameInput{
				Username:    "test-un",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-un?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.GetUserByUsernameInput{
				Username:   "test-un",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-un?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.GetUserByUsernameInput{
				Username:    "test-un",
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "test-un?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.GetUserByUsernameInput{
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

func Test_GetUserByUsernameInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.GetUserByUsernameInput
	}{
		{
			name:   "empty params",
			params: &types.GetUserByUsernameInput{},
		},
		{
			name:   "some params",
			params: &types.GetUserByUsernameInput{Username: "id"},
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

func Test_GetMeInput_SetAccessToken(t *testing.T) {
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
			p := &types.GetMeInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_GetMeInput_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.GetMeInput
		expect string
	}{
		{
			name:   "normal: has no parameter",
			params: &types.GetMeInput{},
			expect: endpointBase,
		},
		{
			name: "normal: with expansions",
			params: &types.GetMeInput{
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.GetMeInput{
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.GetMeInput{
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.GetMeInput{
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

func Test_GetMeInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.GetMeInput
	}{
		{
			name:   "empty params",
			params: &types.GetMeInput{},
		},
		{
			name:   "some params",
			params: &types.GetMeInput{Expansions: fields.ExpansionList{"ex"}},
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
