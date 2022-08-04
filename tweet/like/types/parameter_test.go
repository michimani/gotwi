package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweet/like/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListMaxResults_String(t *testing.T) {
	cases := []struct {
		name   string
		m      types.ListMaxResults
		expect string
	}{
		{
			name:   "normal",
			m:      types.ListMaxResults(1),
			expect: "1",
		},
		{
			name:   "normal: zero",
			m:      types.ListMaxResults(0),
			expect: "0",
		},
		{
			name:   "normal: negative",
			m:      types.ListMaxResults(-1),
			expect: "-1",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			s := c.m.String()
			assert.Equal(tt, c.expect, s)
		})
	}
}

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
			name:   "empty",
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
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.ListUsersInput
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.ListUsersInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "with expansions",
			params: &types.ListUsersInput{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with tweets.fields",
			params: &types.ListUsersInput{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.ListUsersInput{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.ListUsersInput{
				ID:          "test-id",
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListUsersInput{
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
			params: &types.ListUsersInput{ID: "id"},
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
			name:   "empty",
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
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.ListInput
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.ListInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "with expansions",
			params: &types.ListInput{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.ListInput{
				ID:          "test-id",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.ListInput{
				ID:          "test-id",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.ListInput{
				ID:         "test-id",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.ListInput{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.ListInput{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.ListInput{
				ID:              "test-id",
				Expansions:      fields.ExpansionList{"ex"},
				MediaFields:     fields.MediaFieldList{"mf"},
				PlaceFields:     fields.PlaceFieldList{"plf"},
				PollFields:      fields.PollFieldList{"pof"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
				MaxResults:      types.ListMaxResults(10),
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "test-id" + "?expansions=ex&max_results=10&media.fields=mf&pagination_token=ptoken&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListInput{
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
			params: &types.ListInput{ID: "id"},
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

func Test_CreateInput_SetAccessToken(t *testing.T) {
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
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.CreateInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_CreateInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.CreateInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.CreateInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: has no required parameter",
			params: &types.CreateInput{
				TweetID: "tid",
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

func Test_CreateInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateInput
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.CreateInput{
				ID:      "test-id",
				TweetID: "tid",
			},
			expect: strings.NewReader(`{"tweet_id":"tid"}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.CreateInput{ID: "id"},
			expect: strings.NewReader(`{"tweet_id":""}`),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, r)
		})
	}
}

func Test_CreateInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateInput
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.CreateInput{ID: "test-id"},
			expect: map[string]string{},
		},
		{
			name: "normal: has no required parameter",
			params: &types.CreateInput{
				TweetID: "tid",
			},
			expect: map[string]string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			m := c.params.ParameterMap()
			assert.Equal(tt, c.expect, m)
		})
	}
}

func Test_DeleteInput_SetAccessToken(t *testing.T) {
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
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.DeleteInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_DeleteInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id/:tweet_id"
	cases := []struct {
		name   string
		params *types.DeleteInput
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.DeleteInput{
				ID:      "uid",
				TweetID: "tid",
			},
			expect: endpointRoot + "uid" + "/" + "tid",
		},
		{
			name: "normal: has no required parameter",
			params: &types.DeleteInput{
				ID: "uid",
			},
			expect: "",
		},
		{
			name: "normal: has no required parameter",
			params: &types.DeleteInput{
				TweetID: "tid",
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

func Test_DeleteInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.DeleteInput
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.DeleteInput{
				ID:      "uid",
				TweetID: "tid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.DeleteInput{},
			expect: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, r)
		})
	}
}

func Test_DeleteInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.DeleteInput
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.DeleteInput{ID: "test-id"},
			expect: map[string]string{},
		},
		{
			name: "normal: has no required parameter",
			params: &types.DeleteInput{
				TweetID: "tid",
			},
			expect: map[string]string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			m := c.params.ParameterMap()
			assert.Equal(tt, c.expect, m)
		})
	}
}
