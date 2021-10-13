package types_test

import (
	"testing"

	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_SearchTweetsTweetsSearchRecent_SetAccessToken(t *testing.T) {
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
			p := &types.SearchTweetsTweetsSearchRecentParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SearchTweetsTweetsSearchRecent_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.SearchTweetsTweetsSearchRecentParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SearchTweetsTweetsSearchRecentParams{
				Query: "from:testuser",
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser",
		},
		{
			name: "with expansions",
			params: &types.SearchTweetsTweetsSearchRecentParams{
				Query:      "from:testuser",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&max_results=10&query=from%3Atestuser",
		},
		{
			name: "with media.fields",
			params: &types.SearchTweetsTweetsSearchRecentParams{
				Query:       "from:testuser",
				MediaFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&media.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with place.fields",
			params: &types.SearchTweetsTweetsSearchRecentParams{
				Query:       "from:testuser",
				PlaceFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&place.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with poll.fields",
			params: &types.SearchTweetsTweetsSearchRecentParams{
				Query:      "from:testuser",
				PollFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&poll.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with tweets.fields",
			params: &types.SearchTweetsTweetsSearchRecentParams{
				Query:       "from:testuser",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.SearchTweetsTweetsSearchRecentParams{
				Query:      "from:testuser",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SearchTweetsTweetsSearchRecentParams{
				Query:       "from:testuser",
				Expansions:  []string{"ex"},
				MediaFields: []string{"mf"},
				PlaceFields: []string{"plf"},
				PollFields:  []string{"pof"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointBase + "?expansions=ex&max_results=10&media.fields=mf&place.fields=plf&poll.fields=pof&query=from%3Atestuser&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.SearchTweetsTweetsSearchRecentParams{
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
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

func Test_SearchTweetsTweetsSearchRecent_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SearchTweetsTweetsSearchRecentParams
	}{
		{
			name:   "empty params",
			params: &types.SearchTweetsTweetsSearchRecentParams{},
		},
		{
			name:   "some params",
			params: &types.SearchTweetsTweetsSearchRecentParams{Query: "from:testuser"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}

func Test_SearchTweetsTweetsSearchAll_SetAccessToken(t *testing.T) {
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
			p := &types.SearchTweetsTweetsSearchAllParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SearchTweetsTweetsSearchAll_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.SearchTweetsTweetsSearchAllParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SearchTweetsTweetsSearchAllParams{
				Query: "from:testuser",
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser",
		},
		{
			name: "with expansions",
			params: &types.SearchTweetsTweetsSearchAllParams{
				Query:      "from:testuser",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&max_results=10&query=from%3Atestuser",
		},
		{
			name: "with media.fields",
			params: &types.SearchTweetsTweetsSearchAllParams{
				Query:       "from:testuser",
				MediaFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&media.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with place.fields",
			params: &types.SearchTweetsTweetsSearchAllParams{
				Query:       "from:testuser",
				PlaceFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&place.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with poll.fields",
			params: &types.SearchTweetsTweetsSearchAllParams{
				Query:      "from:testuser",
				PollFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&poll.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with tweets.fields",
			params: &types.SearchTweetsTweetsSearchAllParams{
				Query:       "from:testuser",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.SearchTweetsTweetsSearchAllParams{
				Query:      "from:testuser",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SearchTweetsTweetsSearchAllParams{
				Query:       "from:testuser",
				Expansions:  []string{"ex"},
				MediaFields: []string{"mf"},
				PlaceFields: []string{"plf"},
				PollFields:  []string{"pof"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointBase + "?expansions=ex&max_results=10&media.fields=mf&place.fields=plf&poll.fields=pof&query=from%3Atestuser&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.SearchTweetsTweetsSearchAllParams{
				Expansions:  []string{"ex"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
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

func Test_SearchTweetsTweetsSearchAll_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SearchTweetsTweetsSearchAllParams
	}{
		{
			name:   "empty params",
			params: &types.SearchTweetsTweetsSearchAllParams{},
		},
		{
			name:   "some params",
			params: &types.SearchTweetsTweetsSearchAllParams{Query: "from:testuser"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}
