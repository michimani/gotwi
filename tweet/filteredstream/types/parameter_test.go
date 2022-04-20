package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweet/filteredstream/types"
	"github.com/stretchr/testify/assert"
)

func Test_FilteredStreamRulesGet_SetAccessToken(t *testing.T) {
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
			p := &types.ListRulesInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_FilteredStreamRulesGet_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"

	cases := []struct {
		name   string
		params *types.ListRulesInput
		expect string
	}{
		{
			name:   "has no parameter",
			params: &types.ListRulesInput{},
			expect: endpointBase,
		},
		{
			name: "with ids",
			params: &types.ListRulesInput{
				IDs: []string{"rid1", "rid2"},
			},
			expect: endpointBase + "?ids=rid1%2Crid2",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_FilteredStreamRulesGet_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListRulesInput
	}{
		{
			name:   "empty params",
			params: &types.ListRulesInput{},
		},
		{
			name:   "some params",
			params: &types.ListRulesInput{IDs: []string{"rid1", "rid2"}},
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

func Test_CreateRulesInput_SetAccessToken(t *testing.T) {
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
			p := &types.CreateRulesInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_CreateRulesInput_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.CreateRulesInput
		expect string
	}{
		{
			name: "ok: only required parameter (add)",
			params: &types.CreateRulesInput{
				Add: []types.AddingRule{
					{Value: gotwi.String("test")},
				},
			},
			expect: endpoint + "?dry_run=false",
		},
		{
			name: "ok: with dry run (true)",
			params: &types.CreateRulesInput{
				Add: []types.AddingRule{
					{Value: gotwi.String("test")},
				},
				DryRun: true,
			},
			expect: endpoint + "?dry_run=true",
		},
		{
			name: "ok: with dry run (false)",
			params: &types.CreateRulesInput{
				Add: []types.AddingRule{
					{Value: gotwi.String("test")},
				},
				DryRun: false,
			},
			expect: endpoint + "?dry_run=false",
		},
		{
			name: "ng: has no required parameters",
			params: &types.CreateRulesInput{
				DryRun: true,
			},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpoint)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_CreateRulesInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateRulesInput
		expect io.Reader
	}{
		{
			name: "ok: has required parameter (add)",
			params: &types.CreateRulesInput{
				Add: []types.AddingRule{
					{Value: gotwi.String("test-value-1"), Tag: gotwi.String("test-tag-1")},
					{Value: gotwi.String("test-value-2")},
				},
			},
			expect: strings.NewReader(`{"add":[{"value":"test-value-1","tag":"test-tag-1"},{"value":"test-value-2"}]}`),
		},
		{
			name: "ng: has no required parameters",
			params: &types.CreateRulesInput{
				DryRun: true,
			},
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

func Test_CreateRulesInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateRulesInput
		expect map[string]string
	}{
		{
			name:   "ok: has dry_run (true)",
			params: &types.CreateRulesInput{DryRun: true},
			expect: map[string]string{
				"dry_run": "true",
			},
		},
		{
			name:   "ok: has dry_run (false)",
			params: &types.CreateRulesInput{DryRun: false},
			expect: map[string]string{
				"dry_run": "false",
			},
		},
		{
			name:   "ok: has no query parameter (default)",
			params: &types.CreateRulesInput{},
			expect: map[string]string{
				"dry_run": "false",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			m := c.params.ParameterMap()
			assert.Equal(tt, c.expect, m)
		})
	}
}

func Test_DeleteRulesInput_SetAccessToken(t *testing.T) {
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
			p := &types.DeleteRulesInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_DeleteRulesInput_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.DeleteRulesInput
		expect string
	}{
		{
			name: "ok: only required parameter (delete)",
			params: &types.DeleteRulesInput{
				Delete: &types.DeletingRules{IDs: []string{"test-id"}},
			},
			expect: endpoint + "?dry_run=false",
		},
		{
			name: "ok: with dry run (true)",
			params: &types.DeleteRulesInput{
				Delete: &types.DeletingRules{IDs: []string{"test-id"}},
				DryRun: true,
			},
			expect: endpoint + "?dry_run=true",
		},
		{
			name: "ok: with dry run (false)",
			params: &types.DeleteRulesInput{
				Delete: &types.DeletingRules{IDs: []string{"test-id"}},
				DryRun: false,
			},
			expect: endpoint + "?dry_run=false",
		},
		{
			name: "ng: has no required parameters",
			params: &types.DeleteRulesInput{
				DryRun: true,
			},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpoint)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_DeleteRulesInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.DeleteRulesInput
		expect io.Reader
	}{
		{
			name: "ok: has required parameter (delete)",
			params: &types.DeleteRulesInput{
				Delete: &types.DeletingRules{IDs: []string{"test-id-1", "test-id-2"}},
			},
			expect: strings.NewReader(`{"delete":{"ids":["test-id-1","test-id-2"]}}`),
		},
		{
			name: "ng: has no required parameters",
			params: &types.DeleteRulesInput{
				DryRun: true,
			},
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

func Test_DeleteRulesInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.DeleteRulesInput
		expect map[string]string
	}{
		{
			name:   "ok: has dry_run (true)",
			params: &types.DeleteRulesInput{DryRun: true},
			expect: map[string]string{
				"dry_run": "true",
			},
		},
		{
			name:   "ok: has dry_run (false)",
			params: &types.DeleteRulesInput{DryRun: false},
			expect: map[string]string{
				"dry_run": "false",
			},
		},
		{
			name:   "ok: has no query parameter (default)",
			params: &types.DeleteRulesInput{},
			expect: map[string]string{
				"dry_run": "false",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			m := c.params.ParameterMap()
			assert.Equal(tt, c.expect, m)
		})
	}
}

func Test_SearchStreamInput_SetAccessToken(t *testing.T) {
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
			p := &types.SearchStreamInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SearchStreamInput_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.SearchStreamInput
		expect string
	}{
		{
			name:   "ok",
			params: &types.SearchStreamInput{},
			expect: endpoint,
		},
		{
			name: "with valid backfill_minutes",
			params: &types.SearchStreamInput{
				BackfillMinutes: 3,
			},
			expect: endpoint + "?backfill_minutes=3",
		},
		{
			name: "with invalid backfill_minutes",
			params: &types.SearchStreamInput{
				BackfillMinutes: -1,
			},
			expect: endpoint,
		},
		{
			name: "with expansions",
			params: &types.SearchStreamInput{
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpoint + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.SearchStreamInput{
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.SearchStreamInput{
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.SearchStreamInput{
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.SearchStreamInput{
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.SearchStreamInput{
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpoint + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SearchStreamInput{
				BackfillMinutes: 3,
				Expansions:      fields.ExpansionList{"ex"},
				MediaFields:     fields.MediaFieldList{"mf"},
				PlaceFields:     fields.PlaceFieldList{"plf"},
				PollFields:      fields.PollFieldList{"pof"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
			},
			expect: endpoint + "?backfill_minutes=3&expansions=ex&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name:   "nil",
			params: nil,
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpoint)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_SearchStreamInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SearchStreamInput
	}{
		{
			name:   "empty params",
			params: &types.SearchStreamInput{},
		},
		{
			name:   "some params",
			params: &types.SearchStreamInput{Expansions: fields.ExpansionList{"ex"}},
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
