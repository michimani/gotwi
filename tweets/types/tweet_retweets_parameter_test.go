package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
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
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:          "test-id",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:          "test-id",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:         "test-id",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.TweetRetweetsRetweetedByParams{
				ID:          "test-id",
				Expansions:  fields.ExpansionList{"ex"},
				MediaFields: fields.MediaFieldList{"mf"},
				PlaceFields: fields.PlaceFieldList{"plf"},
				PollFields:  fields.PollFieldList{"pof"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.TweetRetweetsRetweetedByParams{
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

func Test_TweetRetweetsPostParams_SetAccessToken(t *testing.T) {
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
			p := &types.TweetRetweetsPostParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
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
				TweetID: gotwi.String("tid"),
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
			name: "ok: has both ob path and json parameters",
			params: &types.TweetRetweetsPostParams{
				ID:      "test-id",
				TweetID: gotwi.String("tid"),
			},
			expect: strings.NewReader(`{"tweet_id":"tid"}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.TweetRetweetsPostParams{ID: "id"},
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

func Test_TweetRetweetsPostParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetRetweetsPostParams
		expect map[string]string
	}{
		{
			name:   "normal: has both of path and json parameters",
			params: &types.TweetRetweetsPostParams{ID: "id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.TweetRetweetsPostParams{},
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

func Test_TweetRetweetsDeleteParams_SetAccessToken(t *testing.T) {
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
			p := &types.TweetRetweetsDeleteParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_TweetRetweetsDeleteParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id/:source_tweet_id"
	cases := []struct {
		name   string
		params *types.TweetRetweetsDeleteParams
		expect string
	}{
		{
			name: "normal: only required parameter",
			params: &types.TweetRetweetsDeleteParams{
				ID:            "uid",
				SourceTweetID: "sid",
			},
			expect: endpointRoot + "uid" + "/" + "sid",
		},
		{
			name: "normal: has no required parameter",
			params: &types.TweetRetweetsDeleteParams{
				ID: "uid",
			},
			expect: "",
		},
		{
			name: "normal: has no required parameter",
			params: &types.TweetRetweetsDeleteParams{
				SourceTweetID: "sid",
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

func Test_TweetRetweetsDeleteParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetRetweetsDeleteParams
		expect io.Reader
	}{
		{
			name: "ok: has required parameters",
			params: &types.TweetRetweetsDeleteParams{
				ID:            "uid",
				SourceTweetID: "sid",
			},
			expect: nil,
		},
		{
			name:   "ok: has no required parameters",
			params: &types.TweetRetweetsDeleteParams{},
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

func Test_TweetRetweetsDeleteParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetRetweetsDeleteParams
		expect map[string]string
	}{
		{
			name:   "normal: has both of path and json parameters",
			params: &types.TweetRetweetsDeleteParams{ID: "id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameter",
			params: &types.TweetRetweetsDeleteParams{},
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
