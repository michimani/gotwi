package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
	"github.com/stretchr/testify/assert"
)

func Test_ManageListFollowsPostParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.ManageListFollowsPostParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ManageListFollowsPostParams{ID: "uid"},
			expect: endpointBase,
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ManageListFollowsPostParams{},
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

func Test_ManageListFollowsPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListFollowsPostParams
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.ManageListFollowsPostParams{
				ID:     "uid",
				ListID: gotwi.String("lid"),
			},
			expect: strings.NewReader(`{"list_id":"lid"}`),
		},
		{
			name: "ok: has no json parameters",
			params: &types.ManageListFollowsPostParams{
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

func Test_ManageListFollowsDeleteParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id/:list_id"
	cases := []struct {
		name   string
		params *types.ManageListFollowsDeleteParams
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.ManageListFollowsDeleteParams{
				ID:     "uid",
				ListID: "lid",
			},
			expect: endpointRoot + "uid" + "/" + "lid",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ManageListFollowsDeleteParams{},
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

func Test_ManageListFollowsDeleteParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListFollowsDeleteParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.ManageListFollowsDeleteParams{
				ID:     "uid",
				ListID: "lid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.ManageListFollowsDeleteParams{},
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
