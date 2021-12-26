package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/lists/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListFollowsFollowers_SetAccessToken(t *testing.T) {
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
			p := &types.ListFollowsFollowersParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListFollowsFollowers_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListFollowsFollowersParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListFollowsFollowersParams{
				ID: "lid",
			},
			expect: endpointRoot + "lid",
		},
		{
			name: "with expansions",
			params: &types.ListFollowsFollowersParams{
				ID:         "lid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "lid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with tweets.fields",
			params: &types.ListFollowsFollowersParams{
				ID:          "lid",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "lid" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.ListFollowsFollowersParams{
				ID:         "lid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "lid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "with max_results",
			params: &types.ListFollowsFollowersParams{
				ID:         "lid",
				MaxResults: 10,
			},
			expect: endpointRoot + "lid" + "?max_results=10",
		},
		{
			name: "with pagination_token",
			params: &types.ListFollowsFollowersParams{
				ID:              "lid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "lid" + "?pagination_token=ptoken",
		},
		{
			name: "all query parameters",
			params: &types.ListFollowsFollowersParams{
				Expansions:      fields.ExpansionList{"ex"},
				ID:              "lid",
				TweetFields:     fields.TweetFieldList{"tf"},
				MaxResults:      10,
				PaginationToken: "ptoken",
				UserFields:      fields.UserFieldList{"uf"},
			},
			expect: endpointRoot + "lid" + "?expansions=ex&max_results=10&pagination_token=ptoken&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListFollowsFollowersParams{
				Expansions:      fields.ExpansionList{"ex"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
				MaxResults:      10,
				PaginationToken: "pagination_token",
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

func Test_ListFollowsFollowers_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListFollowsFollowersParams
	}{
		{
			name:   "empty params",
			params: &types.ListFollowsFollowersParams{},
		},
		{
			name:   "some params",
			params: &types.ListFollowsFollowersParams{ID: "sid"},
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

func Test_ListFollowsFollowedLists_SetAccessToken(t *testing.T) {
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
			p := &types.ListFollowsFollowedListsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListFollowsFollowedLists_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListFollowsFollowedListsParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListFollowsFollowedListsParams{
				ID: "lid",
			},
			expect: endpointRoot + "lid",
		},
		{
			name: "with expansions",
			params: &types.ListFollowsFollowedListsParams{
				ID:         "lid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "lid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with list.fields",
			params: &types.ListFollowsFollowedListsParams{
				ID:         "lid",
				ListFields: fields.ListFieldList{"lf1", "lf2"},
			},
			expect: endpointRoot + "lid" + "?list.fields=lf1%2Clf2",
		},
		{
			name: "with users.fields",
			params: &types.ListFollowsFollowedListsParams{
				ID:         "lid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "lid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "with max_results",
			params: &types.ListFollowsFollowedListsParams{
				ID:         "lid",
				MaxResults: 10,
			},
			expect: endpointRoot + "lid" + "?max_results=10",
		},
		{
			name: "with pagination_token",
			params: &types.ListFollowsFollowedListsParams{
				ID:              "lid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "lid" + "?pagination_token=ptoken",
		},
		{
			name: "all query parameters",
			params: &types.ListFollowsFollowedListsParams{
				Expansions:      fields.ExpansionList{"ex"},
				ID:              "lid",
				ListFields:      fields.ListFieldList{"lf"},
				MaxResults:      10,
				PaginationToken: "ptoken",
				UserFields:      fields.UserFieldList{"uf"},
			},
			expect: endpointRoot + "lid" + "?expansions=ex&list.fields=lf&max_results=10&pagination_token=ptoken&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListFollowsFollowedListsParams{
				Expansions:      fields.ExpansionList{"ex"},
				UserFields:      fields.UserFieldList{"uf"},
				ListFields:      fields.ListFieldList{"lf"},
				MaxResults:      10,
				PaginationToken: "pagination_token",
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

func Test_ListFollowsFollowedLists_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListFollowsFollowedListsParams
	}{
		{
			name:   "empty params",
			params: &types.ListFollowsFollowedListsParams{},
		},
		{
			name:   "some params",
			params: &types.ListFollowsFollowedListsParams{ID: "sid"},
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

func Test_ListFollowsPostParams_SetAccessToken(t *testing.T) {
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
			p := &types.ListFollowsPostParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListFollowsPostParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.ListFollowsPostParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListFollowsPostParams{ID: "uid"},
			expect: endpointBase,
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ListFollowsPostParams{},
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

func Test_ListFollowsPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListFollowsPostParams
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.ListFollowsPostParams{
				ID:     "uid",
				ListID: gotwi.String("lid"),
			},
			expect: strings.NewReader(`{"list_id":"lid"}`),
		},
		{
			name: "ok: has no json parameters",
			params: &types.ListFollowsPostParams{
				ID: "uid",
			},
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

func Test_ListFollowsPostParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListFollowsPostParams
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListFollowsPostParams{ID: "id", ListID: gotwi.String("lid")},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ListFollowsPostParams{},
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

func Test_ListFollowsDeleteParams_SetAccessToken(t *testing.T) {
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
			p := &types.ListFollowsDeleteParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListFollowsDeleteParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id/:list_id"
	cases := []struct {
		name   string
		params *types.ListFollowsDeleteParams
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.ListFollowsDeleteParams{
				ID:     "uid",
				ListID: "lid",
			},
			expect: endpointRoot + "uid" + "/" + "lid",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ListFollowsDeleteParams{},
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

func Test_ListFollowsDeleteParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListFollowsDeleteParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.ListFollowsDeleteParams{
				ID:     "uid",
				ListID: "lid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.ListFollowsDeleteParams{},
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

func Test_ListFollowsDeleteParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListFollowsDeleteParams
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListFollowsDeleteParams{ID: "id", ListID: "lid"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ListFollowsDeleteParams{},
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
