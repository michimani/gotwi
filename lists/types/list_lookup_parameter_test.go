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
