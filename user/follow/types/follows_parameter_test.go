package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/user/follow/types"
	"github.com/stretchr/testify/assert"
)

func Test_FollowsMaxResults_Valid(t *testing.T) {
	cases := []struct {
		name   string
		max    types.FollowsMaxResults
		expect bool
	}{
		{
			name:   "ok: 1",
			max:    types.FollowsMaxResults(1),
			expect: true,
		},
		{
			name:   "ok: 1000",
			max:    types.FollowsMaxResults(1000),
			expect: true,
		},
		{
			name:   "ng: 0",
			max:    types.FollowsMaxResults(0),
			expect: false,
		},
		{
			name:   "ng: 1001",
			max:    types.FollowsMaxResults(1001),
			expect: false,
		},
		{
			name:   "ng: -1",
			max:    types.FollowsMaxResults(-1),
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

func Test_FollowsMaxResults_String(t *testing.T) {
	cases := []struct {
		name   string
		max    types.FollowsMaxResults
		expect string
	}{
		{
			name:   "ok: 1",
			max:    types.FollowsMaxResults(1),
			expect: "1",
		},
		{
			name:   "ok: 0",
			max:    types.FollowsMaxResults(0),
			expect: "0",
		},
		{
			name:   "ok: -1",
			max:    types.FollowsMaxResults(-1),
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

func Test_ListFollowingsInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListFollowingsInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListFollowingsInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.ListFollowingsInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListFollowingsInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: with specific max_result",
			params: &types.ListFollowingsInput{
				ID:         "test-id",
				MaxResults: 111,
			},
			expect: endpointRoot + "test-id?max_results=111",
		},
		{
			name: "normal: with pagination_token",
			params: &types.ListFollowingsInput{
				ID:              "test-id",
				PaginationToken: "p-token",
			},
			expect: endpointRoot + "test-id?pagination_token=p-token",
		},
		{
			name: "normal: with expansions",
			params: &types.ListFollowingsInput{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.ListFollowingsInput{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.ListFollowingsInput{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.ListFollowingsInput{
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
			params: &types.ListFollowingsInput{
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

func Test_ListFollowingsInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListFollowingsInput
	}{
		{
			name:   "empty params",
			params: &types.ListFollowingsInput{},
		},
		{
			name:   "some params",
			params: &types.ListFollowingsInput{ID: "id"},
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

func Test_ListFollowersInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListFollowersInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListFollowersInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.ListFollowersInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.ListFollowersInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: with specific max_result",
			params: &types.ListFollowersInput{
				ID:         "test-id",
				MaxResults: 111,
			},
			expect: endpointRoot + "test-id?max_results=111",
		},
		{
			name: "normal: with pagination_token",
			params: &types.ListFollowersInput{
				ID:              "test-id",
				PaginationToken: "p-token",
			},
			expect: endpointRoot + "test-id?pagination_token=p-token",
		},
		{
			name: "normal: with expansions",
			params: &types.ListFollowersInput{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id?expansions=ex1%2Cex2",
		},
		{
			name: "normal: with tweets.fields",
			params: &types.ListFollowersInput{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "normal: with users.fields",
			params: &types.ListFollowersInput{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id?user.fields=uf1%2Cuf2",
		},
		{
			name: "normal: all query parameters",
			params: &types.ListFollowersInput{
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
			params: &types.ListFollowersInput{
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

func Test_ListFollowersInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListFollowersInput
	}{
		{
			name:   "empty params",
			params: &types.ListFollowersInput{},
		},
		{
			name:   "some params",
			params: &types.ListFollowersInput{ID: "id"},
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

func Test_CreateFollowingInput_SetAccessToken(t *testing.T) {
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
			p := &types.CreateFollowingInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_CreateFollowingInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.CreateFollowingInput
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.CreateFollowingInput{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: has no required parameter",
			params: &types.CreateFollowingInput{
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

func Test_CreateFollowingInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateFollowingInput
		expect io.Reader
	}{
		{
			name: "ok: has both of path and json parameters",
			params: &types.CreateFollowingInput{
				ID:       "test-id",
				TargetID: "tid",
			},
			expect: strings.NewReader(`{"target_user_id":"tid"}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.CreateFollowingInput{ID: "id"},
			expect: strings.NewReader(`{"target_user_id":""}`),
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

func Test_CreateFollowingInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateFollowingInput
		expect map[string]string
	}{
		{
			name:   "normal: has both of path and json parameters",
			params: &types.CreateFollowingInput{ID: "id", TargetID: "tid"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.CreateFollowingInput{},
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

func Test_DeleteFollowingInput_SetAccessToken(t *testing.T) {
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
			p := &types.DeleteFollowingInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_DeleteFollowingInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:source_user_id/:target_user_id"
	cases := []struct {
		name   string
		params *types.DeleteFollowingInput
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.DeleteFollowingInput{
				SourceUserID: "sid",
				TargetID:     "tid",
			},
			expect: endpointRoot + "sid" + "/" + "tid",
		},
		{
			name: "normal: has no required parameter",
			params: &types.DeleteFollowingInput{
				SourceUserID: "sid",
			},
			expect: "",
		},
		{
			name: "normal: has no required parameter",
			params: &types.DeleteFollowingInput{
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

func Test_DeleteFollowingInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.DeleteFollowingInput
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.DeleteFollowingInput{
				SourceUserID: "sid",
				TargetID:     "tid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.DeleteFollowingInput{},
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

func Test_DeleteFollowingInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.DeleteFollowingInput
		expect map[string]string
	}{
		{
			name:   "normal: has parameters",
			params: &types.DeleteFollowingInput{SourceUserID: "id", TargetID: "tid"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.DeleteFollowingInput{},
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
