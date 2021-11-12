package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/spaces/types"
	"github.com/stretchr/testify/assert"
)

func Test_SpacesLookupID_SetAccessToken(t *testing.T) {
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
			p := &types.SpacesLookupIDParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SpacesLookupID_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.SpacesLookupIDParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SpacesLookupIDParams{
				ID: "sid",
			},
			expect: endpointRoot + "sid",
		},
		{
			name: "with expansions",
			params: &types.SpacesLookupIDParams{
				ID:         "sid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with space.fields",
			params: &types.SpacesLookupIDParams{
				ID:          "sid",
				SpaceFields: fields.SpaceFieldList{"sf1", "sf2"},
			},
			expect: endpointRoot + "sid" + "?space.fields=sf1%2Csf2",
		},
		{
			name: "with users.fields",
			params: &types.SpacesLookupIDParams{
				ID:         "sid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "sid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SpacesLookupIDParams{
				Expansions:  fields.ExpansionList{"ex"},
				ID:          "sid",
				SpaceFields: fields.SpaceFieldList{"sf"},
				UserFields:  fields.UserFieldList{"uf"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex&space.fields=sf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.SpacesLookupIDParams{
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

func Test_SpacesLookupID_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SpacesLookupIDParams
	}{
		{
			name:   "empty params",
			params: &types.SpacesLookupIDParams{},
		},
		{
			name:   "some params",
			params: &types.SpacesLookupIDParams{ID: "sid"},
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

func Test_SpacesLookup_SetAccessToken(t *testing.T) {
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
			p := &types.SpacesLookupParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SpacesLookup_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"

	cases := []struct {
		name   string
		params *types.SpacesLookupParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SpacesLookupParams{
				IDs: []string{"sid1", "sid2"},
			},
			expect: endpointBase + "?ids=sid1%2Csid2",
		},
		{
			name: "with expansions",
			params: &types.SpacesLookupParams{
				IDs:        []string{"sid1", "sid2"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=sid1%2Csid2",
		},
		{
			name: "with space.fields",
			params: &types.SpacesLookupParams{
				IDs:         []string{"sid1", "sid2"},
				SpaceFields: fields.SpaceFieldList{"sf1", "sf2"},
			},
			expect: endpointBase + "?ids=sid1%2Csid2&space.fields=sf1%2Csf2",
		},
		{
			name: "with users.fields",
			params: &types.SpacesLookupParams{
				IDs:        []string{"sid1", "sid2"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=sid1%2Csid2&user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SpacesLookupParams{
				Expansions:  fields.ExpansionList{"ex"},
				IDs:         []string{"sid1", "sid2"},
				SpaceFields: fields.SpaceFieldList{"sf"},
				UserFields:  fields.UserFieldList{"uf"},
			},
			expect: endpointBase + "?expansions=ex&ids=sid1%2Csid2&space.fields=sf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.SpacesLookupParams{
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

func Test_SpacesLookup_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SpacesLookupParams
	}{
		{
			name:   "empty params",
			params: &types.SpacesLookupParams{},
		},
		{
			name:   "some params",
			params: &types.SpacesLookupParams{IDs: []string{"sid"}},
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
