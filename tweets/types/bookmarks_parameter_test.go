package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_BookmarksMaxResults_String(t *testing.T) {
	cases := []struct {
		name   string
		m      types.BookmarksMaxResults
		expect string
	}{
		{
			name:   "normal",
			m:      types.BookmarksMaxResults(1),
			expect: "1",
		},
		{
			name:   "normal: zero",
			m:      types.BookmarksMaxResults(0),
			expect: "0",
		},
		{
			name:   "normal: negative",
			m:      types.BookmarksMaxResults(-1),
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

func Test_BookmarksParams_SetAccessToken(t *testing.T) {
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
			p := &types.BookmarksParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_BookmarksParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.BookmarksParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.BookmarksParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "with expansions",
			params: &types.BookmarksParams{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.BookmarksParams{
				ID:          "test-id",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.BookmarksParams{
				ID:          "test-id",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.BookmarksParams{
				ID:         "test-id",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.BookmarksParams{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.BookmarksParams{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.BookmarksParams{
				ID:              "test-id",
				Expansions:      fields.ExpansionList{"ex"},
				MediaFields:     fields.MediaFieldList{"mf"},
				PlaceFields:     fields.PlaceFieldList{"plf"},
				PollFields:      fields.PollFieldList{"pof"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
				MaxResults:      types.BookmarksMaxResults(10),
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "test-id" + "?expansions=ex&max_results=10&media.fields=mf&pagination_token=ptoken&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.BookmarksParams{
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

func Test_BookmarksParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BookmarksParams
	}{
		{
			name:   "empty params",
			params: &types.BookmarksParams{},
		},
		{
			name:   "some params",
			params: &types.BookmarksParams{ID: "id"},
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

func Test_BookmarksPostParams_SetAccessToken(t *testing.T) {
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
			p := &types.BookmarksPostParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_BookmarksPostParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.BookmarksPostParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.BookmarksPostParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: has no required parameter",
			params: &types.BookmarksPostParams{
				TweetID: gotwi.String("tid"),
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

func Test_BookmarksPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BookmarksPostParams
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.BookmarksPostParams{
				ID:      "test-id",
				TweetID: gotwi.String("tid"),
			},
			expect: strings.NewReader(`{"tweet_id":"tid"}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.BookmarksPostParams{ID: "id"},
			expect: strings.NewReader(`{}`),
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

func Test_BookmarksPostParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BookmarksPostParams
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.BookmarksPostParams{ID: "test-id"},
			expect: map[string]string{},
		},
		{
			name: "normal: has no required parameter",
			params: &types.BookmarksPostParams{
				TweetID: gotwi.String("tid"),
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
