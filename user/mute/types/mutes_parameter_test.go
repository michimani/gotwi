package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/user/mute/types"
	"github.com/stretchr/testify/assert"
)

func Test_MutesMaxResults_Valid(t *testing.T) {
	cases := []struct {
		name   string
		max    types.MutesMaxResults
		expect bool
	}{
		{
			name:   "ok: 1",
			max:    types.MutesMaxResults(1),
			expect: true,
		},
		{
			name:   "ok: 1000",
			max:    types.MutesMaxResults(1000),
			expect: true,
		},
		{
			name:   "ng: 0",
			max:    types.MutesMaxResults(0),
			expect: false,
		},
		{
			name:   "ng: 1001",
			max:    types.MutesMaxResults(1001),
			expect: false,
		},
		{
			name:   "ng: -1",
			max:    types.MutesMaxResults(-1),
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

func Test_MutesMaxResults_String(t *testing.T) {
	cases := []struct {
		name   string
		max    types.MutesMaxResults
		expect string
	}{
		{
			name:   "ok: 1",
			max:    types.MutesMaxResults(1),
			expect: "1",
		},
		{
			name:   "ok: 0",
			max:    types.MutesMaxResults(0),
			expect: "0",
		},
		{
			name:   "ok: -1",
			max:    types.MutesMaxResults(-1),
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

func Test_ListsInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListsInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListsInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.ListsInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListsInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: with specific max_result",
			params: &types.ListsInput{
				ID:         "test-id",
				MaxResults: 111,
			},
			expect: endpointRoot + "test-id?max_results=111",
		},
		{
			name: "normal: with pagination_token",
			params: &types.ListsInput{
				ID:              "test-id",
				PaginationToken: "p-token",
			},
			expect: endpointRoot + "test-id?pagination_token=p-token",
		},
		{
			name: "normal: with expansions",
			params: &types.ListsInput{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.ListsInput{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.ListsInput{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.ListsInput{
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
			params: &types.ListsInput{
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

func Test_ListsInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListsInput
	}{
		{
			name:   "empty params",
			params: &types.ListsInput{},
		},
		{
			name:   "some params",
			params: &types.ListsInput{ID: "id"},
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
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.CreateInput
		expect string
	}{
		{
			name: "ok",
			params: &types.CreateInput{
				ID: "test-id", TargetID: "tid",
			},
			expect: endpointRoot + "test-id",
		},
		{
			name: "ok: has no required parameter (ID)",
			params: &types.CreateInput{
				TargetID: "tid",
			},
			expect: "",
		},
		{
			name: "ok: has no required parameter (TargetID)",
			params: &types.CreateInput{
				ID: "tid",
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

func Test_CreateInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateInput
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.CreateInput{
				ID:       "test-id",
				TargetID: "tid",
			},
			expect: strings.NewReader(`{"target_user_id":"tid"}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.CreateInput{ID: "id"},
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
			params: &types.CreateInput{ID: "id", TargetID: "tid"},
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
	const endpointBase = "test/endpoint/:source_user_id/:target_user_id"
	cases := []struct {
		name   string
		params *types.DeleteInput
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.DeleteInput{
				SourceUserID: "sid",
				TargetID:     "tid",
			},
			expect: endpointRoot + "sid" + "/" + "tid",
		},
		{
			name: "normal: has no required parameter",
			params: &types.DeleteInput{
				SourceUserID: "sid",
			},
			expect: "",
		},
		{
			name: "normal: has no required parameter",
			params: &types.DeleteInput{
				TargetID: "tid",
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

func Test_DeleteInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.DeleteInput
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.DeleteInput{
				SourceUserID: "sid",
				TargetID:     "tid",
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
			name:   "normal: has parameters",
			params: &types.DeleteInput{SourceUserID: "id", TargetID: "tid"},
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
