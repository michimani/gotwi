package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi/tweet/hidereply/types"
	"github.com/stretchr/testify/assert"
)

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
			name:   "normal: empty",
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
			name: "normal: has no required parameter",
			params: &types.UpdateInput{
				Hidden: true,
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

func Test_UpdateInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.UpdateInput
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters (true)",
			params: &types.UpdateInput{
				ID:     "test-id",
				Hidden: true,
			},
			expect: strings.NewReader(`{"hidden":true}`),
		},
		{
			name: "ok: has both of path and json parameters (false)",
			params: &types.UpdateInput{
				ID:     "test-id",
				Hidden: false,
			},
			expect: strings.NewReader(`{"hidden":false}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.UpdateInput{ID: "id"},
			expect: strings.NewReader(`{"hidden":false}`),
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
			name:   "normal: has both of path and json parameters",
			params: &types.UpdateInput{ID: "id", Hidden: true},
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
