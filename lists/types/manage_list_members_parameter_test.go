package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
	"github.com/stretchr/testify/assert"
)

func Test_ManageListMembersPostParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.ManageListMembersPostParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ManageListMembersPostParams{ID: "lid"},
			expect: endpointBase,
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ManageListMembersPostParams{},
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

func Test_ManageListMembersPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListMembersPostParams
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.ManageListMembersPostParams{
				ID:     "lid",
				UserID: gotwi.String("uid"),
			},
			expect: strings.NewReader(`{"user_id":"uid"}`),
		},
		{
			name: "ok: has no json parameters",
			params: &types.ManageListMembersPostParams{
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

func Test_ManageListMembersDeleteParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id/:user_id"
	cases := []struct {
		name   string
		params *types.ManageListMembersDeleteParams
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.ManageListMembersDeleteParams{
				ID:     "lid",
				UserID: "uid",
			},
			expect: endpointRoot + "lid" + "/" + "uid",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ManageListMembersDeleteParams{},
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

func Test_ManageListMembersDeleteParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListMembersDeleteParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.ManageListMembersDeleteParams{
				ID:     "lid",
				UserID: "uid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.ManageListMembersDeleteParams{},
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
