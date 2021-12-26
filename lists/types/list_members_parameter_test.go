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

func Test_ListMembersListMemberships_SetAccessToken(t *testing.T) {
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
			p := &types.ListMembersListMembershipsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListMembersListMemberships_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListMembersListMembershipsParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListMembersListMembershipsParams{
				ID: "lid",
			},
			expect: endpointRoot + "lid",
		},
		{
			name: "with expansions",
			params: &types.ListMembersListMembershipsParams{
				ID:         "lid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "lid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with list.fields",
			params: &types.ListMembersListMembershipsParams{
				ID:         "lid",
				ListFields: fields.ListFieldList{"lf1", "lf2"},
			},
			expect: endpointRoot + "lid" + "?list.fields=lf1%2Clf2",
		},
		{
			name: "with users.fields",
			params: &types.ListMembersListMembershipsParams{
				ID:         "lid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "lid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "with max_results",
			params: &types.ListMembersListMembershipsParams{
				ID:         "lid",
				MaxResults: 10,
			},
			expect: endpointRoot + "lid" + "?max_results=10",
		},
		{
			name: "with pagination_token",
			params: &types.ListMembersListMembershipsParams{
				ID:              "lid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "lid" + "?pagination_token=ptoken",
		},
		{
			name: "all query parameters",
			params: &types.ListMembersListMembershipsParams{
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
			params: &types.ListMembersListMembershipsParams{
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

func Test_ListMembersListMemberships_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListMembersListMembershipsParams
	}{
		{
			name:   "empty params",
			params: &types.ListMembersListMembershipsParams{},
		},
		{
			name:   "some params",
			params: &types.ListMembersListMembershipsParams{ID: "sid"},
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

func Test_ListMembersGet_SetAccessToken(t *testing.T) {
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
			p := &types.ListMembersGetParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListMembersGet_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListMembersGetParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListMembersGetParams{
				ID: "uid",
			},
			expect: endpointRoot + "uid",
		},
		{
			name: "with expansions",
			params: &types.ListMembersGetParams{
				ID:         "uid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "uid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with list.fields",
			params: &types.ListMembersGetParams{
				ID:         "uid",
				ListFields: fields.ListFieldList{"lf1", "lf2"},
			},
			expect: endpointRoot + "uid" + "?list.fields=lf1%2Clf2",
		},
		{
			name: "with users.fields",
			params: &types.ListMembersGetParams{
				ID:         "uid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "uid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "with max_results",
			params: &types.ListMembersGetParams{
				ID:         "uid",
				MaxResults: 10,
			},
			expect: endpointRoot + "uid" + "?max_results=10",
		},
		{
			name: "with pagination_token",
			params: &types.ListMembersGetParams{
				ID:              "uid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "uid" + "?pagination_token=ptoken",
		},
		{
			name: "all query parameters",
			params: &types.ListMembersGetParams{
				Expansions:      fields.ExpansionList{"ex"},
				ID:              "uid",
				ListFields:      fields.ListFieldList{"lf"},
				MaxResults:      10,
				PaginationToken: "ptoken",
				UserFields:      fields.UserFieldList{"uf"},
			},
			expect: endpointRoot + "uid" + "?expansions=ex&list.fields=lf&max_results=10&pagination_token=ptoken&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListMembersGetParams{
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

func Test_ListMembersGet_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListMembersGetParams
	}{
		{
			name:   "empty params",
			params: &types.ListMembersGetParams{},
		},
		{
			name:   "some params",
			params: &types.ListMembersGetParams{ID: "sid"},
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

func Test_ListMembersPostParams_SetAccessToken(t *testing.T) {
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
			p := &types.ListMembersPostParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListMembersPostParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.ListMembersPostParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListMembersPostParams{ID: "lid"},
			expect: endpointBase,
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ListMembersPostParams{},
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

func Test_ListMembersPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListMembersPostParams
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.ListMembersPostParams{
				ID:     "lid",
				UserID: gotwi.String("uid"),
			},
			expect: strings.NewReader(`{"user_id":"uid"}`),
		},
		{
			name: "ok: has no json parameters",
			params: &types.ListMembersPostParams{
				ID: "lid",
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

func Test_ListMembersPostParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListMembersPostParams
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListMembersPostParams{ID: "id", UserID: gotwi.String("uid")},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ListMembersPostParams{},
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

func Test_ListMembersDeleteParams_SetAccessToken(t *testing.T) {
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
			p := &types.ListMembersDeleteParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListMembersDeleteParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id/:user_id"
	cases := []struct {
		name   string
		params *types.ListMembersDeleteParams
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.ListMembersDeleteParams{
				ID:     "lid",
				UserID: "uid",
			},
			expect: endpointRoot + "lid" + "/" + "uid",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ListMembersDeleteParams{},
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

func Test_ListMembersDeleteParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListMembersDeleteParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.ListMembersDeleteParams{
				ID:     "lid",
				UserID: "uid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.ListMembersDeleteParams{},
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

func Test_ListMembersDeleteParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListMembersDeleteParams
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListMembersDeleteParams{ID: "id", UserID: "uid"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ListMembersDeleteParams{},
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
