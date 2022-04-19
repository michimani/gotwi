package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweet/volumestream/types"
	"github.com/stretchr/testify/assert"
)

func Test_SampleStreamInput_SetAccessToken(t *testing.T) {
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
			p := &types.SampleStreamInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SampleStreamInput_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.SampleStreamInput
		expect string
	}{
		{
			name:   "ok",
			params: &types.SampleStreamInput{},
			expect: endpoint,
		},
		{
			name: "with expansions",
			params: &types.SampleStreamInput{
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpoint + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.SampleStreamInput{
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.SampleStreamInput{
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.SampleStreamInput{
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.SampleStreamInput{
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.SampleStreamInput{
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpoint + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SampleStreamInput{
				Expansions:  fields.ExpansionList{"ex"},
				MediaFields: fields.MediaFieldList{"mf"},
				PlaceFields: fields.PlaceFieldList{"plf"},
				PollFields:  fields.PollFieldList{"pof"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpoint + "?expansions=ex&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
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

func Test_SampleStreamInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SampleStreamInput
	}{
		{
			name:   "empty params",
			params: &types.SampleStreamInput{},
		},
		{
			name:   "some params",
			params: &types.SampleStreamInput{Expansions: fields.ExpansionList{"ex"}},
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
