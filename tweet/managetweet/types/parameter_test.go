package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet/types"
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
			name:   "normal: empty",
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
			name:   "normal: some parameter",
			params: &types.CreateInput{},
			expect: endpointBase,
		},
		{
			name:   "normal: has no parameters",
			params: &types.CreateInput{Text: gotwi.String("test")},
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

func Test_CreateInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateInput
		expect io.Reader
	}{
		{
			name: "ok: has some json parameters",
			params: &types.CreateInput{
				Text: gotwi.String("test text"),
				Poll: &types.CreateInputPoll{
					DurationMinutes: gotwi.Int(5),
					Options:         []string{"op1", "op2"},
				},
				QuoteTweetID: gotwi.String("quote_tweet_id"),
			},
			expect: strings.NewReader(`{"poll":{"duration_minutes":5,"options":["op1","op2"]},"quote_tweet_id":"quote_tweet_id","text":"test text"}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.CreateInput{},
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

func Test_CreateInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateInput
		expect map[string]string
	}{
		{
			name:   "normal: has both of path and json parameters",
			params: &types.CreateInput{Text: gotwi.String("test")},
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
			name:   "normal: empty",
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
			name:   "normal: only required parameter",
			params: &types.DeleteInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
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
			name: "ok: has path parameters",
			params: &types.DeleteInput{
				ID: "test-id",
			},
			expect: nil,
		},
		{
			name:   "ok: has no json parameters",
			params: &types.DeleteInput{ID: "id"},
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
			name:   "normal: has both of path and json parameters",
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
