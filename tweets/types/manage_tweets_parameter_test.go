package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

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
