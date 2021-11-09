package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/spaces/types"
	"github.com/stretchr/testify/assert"
)

func Test_SearchSpaces_SetAccessToken(t *testing.T) {
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
			p := &types.SearchSpacesParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SearchSpaces_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"

	cases := []struct {
		name   string
		params *types.SearchSpacesParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SearchSpacesParams{
				Query: "hello",
			},
			expect: endpointBase + "?query=hello",
		},
		{
			name: "with expansions",
			params: &types.SearchSpacesParams{
				Query:      "hello",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&query=hello",
		},
		{
			name: "with max_result",
			params: &types.SearchSpacesParams{
				Query:      "hello",
				MaxResults: 10,
			},
			expect: endpointBase + "?max_results=10&query=hello",
		},
		{
			name: "with space.fields",
			params: &types.SearchSpacesParams{
				Query:       "hello",
				SpaceFields: fields.SpaceFieldList{"sf1", "sf2"},
			},
			expect: endpointBase + "?query=hello&space.fields=sf1%2Csf2",
		},
		{
			name: "with users.fields",
			params: &types.SearchSpacesParams{
				Query:      "hello",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?query=hello&user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SearchSpacesParams{
				Expansions:  fields.ExpansionList{"ex"},
				MaxResults:  10,
				Query:       "hello",
				SpaceFields: fields.SpaceFieldList{"sf"},
				State:       fields.StateAll,
				UserFields:  fields.UserFieldList{"uf"},
			},
			expect: endpointBase + "?expansions=ex&max_results=10&query=hello&space.fields=sf&state=all&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.SearchSpacesParams{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				SpaceFields: fields.SpaceFieldList{"sf"},
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

func Test_SearchSpaces_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SearchSpacesParams
	}{
		{
			name:   "empty params",
			params: &types.SearchSpacesParams{},
		},
		{
			name:   "some params",
			params: &types.SearchSpacesParams{Query: "hello"},
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
