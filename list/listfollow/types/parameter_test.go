package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/list/listfollow/types"
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
			p := &types.ListFollowersInput{}
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
		params *types.ListFollowersInput
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListFollowersInput{
				ID: "lid",
			},
			expect: endpointRoot + "lid",
		},
		{
			name: "with expansions",
			params: &types.ListFollowersInput{
				ID:         "lid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "lid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with tweets.fields",
			params: &types.ListFollowersInput{
				ID:          "lid",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "lid" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.ListFollowersInput{
				ID:         "lid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "lid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "with max_results",
			params: &types.ListFollowersInput{
				ID:         "lid",
				MaxResults: 10,
			},
			expect: endpointRoot + "lid" + "?max_results=10",
		},
		{
			name: "with pagination_token",
			params: &types.ListFollowersInput{
				ID:              "lid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "lid" + "?pagination_token=ptoken",
		},
		{
			name: "all query parameters",
			params: &types.ListFollowersInput{
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
			params: &types.ListFollowersInput{
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
		params *types.ListFollowersInput
	}{
		{
			name:   "empty params",
			params: &types.ListFollowersInput{},
		},
		{
			name:   "some params",
			params: &types.ListFollowersInput{ID: "sid"},
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
			p := &types.ListFollowedInput{}
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
		params *types.ListFollowedInput
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListFollowedInput{
				ID: "lid",
			},
			expect: endpointRoot + "lid",
		},
		{
			name: "with expansions",
			params: &types.ListFollowedInput{
				ID:         "lid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "lid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with list.fields",
			params: &types.ListFollowedInput{
				ID:         "lid",
				ListFields: fields.ListFieldList{"lf1", "lf2"},
			},
			expect: endpointRoot + "lid" + "?list.fields=lf1%2Clf2",
		},
		{
			name: "with users.fields",
			params: &types.ListFollowedInput{
				ID:         "lid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "lid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "with max_results",
			params: &types.ListFollowedInput{
				ID:         "lid",
				MaxResults: 10,
			},
			expect: endpointRoot + "lid" + "?max_results=10",
		},
		{
			name: "with pagination_token",
			params: &types.ListFollowedInput{
				ID:              "lid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "lid" + "?pagination_token=ptoken",
		},
		{
			name: "all query parameters",
			params: &types.ListFollowedInput{
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
			params: &types.ListFollowedInput{
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
		params *types.ListFollowedInput
	}{
		{
			name:   "empty params",
			params: &types.ListFollowedInput{},
		},
		{
			name:   "some params",
			params: &types.ListFollowedInput{ID: "sid"},
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
	const endpointBase = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.CreateInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.CreateInput{ID: "uid"},
			expect: endpointBase,
		},
		{
			name:   "normal: has no required parameter",
			params: &types.CreateInput{},
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
				ID:     "uid",
				ListID: "lid",
			},
			expect: strings.NewReader(`{"list_id":"lid"}`),
		},
		{
			name: "ok: has no json parameters",
			params: &types.CreateInput{
				ID: "uid",
			},
			expect: strings.NewReader(`{"list_id":""}`),
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
			params: &types.CreateInput{ID: "id", ListID: "lid"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no required parameter",
			params: &types.CreateInput{},
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
	const endpointBase = "test/endpoint/:id/:list_id"
	cases := []struct {
		name   string
		params *types.DeleteInput
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.DeleteInput{
				ID:     "uid",
				ListID: "lid",
			},
			expect: endpointRoot + "uid" + "/" + "lid",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.DeleteInput{},
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
				ID:     "uid",
				ListID: "lid",
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
			params: &types.DeleteInput{ID: "id", ListID: "lid"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no required parameter",
			params: &types.DeleteInput{},
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
