package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
	"github.com/stretchr/testify/assert"
)

func Test_ManageListsPostParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.ManageListsPostParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ManageListsPostParams{Name: gotwi.String("test-list-name")},
			expect: endpointBase,
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ManageListsPostParams{},
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

func Test_ManageListsPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListsPostParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.ManageListsPostParams{
				Name: gotwi.String("test-list-name"),
			},
			expect: strings.NewReader(`{"name":"test-list-name"}`),
		},
		{
			name: "ok: has some parameters",
			params: &types.ManageListsPostParams{
				Name:        gotwi.String("test-list-name"),
				Description: gotwi.String("test description"),
			},
			expect: strings.NewReader(`{"name":"test-list-name","description":"test description"}`),
		},
		{
			name: "ok: has some parameters",
			params: &types.ManageListsPostParams{
				Name:        gotwi.String("test-list-name"),
				Description: gotwi.String("test description"),
				Private:     gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"name":"test-list-name","description":"test description","private":true}`),
		},
		{
			name:   "ok: has no required parameters",
			params: &types.ManageListsPostParams{},
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
