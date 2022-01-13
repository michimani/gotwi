package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_ManageTweetsPostParams_SetAccessToken(t *testing.T) {
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
			p := &types.ManageTweetsPostParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ManageTweetsPostParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.ManageTweetsPostParams
		expect string
	}{
		{
			name:   "normal: some parameter",
			params: &types.ManageTweetsPostParams{},
			expect: endpointBase,
		},
		{
			name:   "normal: has no parameters",
			params: &types.ManageTweetsPostParams{Text: gotwi.String("test")},
			expect: endpointBase,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_ManageTweetsPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageTweetsPostParams
		expect io.Reader
	}{
		{
			name: "ok: has some json parameters",
			params: &types.ManageTweetsPostParams{
				Text: gotwi.String("test text"),
				Poll: &types.ManageTweetsPostParamsPoll{
					DurationMinutes: gotwi.Int(5),
					Options:         []string{"op1", "op2"},
				},
			},
			expect: strings.NewReader(`{"poll":{"duration_minutes":5,"options":["op1","op2"]},"text":"test text"}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.ManageTweetsPostParams{},
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

func Test_ManageTweetsPostParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageTweetsPostParams
		expect map[string]string
	}{
		{
			name:   "normal: has both of path and json parameters",
			params: &types.ManageTweetsPostParams{Text: gotwi.String("test")},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.ManageTweetsPostParams{},
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

func Test_ManageTweetsDeleteParams_SetAccessToken(t *testing.T) {
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
			p := &types.ManageTweetsDeleteParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ManageTweetsDeleteParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.ManageTweetsDeleteParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ManageTweetsDeleteParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name:   "normal: has no required parameter",
			params: &types.ManageTweetsDeleteParams{},
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

func Test_ManageTweetsDeleteParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageTweetsDeleteParams
		expect io.Reader
	}{
		{
			name: "ok: has path parameters",
			params: &types.ManageTweetsDeleteParams{
				ID: "test-id",
			},
			expect: nil,
		},
		{
			name:   "ok: has no json parameters",
			params: &types.ManageTweetsDeleteParams{ID: "id"},
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

func Test_ManageTweetsDeleteParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ManageTweetsDeleteParams
		expect map[string]string
	}{
		{
			name:   "normal: has both of path and json parameters",
			params: &types.ManageTweetsDeleteParams{ID: "id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.ManageTweetsDeleteParams{},
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
