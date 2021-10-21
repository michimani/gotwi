package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetRetweetsRetweetedByParams_SetAccessToken(t *testing.T) {
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
			p := &types.TweetRetweetsRetweetedByParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_TweetRetweetsRetweetedByParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.TweetRetweetsRetweetedByParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.TweetRetweetsRetweetedByParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "with expansions",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:         "test-id",
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:          "test-id",
				MediaFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:          "test-id",
				PlaceFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:         "test-id",
				PollFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:          "test-id",
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:         "test-id",
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:          "test-id",
				Expansions:  []string{"ex"},
				MediaFields: []string{"mf"},
				PlaceFields: []string{"plf"},
				PollFields:  []string{"pof"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.TweetRetweetsRetweetedByParams{
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

func Test_TweetRetweetsRetweetedByParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetRetweetsRetweetedByParams
	}{
		{
			name:   "empty params",
			params: &types.TweetRetweetsRetweetedByParams{},
		},
		{
			name:   "some params",
			params: &types.TweetRetweetsRetweetedByParams{ID: "id"},
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

func Test_TweetRetweetsPostParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.TweetRetweetsPostParams
		expect string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.TweetRetweetsPostParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "normal: has no required parameter",
			params: &types.TweetRetweetsPostParams{
				TweetID: "tid",
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

func Test_TweetRetweetsPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetRetweetsPostParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.TweetRetweetsPostParams{
				ID:      "test-id",
				TweetID: "tid",
			},
			expect: strings.NewReader(`{"tweet_id":"tid"}`),
		},
		{
			name:   "ok: has no required parameters",
			params: &types.TweetRetweetsPostParams{ID: "id"},
			expect: strings.NewReader(`{"tweet_id":""}`),
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
