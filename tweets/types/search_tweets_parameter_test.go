package types_test

import (
	"testing"
	"time"

	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_SearchTweetsRecent_SetAccessToken(t *testing.T) {
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
			p := &types.SearchTweetsRecentParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SearchTweetsRecent_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	startTime := time.Date(2021, 10, 24, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2021, 10, 24, 23, 59, 59, 59, time.UTC)

	cases := []struct {
		name   string
		params *types.SearchTweetsRecentParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SearchTweetsRecentParams{
				Query: "from:testuser",
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser",
		},
		{
			name: "with expansions",
			params: &types.SearchTweetsRecentParams{
				Query:      "from:testuser",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&max_results=10&query=from%3Atestuser",
		},
		{
			name: "with media.fields",
			params: &types.SearchTweetsRecentParams{
				Query:       "from:testuser",
				MediaFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&media.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with place.fields",
			params: &types.SearchTweetsRecentParams{
				Query:       "from:testuser",
				PlaceFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&place.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with poll.fields",
			params: &types.SearchTweetsRecentParams{
				Query:      "from:testuser",
				PollFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&poll.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with tweets.fields",
			params: &types.SearchTweetsRecentParams{
				Query:       "from:testuser",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.SearchTweetsRecentParams{
				Query:      "from:testuser",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&user.fields=uf1%2Cuf2",
		},
		{
			name: "with end_time",
			params: &types.SearchTweetsRecentParams{
				Query:   "from:testuser",
				EndTime: &endTime,
			},
			expect: endpointBase + "?end_time=2021-10-24T23%3A59%3A59Z&max_results=10&query=from%3Atestuser",
		},
		{
			name: "with start_time",
			params: &types.SearchTweetsRecentParams{
				Query:     "from:testuser",
				StartTime: &startTime,
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&start_time=2021-10-24T00%3A00%3A00Z",
		},
		{
			name: "with since_id",
			params: &types.SearchTweetsRecentParams{
				Query:   "from:testuser",
				SinceID: "sid",
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&since_id=sid",
		},
		{
			name: "with until_id",
			params: &types.SearchTweetsRecentParams{
				Query:   "from:testuser",
				UntilID: "uid",
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&until_id=uid",
		},
		{
			name: "all query parameters",
			params: &types.SearchTweetsRecentParams{
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
			params: &types.SearchTweetsRecentParams{
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

func Test_SearchTweetsRecent_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SearchTweetsRecentParams
	}{
		{
			name:   "empty params",
			params: &types.SearchTweetsRecentParams{},
		},
		{
			name:   "some params",
			params: &types.SearchTweetsRecentParams{Query: "from:testuser"},
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

func Test_SearchTweetsAll_SetAccessToken(t *testing.T) {
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
			p := &types.SearchTweetsAllParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SearchTweetsAll_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	startTime := time.Date(2021, 10, 24, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2021, 10, 24, 23, 59, 59, 59, time.UTC)

	cases := []struct {
		name   string
		params *types.SearchTweetsAllParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SearchTweetsAllParams{
				Query: "from:testuser",
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser",
		},
		{
			name: "with expansions",
			params: &types.SearchTweetsAllParams{
				Query:      "from:testuser",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&max_results=10&query=from%3Atestuser",
		},
		{
			name: "with media.fields",
			params: &types.SearchTweetsAllParams{
				Query:       "from:testuser",
				MediaFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&media.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with place.fields",
			params: &types.SearchTweetsAllParams{
				Query:       "from:testuser",
				PlaceFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&place.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with poll.fields",
			params: &types.SearchTweetsAllParams{
				Query:      "from:testuser",
				PollFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&poll.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with tweets.fields",
			params: &types.SearchTweetsAllParams{
				Query:       "from:testuser",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.SearchTweetsAllParams{
				Query:      "from:testuser",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&user.fields=uf1%2Cuf2",
		},
		{
			name: "with end_time",
			params: &types.SearchTweetsAllParams{
				Query:   "from:testuser",
				EndTime: &endTime,
			},
			expect: endpointBase + "?end_time=2021-10-24T23%3A59%3A59Z&max_results=10&query=from%3Atestuser",
		},
		{
			name: "with start_time",
			params: &types.SearchTweetsAllParams{
				Query:     "from:testuser",
				StartTime: &startTime,
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&start_time=2021-10-24T00%3A00%3A00Z",
		},
		{
			name: "with since_id",
			params: &types.SearchTweetsAllParams{
				Query:   "from:testuser",
				SinceID: "sid",
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&since_id=sid",
		},
		{
			name: "with until_id",
			params: &types.SearchTweetsAllParams{
				Query:   "from:testuser",
				UntilID: "uid",
			},
			expect: endpointBase + "?max_results=10&query=from%3Atestuser&until_id=uid",
		},
		{
			name: "all query parameters",
			params: &types.SearchTweetsAllParams{
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
			params: &types.SearchTweetsAllParams{
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

func Test_SearchTweetsAll_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SearchTweetsAllParams
	}{
		{
			name:   "empty params",
			params: &types.SearchTweetsAllParams{},
		},
		{
			name:   "some params",
			params: &types.SearchTweetsAllParams{Query: "from:testuser"},
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
