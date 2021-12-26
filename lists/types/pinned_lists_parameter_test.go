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

func Test_PinnedListsGet_SetAccessToken(t *testing.T) {
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
			p := &types.PinnedListsGetParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_PinnedListsGet_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.PinnedListsGetParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.PinnedListsGetParams{
				ID: "uid",
			},
			expect: endpointRoot + "uid",
		},
		{
			name: "with expansions",
			params: &types.PinnedListsGetParams{
				ID:         "uid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "uid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with list.fields",
			params: &types.PinnedListsGetParams{
				ID:         "uid",
				ListFields: fields.ListFieldList{"lf1", "lf2"},
			},
			expect: endpointRoot + "uid" + "?list.fields=lf1%2Clf2",
		},
		{
			name: "with users.fields",
			params: &types.PinnedListsGetParams{
				ID:         "uid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "uid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.PinnedListsGetParams{
				Expansions: fields.ExpansionList{"ex"},
				ID:         "uid",
				ListFields: fields.ListFieldList{"lf"},
				UserFields: fields.UserFieldList{"uf"},
			},
			expect: endpointRoot + "uid" + "?expansions=ex&list.fields=lf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.PinnedListsGetParams{
				Expansions: fields.ExpansionList{"ex"},
				UserFields: fields.UserFieldList{"uf"},
				ListFields: fields.ListFieldList{"lf"},
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

func Test_PinnedListsGet_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.PinnedListsGetParams
	}{
		{
			name:   "empty params",
			params: &types.PinnedListsGetParams{},
		},
		{
			name:   "some params",
			params: &types.PinnedListsGetParams{ID: "sid"},
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

func Test_PinnedListsPostParams_SetAccessToken(t *testing.T) {
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
			p := &types.PinnedListsPostParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_PinnedListsPostParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.PinnedListsPostParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.PinnedListsPostParams{ID: "uid"},
			expect: endpointBase,
		},
		{
			name:   "normal: has no required parameter",
			params: &types.PinnedListsPostParams{},
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

func Test_PinnedListsPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.PinnedListsPostParams
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.PinnedListsPostParams{
				ID:     "uid",
				ListID: gotwi.String("lid"),
			},
			expect: strings.NewReader(`{"list_id":"lid"}`),
		},
		{
			name: "ok: has no json parameters",
			params: &types.PinnedListsPostParams{
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

func Test_PinnedListsPostParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.PinnedListsPostParams
		expect map[string]string
	}{
		{
			name:   "normal: some parameters",
			params: &types.PinnedListsPostParams{ID: "id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.PinnedListsPostParams{},
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

func Test_PinnedListsDeleteParams_SetAccessToken(t *testing.T) {
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
			p := &types.PinnedListsDeleteParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_PinnedListsDeleteParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id/:list_id"
	cases := []struct {
		name   string
		params *types.PinnedListsDeleteParams
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.PinnedListsDeleteParams{
				ID:     "uid",
				ListID: "lid",
			},
			expect: endpointRoot + "uid" + "/" + "lid",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.PinnedListsDeleteParams{},
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

func Test_PinnedListsDeleteParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.PinnedListsDeleteParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.PinnedListsDeleteParams{
				ID:     "uid",
				ListID: "lid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.PinnedListsDeleteParams{},
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

func Test_PinnedListsDeleteParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.PinnedListsDeleteParams
		expect map[string]string
	}{
		{
			name:   "normal: some parameters",
			params: &types.PinnedListsDeleteParams{ID: "id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.PinnedListsDeleteParams{},
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
