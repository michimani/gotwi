package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/users/types"
	"github.com/stretchr/testify/assert"
)

func Test_BlocksMaxResults_Valid(t *testing.T) {
	cases := []struct {
		name   string
		max    types.BlocksMaxResults
		expect bool
	}{
		{
			name:   "ok: 1",
			max:    types.BlocksMaxResults(1),
			expect: true,
		},
		{
			name:   "ok: 1000",
			max:    types.BlocksMaxResults(1000),
			expect: true,
		},
		{
			name:   "ng: 0",
			max:    types.BlocksMaxResults(0),
			expect: false,
		},
		{
			name:   "ng: 1001",
			max:    types.BlocksMaxResults(1001),
			expect: false,
		},
		{
			name:   "ng: -1",
			max:    types.BlocksMaxResults(-1),
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			b := c.max.Valid()
			assert.Equal(tt, c.expect, b)
		})
	}
}

func Test_BlocksMaxResults_String(t *testing.T) {
	cases := []struct {
		name   string
		max    types.BlocksMaxResults
		expect string
	}{
		{
			name:   "ok: 1",
			max:    types.BlocksMaxResults(1),
			expect: "1",
		},
		{
			name:   "ok: 0",
			max:    types.BlocksMaxResults(0),
			expect: "0",
		},
		{
			name:   "ok: -1",
			max:    types.BlocksMaxResults(-1),
			expect: "-1",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			s := c.max.String()
			assert.Equal(tt, c.expect, s)
		})
	}
}

func Test_BlocksBlockingGetParams_SetAccessToken(t *testing.T) {
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
			p := &types.BlocksBlockingGetParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_BlocksBlockingGetParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.BlocksBlockingGetParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.BlocksBlockingGetParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: with specific max_result",
			params: &types.BlocksBlockingGetParams{
				ID:         "test-id",
				MaxResults: 111,
			},
			expect: endpointRoot + "test-id?max_results=111",
		},
		{
			name: "normal: with pagination_token",
			params: &types.BlocksBlockingGetParams{
				ID:              "test-id",
				PaginationToken: "p-token",
			},
			expect: endpointRoot + "test-id?pagination_token=p-token",
		},
		{
			name: "normal: with expansions",
			params: &types.BlocksBlockingGetParams{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.BlocksBlockingGetParams{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.BlocksBlockingGetParams{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.BlocksBlockingGetParams{
				ID:              "test-id",
				MaxResults:      111,
				PaginationToken: "p-token",
				Expansions:      fields.ExpansionList{"ex"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "test-id?expansions=ex&max_results=111&pagination_token=p-token&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "normal: has no required parameter",
			params: &types.BlocksBlockingGetParams{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
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

func Test_BlocksBlockingGetParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BlocksBlockingGetParams
	}{
		{
			name:   "empty params",
			params: &types.BlocksBlockingGetParams{},
		},
		{
			name:   "some params",
			params: &types.BlocksBlockingGetParams{ID: "id"},
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

func Test_BlocksBlockingPostParams_SetAccessToken(t *testing.T) {
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
			p := &types.BlocksBlockingPostParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_BlocksBlockingPostParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.BlocksBlockingPostParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.BlocksBlockingPostParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: has no required parameter",
			params: &types.BlocksBlockingPostParams{
				TargetUserID: gotwi.String("tid"),
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

func Test_BlocksBlockingPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BlocksBlockingPostParams
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.BlocksBlockingPostParams{
				ID:           "test-id",
				TargetUserID: gotwi.String("tid"),
			},
			expect: strings.NewReader(`{"target_user_id":"tid"}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.BlocksBlockingPostParams{ID: "id"},
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

func Test_BlocksBlockingPostParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BlocksBlockingPostParams
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.BlocksBlockingPostParams{ID: "test-id"},
			expect: map[string]string{},
		},
		{
			name: "normal: has no required parameter",
			params: &types.BlocksBlockingPostParams{
				ID:           "tid",
				TargetUserID: gotwi.String("tid"),
			},
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

func Test_BlocksBlockingDeleteParams_SetAccessToken(t *testing.T) {
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
			p := &types.BlocksBlockingDeleteParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_BlocksBlockingDeleteParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:source_user_id/:target_user_id"
	cases := []struct {
		name   string
		params *types.BlocksBlockingDeleteParams
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.BlocksBlockingDeleteParams{
				SourceUserID: "sid",
				TargetUserID: "tid",
			},
			expect: endpointRoot + "sid" + "/" + "tid",
		},
		{
			name: "normal: has no required parameter",
			params: &types.BlocksBlockingDeleteParams{
				SourceUserID: "sid",
			},
			expect: "",
		},
		{
			name: "normal: has no required parameter",
			params: &types.BlocksBlockingDeleteParams{
				TargetUserID: "tid",
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

func Test_BlocksBlockingDeleteParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BlocksBlockingDeleteParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.BlocksBlockingDeleteParams{
				SourceUserID: "sid",
				TargetUserID: "tid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.BlocksBlockingDeleteParams{},
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

func Test_BlocksBlockingDeleteParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BlocksBlockingDeleteParams
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.BlocksBlockingDeleteParams{SourceUserID: "sid", TargetUserID: "tid"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no required parameter",
			params: &types.BlocksBlockingDeleteParams{},
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
