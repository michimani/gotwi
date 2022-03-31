package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_QuoteTweetsParams_SetAccessToken(t *testing.T) {
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
			p := &types.QuoteTweetsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_QuoteTweets_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.QuoteTweetsParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.QuoteTweetsParams{
				ID: "tid",
			},
			expect: endpointRoot + "tid",
		},
		{
			name: "with expansions",
			params: &types.QuoteTweetsParams{
				ID:         "tid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "tid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.QuoteTweetsParams{
				ID:          "tid",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "tid" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.QuoteTweetsParams{
				ID:          "tid",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "tid" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.QuoteTweetsParams{
				ID:         "tid",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "tid" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.QuoteTweetsParams{
				ID:          "tid",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "tid" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.QuoteTweetsParams{
				ID:         "tid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "tid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.QuoteTweetsParams{
				ID:              "tid",
				Expansions:      fields.ExpansionList{"ex"},
				MediaFields:     fields.MediaFieldList{"mf"},
				PlaceFields:     fields.PlaceFieldList{"plf"},
				PollFields:      fields.PollFieldList{"pof"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
				MaxResults:      10,
				PaginationToken: "token",
			},
			expect: endpointRoot + "tid" + "?expansions=ex&max_results=10&media.fields=mf&pagination_token=token&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.QuoteTweetsParams{
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

func Test_QuoteTweets_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.QuoteTweetsParams
	}{
		{
			name:   "empty params",
			params: &types.QuoteTweetsParams{},
		},
		{
			name:   "some params",
			params: &types.QuoteTweetsParams{ID: "tid"},
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
