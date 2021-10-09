package types_test

import (
	"testing"

	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetLookupTweetsParams_SetAccessToken(t *testing.T) {
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
			p := &types.TweetLookupTweetsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_TweetLookupTweetsParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.TweetLookupTweetsParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.TweetLookupTweetsParams{IDs: []string{"test-id1", "test-id2"}},
			expect: endpointBase + "?ids=test-id1%2Ctest-id2",
		},
		{
			name: "with expansions",
			params: &types.TweetLookupTweetsParams{
				IDs:        []string{"test-id"},
				Expansions: []string{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=test-id",
		},
		{
			name: "with media.fields",
			params: &types.TweetLookupTweetsParams{
				IDs:         []string{"test-id"},
				MediaFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.TweetLookupTweetsParams{
				IDs:         []string{"test-id"},
				PlaceFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.TweetLookupTweetsParams{
				IDs:        []string{"test-id"},
				PollFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.TweetLookupTweetsParams{
				IDs:         []string{"test-id"},
				TweetFields: []string{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.TweetLookupTweetsParams{
				IDs:        []string{"test-id"},
				UserFields: []string{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=test-id&user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.TweetLookupTweetsParams{
				IDs:         []string{"test-id"},
				Expansions:  []string{"ex"},
				MediaFields: []string{"mf"},
				PlaceFields: []string{"plf"},
				PollFields:  []string{"pof"},
				UserFields:  []string{"uf"},
				TweetFields: []string{"tf"},
			},
			expect: endpointBase + "?expansions=ex&ids=test-id&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.TweetLookupTweetsParams{
				IDs:         []string{},
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

func Test_TweetLookupTweetsParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetLookupTweetsParams
	}{
		{
			name:   "empty params",
			params: &types.TweetLookupTweetsParams{},
		},
		{
			name:   "some params",
			params: &types.TweetLookupTweetsParams{IDs: []string{"id"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r := c.params.Body()
			assert.Nil(tt, r)
		})
	}
}
