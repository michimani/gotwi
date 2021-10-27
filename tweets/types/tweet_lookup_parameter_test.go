package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetLookupParams_SetAccessToken(t *testing.T) {
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
			p := &types.TweetLookupParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_TweetLookupParams_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	cases := []struct {
		name   string
		params *types.TweetLookupParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.TweetLookupParams{IDs: []string{"test-id1", "test-id2"}},
			expect: endpointBase + "?ids=test-id1%2Ctest-id2",
		},
		{
			name: "with expansions",
			params: &types.TweetLookupParams{
				IDs:        []string{"test-id"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=test-id",
		},
		{
			name: "with media.fields",
			params: &types.TweetLookupParams{
				IDs:         []string{"test-id"},
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.TweetLookupParams{
				IDs:         []string{"test-id"},
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.TweetLookupParams{
				IDs:        []string{"test-id"},
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.TweetLookupParams{
				IDs:         []string{"test-id"},
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?ids=test-id&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.TweetLookupParams{
				IDs:        []string{"test-id"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=test-id&user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.TweetLookupParams{
				IDs:         []string{"test-id"},
				Expansions:  fields.ExpansionList{"ex"},
				MediaFields: fields.MediaFieldList{"mf"},
				PlaceFields: fields.PlaceFieldList{"plf"},
				PollFields:  fields.PollFieldList{"pof"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointBase + "?expansions=ex&ids=test-id&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.TweetLookupParams{
				IDs:         []string{},
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

func Test_TweetLookupParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetLookupParams
	}{
		{
			name:   "empty params",
			params: &types.TweetLookupParams{},
		},
		{
			name:   "some params",
			params: &types.TweetLookupParams{IDs: []string{"id"}},
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

func Test_TweetLookupIDParams_SetAccessToken(t *testing.T) {
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
			p := &types.TweetLookupIDParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_TweetLookupIDParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = endpointRoot + ":id"
	cases := []struct {
		name   string
		params *types.TweetLookupIDParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.TweetLookupIDParams{ID: "test-id"},
			expect: endpointRoot + "test-id",
		},
		{
			name: "with expansions",
			params: &types.TweetLookupIDParams{
				ID:         "test-id",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "test-id" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.TweetLookupIDParams{
				ID:          "test-id",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.TweetLookupIDParams{
				ID:          "test-id",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.TweetLookupIDParams{
				ID:         "test-id",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.TweetLookupIDParams{
				ID:          "test-id",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "test-id" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.TweetLookupIDParams{
				ID:         "test-id",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "test-id" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.TweetLookupIDParams{
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
			params: &types.TweetLookupIDParams{
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

func Test_TweetLookupIDParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.TweetLookupIDParams
	}{
		{
			name:   "empty params",
			params: &types.TweetLookupIDParams{},
		},
		{
			name:   "some params",
			params: &types.TweetLookupIDParams{ID: "id"},
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
