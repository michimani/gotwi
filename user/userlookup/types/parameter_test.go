package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/user/userlookup/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListInput_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.ListInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListInput{IDs: []string{"test-id1", "test-id2"}},
			expect: endpointBase + "?ids=test-id1%2Ctest-id2",
		},
		{
			name: "normal: with expansions",
			params: &types.ListInput{
				IDs:        []string{"test-id"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=test-id",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.ListInput{
				IDs:         []string{"test-id"},
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.ListInput{
				IDs:        []string{"test-id"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=test-id&user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.ListInput{
				IDs:         []string{"test-id"},
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointBase + "?expansions=ex&ids=test-id&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.ListInput{
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

func Test_ListInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListInput
	}{
		{
			name:   "empty params",
			params: &types.ListInput{},
		},
		{
			name:   "some params",
			params: &types.ListInput{IDs: []string{"id"}},
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

func Test_GetInput_SetAccessToken(t *testing.T) {
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
			p := &types.GetInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_GetInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.GetInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.GetInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: with expansions",
			params: &types.GetInput{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.GetInput{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.GetInput{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.GetInput{
				ID:          "test-id",
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "test-id?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.GetInput{
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

func Test_GetInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.GetInput
	}{
		{
			name:   "empty params",
			params: &types.GetInput{},
		},
		{
			name:   "some params",
			params: &types.GetInput{ID: "id"},
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

func Test_ListByUsernamesInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListByUsernamesInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListByUsernamesInput_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.ListByUsernamesInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListByUsernamesInput{Usernames: []string{"test-un1", "test-un2"}},
			expect: endpointBase + "?usernames=test-un1%2Ctest-un2",
		},
		{
			name: "normal: with expansions",
			params: &types.ListByUsernamesInput{
				Usernames:  []string{"test-un"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&usernames=test-un",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.ListByUsernamesInput{
				Usernames:   []string{"test-un"},
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?tweet.fields=tf1%2Ctf2&usernames=test-un",
		},
		{
			name: "normal: with users.fields",
			params: &types.ListByUsernamesInput{
				Usernames:  []string{"test-un"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?user.fields=uf1%2Cuf2&usernames=test-un",
		},
		{
			name: "normal: all query parameters",
			params: &types.ListByUsernamesInput{
				Usernames:   []string{"test-un"},
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointBase + "?expansions=ex&tweet.fields=tf&user.fields=uf&usernames=test-un",
		},
		{
			name: "normal: has no required parameter",
			params: &types.ListByUsernamesInput{
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

func Test_ListByUsernamesInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListByUsernamesInput
	}{
		{
			name:   "empty params",
			params: &types.ListByUsernamesInput{},
		},
		{
			name:   "some params",
			params: &types.ListByUsernamesInput{Usernames: []string{"id"}},
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

func Test_GetByUsernameInput_SetAccessToken(t *testing.T) {
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
			p := &types.GetByUsernameInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_GetByUsernameInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:username"
	cases := []struct {
		name   string
		params *types.GetByUsernameInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.GetByUsernameInput{Username: "test-un"},
			expect: endpointRoot + "test-un",
		},
		{
			name: "normal: with expansions",
			params: &types.GetByUsernameInput{
				Username:   "test-un",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-un?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.GetByUsernameInput{
				Username:    "test-un",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-un?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.GetByUsernameInput{
				Username:   "test-un",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-un?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.GetByUsernameInput{
				Username:    "test-un",
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "test-un?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.GetByUsernameInput{
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

func Test_GetByUsernameInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.GetByUsernameInput
	}{
		{
			name:   "empty params",
			params: &types.GetByUsernameInput{},
		},
		{
			name:   "some params",
			params: &types.GetByUsernameInput{Username: "id"},
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
