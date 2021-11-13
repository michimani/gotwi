package types_test

import (
	"testing"

	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_FilteredStreamRulesGet_SetAccessToken(t *testing.T) {
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
			p := &types.FilteredStreamRulesGetParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_FilteredStreamRulesGet_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"

	cases := []struct {
		name   string
		params *types.FilteredStreamRulesGetParams
		expect string
	}{
		{
			name:   "has no parameter",
			params: &types.FilteredStreamRulesGetParams{},
			expect: endpointBase,
		},
		{
			name: "with ids",
			params: &types.FilteredStreamRulesGetParams{
				IDs: []string{"rid1", "rid2"},
			},
			expect: endpointBase + "?ids=rid1%2Crid2",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_FilteredStreamRulesGet_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.FilteredStreamRulesGetParams
	}{
		{
			name:   "empty params",
			params: &types.FilteredStreamRulesGetParams{},
		},
		{
			name:   "some params",
			params: &types.FilteredStreamRulesGetParams{IDs: []string{"rid1", "rid2"}},
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
