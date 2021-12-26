package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/lists/types"
	"github.com/stretchr/testify/assert"
)

func Test_ManageListsPostParams_SetAccessToken(t *testing.T) {
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
			p := &types.ManageListsPostParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

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

func Test_ManageListsPostParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListsPostParams
		expect map[string]string
	}{
		{
			name:   "normal: some parameters",
			params: &types.ManageListsPostParams{Name: gotwi.String("name")},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.ManageListsPostParams{},
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

func Test_ManageListsPutParams_SetAccessToken(t *testing.T) {
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
			p := &types.ManageListsPutParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ManageListsPutParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.ManageListsPutParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ManageListsPutParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ManageListsPutParams{},
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

func Test_ManageListsPutParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListsPutParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.ManageListsPutParams{
				ID: "test-id",
			},
			expect: strings.NewReader(`{}`),
		},
		{
			name: "ok: has some parameters",
			params: &types.ManageListsPutParams{
				ID:          "test-id",
				Name:        gotwi.String("test-list-name"),
				Description: gotwi.String("test description"),
			},
			expect: strings.NewReader(`{"name":"test-list-name","description":"test description"}`),
		},
		{
			name: "ok: has some parameters",
			params: &types.ManageListsPutParams{
				ID:          "test-id",
				Name:        gotwi.String("test-list-name"),
				Description: gotwi.String("test description"),
				Private:     gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"name":"test-list-name","description":"test description","private":true}`),
		},
		{
			name: "ok: has no required parameters (no effect)",
			params: &types.ManageListsPutParams{
				Name:        gotwi.String("test-list-name"),
				Description: gotwi.String("test description"),
				Private:     gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"name":"test-list-name","description":"test description","private":true}`),
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

func Test_ManageListsPutParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListsPutParams
		expect map[string]string
	}{
		{
			name:   "normal: some parameters",
			params: &types.ManageListsPutParams{Name: gotwi.String("name")},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.ManageListsPutParams{},
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

func Test_ManageListsDeleteParams_SetAccessToken(t *testing.T) {
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
			p := &types.ManageListsDeleteParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ManageListsDeleteParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.ManageListsDeleteParams
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.ManageListsDeleteParams{
				ID: "lid",
			},
			expect: endpointRoot + "lid",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ManageListsDeleteParams{},
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

func Test_ManageListsDeleteParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListsDeleteParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.ManageListsDeleteParams{
				ID: "lid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.ManageListsDeleteParams{},
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

func Test_ManageListsDeleteParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageListsDeleteParams
		expect map[string]string
	}{
		{
			name:   "normal: some parameters",
			params: &types.ManageListsDeleteParams{ID: "id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.ManageListsDeleteParams{},
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
