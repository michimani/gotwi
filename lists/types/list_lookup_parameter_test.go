package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/lists/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListLookupID_SetAccessToken(t *testing.T) {
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
			p := &types.ListLookupIDParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListLookupID_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListLookupIDParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListLookupIDParams{
				ID: "sid",
			},
			expect: endpointRoot + "sid",
		},
		{
			name: "with expansions",
			params: &types.ListLookupIDParams{
				ID:         "sid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with list.fields",
			params: &types.ListLookupIDParams{
				ID:         "sid",
				ListFields: fields.ListFieldList{"lf1", "lf2"},
			},
			expect: endpointRoot + "sid" + "?list.fields=lf1%2Clf2",
		},
		{
			name: "with users.fields",
			params: &types.ListLookupIDParams{
				ID:         "sid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "sid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.ListLookupIDParams{
				Expansions: fields.ExpansionList{"ex"},
				ID:         "sid",
				ListFields: fields.ListFieldList{"lf"},
				UserFields: fields.UserFieldList{"uf"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex&list.fields=lf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListLookupIDParams{
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

func Test_ListLookupID_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListLookupIDParams
	}{
		{
			name:   "empty params",
			params: &types.ListLookupIDParams{},
		},
		{
			name:   "some params",
			params: &types.ListLookupIDParams{ID: "sid"},
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

func Test_ListLookupOwnedLists_SetAccessToken(t *testing.T) {
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
			p := &types.ListLookupOwnedListsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListLookupOwnedLists_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListLookupOwnedListsParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListLookupOwnedListsParams{
				ID: "uid",
			},
			expect: endpointRoot + "uid",
		},
		{
			name: "with expansions",
			params: &types.ListLookupOwnedListsParams{
				ID:         "uid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "uid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with list.fields",
			params: &types.ListLookupOwnedListsParams{
				ID:         "uid",
				ListFields: fields.ListFieldList{"lf1", "lf2"},
			},
			expect: endpointRoot + "uid" + "?list.fields=lf1%2Clf2",
		},
		{
			name: "with users.fields",
			params: &types.ListLookupOwnedListsParams{
				ID:         "uid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "uid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "with max_results",
			params: &types.ListLookupOwnedListsParams{
				ID:         "uid",
				MaxResults: 10,
			},
			expect: endpointRoot + "uid" + "?max_results=10",
		},
		{
			name: "with pagination_token",
			params: &types.ListLookupOwnedListsParams{
				ID:              "uid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "uid" + "?pagination_token=ptoken",
		},
		{
			name: "all query parameters",
			params: &types.ListLookupOwnedListsParams{
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
			params: &types.ListLookupOwnedListsParams{
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

func Test_ListLookupOwnedLists_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListLookupOwnedListsParams
	}{
		{
			name:   "empty params",
			params: &types.ListLookupOwnedListsParams{},
		},
		{
			name:   "some params",
			params: &types.ListLookupOwnedListsParams{ID: "sid"},
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
