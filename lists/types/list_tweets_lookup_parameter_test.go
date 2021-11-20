package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/lists/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListTweetsLookup_SetAccessToken(t *testing.T) {
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
			p := &types.ListTweetsLookupParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListTweetsLookup_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListTweetsLookupParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListTweetsLookupParams{
				ID: "lid",
			},
			expect: endpointRoot + "lid",
		},
		{
			name: "with expansions",
			params: &types.ListTweetsLookupParams{
				ID:         "lid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "lid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with tweets.fields",
			params: &types.ListTweetsLookupParams{
				ID:          "lid",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "lid" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.ListTweetsLookupParams{
				ID:         "lid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "lid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "with max_results",
			params: &types.ListTweetsLookupParams{
				ID:         "lid",
				MaxResults: 10,
			},
			expect: endpointRoot + "lid" + "?max_results=10",
		},
		{
			name: "with pagination_token",
			params: &types.ListTweetsLookupParams{
				ID:              "lid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "lid" + "?pagination_token=ptoken",
		},
		{
			name: "all query parameters",
			params: &types.ListTweetsLookupParams{
				Expansions:      fields.ExpansionList{"ex"},
				ID:              "lid",
				TweetFields:     fields.TweetFieldList{"tf"},
				MaxResults:      10,
				PaginationToken: "ptoken",
				UserFields:      fields.UserFieldList{"uf"},
			},
			expect: endpointRoot + "lid" + "?expansions=ex&max_results=10&pagination_token=ptoken&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListTweetsLookupParams{
				Expansions:      fields.ExpansionList{"ex"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
				MaxResults:      10,
				PaginationToken: "pagination_token",
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

func Test_ListTweetsLookup_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListTweetsLookupParams
	}{
		{
			name:   "empty params",
			params: &types.ListTweetsLookupParams{},
		},
		{
			name:   "some params",
			params: &types.ListTweetsLookupParams{ID: "sid"},
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
