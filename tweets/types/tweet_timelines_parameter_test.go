package types_test

import (
	"testing"
	"time"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetTimelinesTweetsParams_SetAccessToken(t *testing.T) {
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
			p := &types.TweetTimelinesTweetsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_TweetTimelinesTweetsParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	startTime := time.Unix(1640499309, 0).UTC()
	endTime := time.Unix(1640585709, 0).UTC()
	cases := []struct {
		name   string
		params *types.TweetTimelinesTweetsParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.TweetTimelinesTweetsParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "with exclude",
			params: &types.TweetTimelinesTweetsParams{
				ID:      "test-id",
				Exclude: fields.ExcludeList{"exc1", "exc2"},
			},
			expect: endpointRoot + "test-id" + "?exclude=exc1%2Cexc2",
		},
		{
			name: "with expansions",
			params: &types.TweetTimelinesTweetsParams{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with max_results",
			params: &types.TweetTimelinesTweetsParams{
				ID:         "test-id",
				MaxResults: 50,
			},
			expect: endpointRoot + "test-id" + "?max_results=50",
		},
		{
			name: "with media.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:          "test-id",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:          "test-id",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:         "test-id",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.TweetTimelinesTweetsParams{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.TweetTimelinesTweetsParams{
				ID:              "test-id",
				Exclude:         fields.ExcludeList{"exc"},
				Expansions:      fields.ExpansionList{"ex"},
				MaxResults:      50,
				MediaFields:     fields.MediaFieldList{"mf"},
				PlaceFields:     fields.PlaceFieldList{"plf"},
				PollFields:      fields.PollFieldList{"pof"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
				StartTime:       &startTime,
				EndTime:         &endTime,
				SinceID:         "sid",
				UntilID:         "uid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "test-id" + "?end_time=2021-12-27T06%3A15%3A09Z&exclude=exc&expansions=ex&max_results=50&media.fields=mf&pagination_token=ptoken&place.fields=plf&poll.fields=pof&since_id=sid&start_time=2021-12-26T06%3A15%3A09Z&tweet.fields=tf&until_id=uid&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.TweetTimelinesTweetsParams{
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

func Test_TweetTimelinesTweetsParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetTimelinesTweetsParams
	}{
		{
			name:   "empty params",
			params: &types.TweetTimelinesTweetsParams{},
		},
		{
			name:   "some params",
			params: &types.TweetTimelinesTweetsParams{ID: "id"},
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

func Test_TweetTimelinesMentionsParams_SetAccessToken(t *testing.T) {
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
			p := &types.TweetTimelinesMentionsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_TweetTimelinesMentionsParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	startTime := time.Unix(1640499309, 0).UTC()
	endTime := time.Unix(1640585709, 0).UTC()
	cases := []struct {
		name   string
		params *types.TweetTimelinesMentionsParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.TweetTimelinesMentionsParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "with exclude",
			params: &types.TweetTimelinesMentionsParams{
				ID:      "test-id",
				Exclude: fields.ExcludeList{"exc1", "exc2"},
			},
			expect: endpointRoot + "test-id" + "?exclude=exc1%2Cexc2",
		},
		{
			name: "with expansions",
			params: &types.TweetTimelinesMentionsParams{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with max_results",
			params: &types.TweetTimelinesMentionsParams{
				ID:         "test-id",
				MaxResults: 50,
			},
			expect: endpointRoot + "test-id" + "?max_results=50",
		},
		{
			name: "with media.fields",
			params: &types.TweetTimelinesMentionsParams{
				ID:          "test-id",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.TweetTimelinesMentionsParams{
				ID:          "test-id",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.TweetTimelinesMentionsParams{
				ID:         "test-id",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.TweetTimelinesMentionsParams{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.TweetTimelinesMentionsParams{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.TweetTimelinesMentionsParams{
				ID:              "test-id",
				Exclude:         fields.ExcludeList{"exc"},
				Expansions:      fields.ExpansionList{"ex"},
				MaxResults:      50,
				MediaFields:     fields.MediaFieldList{"mf"},
				PlaceFields:     fields.PlaceFieldList{"plf"},
				PollFields:      fields.PollFieldList{"pof"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
				StartTime:       &startTime,
				EndTime:         &endTime,
				SinceID:         "sid",
				UntilID:         "uid",
				PaginationToken: "ptoken",
			},
			expect: endpointRoot + "test-id" + "?end_time=2021-12-27T06%3A15%3A09Z&exclude=exc&expansions=ex&max_results=50&media.fields=mf&pagination_token=ptoken&place.fields=plf&poll.fields=pof&since_id=sid&start_time=2021-12-26T06%3A15%3A09Z&tweet.fields=tf&until_id=uid&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.TweetTimelinesMentionsParams{
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

func Test_TweetTimelinesMentionsParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetTimelinesMentionsParams
	}{
		{
			name:   "empty params",
			params: &types.TweetTimelinesMentionsParams{},
		},
		{
			name:   "some params",
			params: &types.TweetTimelinesMentionsParams{ID: "id"},
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
