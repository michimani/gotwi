package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/list/managelist/types"
	"github.com/stretchr/testify/assert"
)

func Test_CreateInput_SetAccessToken(t *testing.T) {
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
			p := &types.CreateInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_CreateInput_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.CreateInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.CreateInput{Name: "test-list-name"},
			expect: endpointBase,
		},
		{
			name:   "normal: has no required parameter",
			params: &types.CreateInput{},
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

func Test_CreateInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateInput
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.CreateInput{
				Name: "test-list-name",
			},
			expect: strings.NewReader(`{"name":"test-list-name"}`),
		},
		{
			name: "ok: has some parameters",
			params: &types.CreateInput{
				Name:        "test-list-name",
				Description: gotwi.String("test description"),
			},
			expect: strings.NewReader(`{"name":"test-list-name","description":"test description"}`),
		},
		{
			name: "ok: has some parameters",
			params: &types.CreateInput{
				Name:        "test-list-name",
				Description: gotwi.String("test description"),
				Private:     gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"name":"test-list-name","description":"test description","private":true}`),
		},
		{
			name:   "ok: has no required parameters",
			params: &types.CreateInput{},
			expect: strings.NewReader(`{"name":""}`),
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

func Test_CreateInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateInput
		expect map[string]string
	}{
		{
			name:   "normal: some parameters",
			params: &types.CreateInput{Name: "name"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.CreateInput{},
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

func Test_UpdateInput_SetAccessToken(t *testing.T) {
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
			p := &types.UpdateInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_UpdateInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.UpdateInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.UpdateInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.UpdateInput{},
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

func Test_UpdateInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UpdateInput
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.UpdateInput{
				ID: "test-id",
			},
			expect: strings.NewReader(`{}`),
		},
		{
			name: "ok: has some parameters",
			params: &types.UpdateInput{
				ID:          "test-id",
				Name:        gotwi.String("test-list-name"),
				Description: gotwi.String("test description"),
			},
			expect: strings.NewReader(`{"name":"test-list-name","description":"test description"}`),
		},
		{
			name: "ok: has some parameters",
			params: &types.UpdateInput{
				ID:          "test-id",
				Name:        gotwi.String("test-list-name"),
				Description: gotwi.String("test description"),
				Private:     gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"name":"test-list-name","description":"test description","private":true}`),
		},
		{
			name: "ok: has no required parameters (no effect)",
			params: &types.UpdateInput{
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

func Test_UpdateInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UpdateInput
		expect map[string]string
	}{
		{
			name:   "normal: some parameters",
			params: &types.UpdateInput{Name: gotwi.String("name")},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.UpdateInput{},
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

func Test_DeleteInput_SetAccessToken(t *testing.T) {
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
			p := &types.DeleteInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_DeleteInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.DeleteInput
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.DeleteInput{
				ID: "lid",
			},
			expect: endpointRoot + "lid",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.DeleteInput{},
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

func Test_DeleteInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.DeleteInput
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.DeleteInput{
				ID: "lid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.DeleteInput{},
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

func Test_DeleteInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.DeleteInput
		expect map[string]string
	}{
		{
			name:   "normal: some parameters",
			params: &types.DeleteInput{ID: "id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.DeleteInput{},
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
