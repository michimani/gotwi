package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/space/spacelookup/types"
	"github.com/stretchr/testify/assert"
)

func Test_SpacesLookupID_SetAccessToken(t *testing.T) {
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
			p := &types.GetInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SpacesLookupID_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.GetInput
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.GetInput{
				ID: "sid",
			},
			expect: endpointRoot + "sid",
		},
		{
			name: "with expansions",
			params: &types.GetInput{
				ID:         "sid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with space.fields",
			params: &types.GetInput{
				ID:          "sid",
				SpaceFields: fields.SpaceFieldList{"sf1", "sf2"},
			},
			expect: endpointRoot + "sid" + "?space.fields=sf1%2Csf2",
		},
		{
			name: "with users.fields",
			params: &types.GetInput{
				ID:         "sid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "sid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.GetInput{
				Expansions:  fields.ExpansionList{"ex"},
				ID:          "sid",
				SpaceFields: fields.SpaceFieldList{"sf"},
				UserFields:  fields.UserFieldList{"uf"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex&space.fields=sf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.GetInput{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				SpaceFields: fields.SpaceFieldList{"sf"},
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

func Test_SpacesLookupID_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.GetInput
	}{
		{
			name:   "empty params",
			params: &types.GetInput{},
		},
		{
			name:   "some params",
			params: &types.GetInput{ID: "sid"},
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

func Test_SpacesLookup_SetAccessToken(t *testing.T) {
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
			p := &types.ListInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SpacesLookup_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"

	cases := []struct {
		name   string
		params *types.ListInput
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListInput{
				IDs: []string{"sid1", "sid2"},
			},
			expect: endpointBase + "?ids=sid1%2Csid2",
		},
		{
			name: "with expansions",
			params: &types.ListInput{
				IDs:        []string{"sid1", "sid2"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=sid1%2Csid2",
		},
		{
			name: "with space.fields",
			params: &types.ListInput{
				IDs:         []string{"sid1", "sid2"},
				SpaceFields: fields.SpaceFieldList{"sf1", "sf2"},
			},
			expect: endpointBase + "?ids=sid1%2Csid2&space.fields=sf1%2Csf2",
		},
		{
			name: "with users.fields",
			params: &types.ListInput{
				IDs:        []string{"sid1", "sid2"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=sid1%2Csid2&user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.ListInput{
				Expansions:  fields.ExpansionList{"ex"},
				IDs:         []string{"sid1", "sid2"},
				SpaceFields: fields.SpaceFieldList{"sf"},
				UserFields:  fields.UserFieldList{"uf"},
			},
			expect: endpointBase + "?expansions=ex&ids=sid1%2Csid2&space.fields=sf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListInput{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				SpaceFields: fields.SpaceFieldList{"sf"},
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

func Test_SpacesLookup_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListInput
	}{
		{
			name:   "empty params",
			params: &types.ListInput{},
		},
		{
			name:   "some params",
			params: &types.ListInput{IDs: []string{"sid"}},
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

func Test_SpacesLookupByCreatorIDs_SetAccessToken(t *testing.T) {
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
			p := &types.ListByCreatorIDsInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SpacesLookupByCreatorIDs_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"

	cases := []struct {
		name   string
		params *types.ListByCreatorIDsInput
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListByCreatorIDsInput{
				UserIDs: []string{"uid1", "uid2"},
			},
			expect: endpointBase + "?user_ids=uid1%2Cuid2",
		},
		{
			name: "with expansions",
			params: &types.ListByCreatorIDsInput{
				UserIDs:    []string{"uid1", "uid2"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&user_ids=uid1%2Cuid2",
		},
		{
			name: "with space.fields",
			params: &types.ListByCreatorIDsInput{
				UserIDs:     []string{"uid1", "uid2"},
				SpaceFields: fields.SpaceFieldList{"sf1", "sf2"},
			},
			expect: endpointBase + "?space.fields=sf1%2Csf2&user_ids=uid1%2Cuid2",
		},
		{
			name: "with users.fields",
			params: &types.ListByCreatorIDsInput{
				UserIDs:    []string{"uid1", "uid2"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?user.fields=uf1%2Cuf2&user_ids=uid1%2Cuid2",
		},
		{
			name: "all query parameters",
			params: &types.ListByCreatorIDsInput{
				Expansions:  fields.ExpansionList{"ex"},
				UserIDs:     []string{"uid1", "uid2"},
				SpaceFields: fields.SpaceFieldList{"sf"},
				UserFields:  fields.UserFieldList{"uf"},
			},
			expect: endpointBase + "?expansions=ex&space.fields=sf&user.fields=uf&user_ids=uid1%2Cuid2",
		},
		{
			name: "has no required parameter",
			params: &types.ListByCreatorIDsInput{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				SpaceFields: fields.SpaceFieldList{"sf"},
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

func Test_SpacesLookupByCreatorIDs_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListByCreatorIDsInput
	}{
		{
			name:   "empty params",
			params: &types.ListByCreatorIDsInput{},
		},
		{
			name:   "some params",
			params: &types.ListByCreatorIDsInput{UserIDs: []string{"uid"}},
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

func Test_ListBuyersInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListBuyersInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListBuyersInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListBuyersInput
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.ListBuyersInput{ID: "sid"},
			expect: endpointRoot + "sid",
		},
		{
			name: "with expansions",
			params: &types.ListBuyersInput{
				ID:         "sid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.ListBuyersInput{
				ID:          "sid",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.ListBuyersInput{
				ID:          "sid",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.ListBuyersInput{
				ID:         "sid",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.ListBuyersInput{
				ID:          "sid",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.ListBuyersInput{
				ID:         "sid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "sid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.ListBuyersInput{
				ID:          "sid",
				Expansions:  fields.ExpansionList{"ex"},
				MediaFields: fields.MediaFieldList{"mf"},
				PlaceFields: fields.PlaceFieldList{"plf"},
				PollFields:  fields.PollFieldList{"pof"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListBuyersInput{
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

func Test_ListBuyersInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListBuyersInput
	}{
		{
			name:   "empty params",
			params: &types.ListBuyersInput{},
		},
		{
			name:   "some params",
			params: &types.ListBuyersInput{ID: "id"},
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

func Test_ListTweetsInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListTweetsInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListTweetsInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.ListTweetsInput
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.ListTweetsInput{ID: "sid"},
			expect: endpointRoot + "sid",
		},
		{
			name: "with expansions",
			params: &types.ListTweetsInput{
				ID:         "sid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.ListTweetsInput{
				ID:          "sid",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.ListTweetsInput{
				ID:          "sid",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.ListTweetsInput{
				ID:         "sid",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.ListTweetsInput{
				ID:          "sid",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.ListTweetsInput{
				ID:         "sid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "sid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.ListTweetsInput{
				ID:          "sid",
				Expansions:  fields.ExpansionList{"ex"},
				MediaFields: fields.MediaFieldList{"mf"},
				PlaceFields: fields.PlaceFieldList{"plf"},
				PollFields:  fields.PollFieldList{"pof"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListTweetsInput{
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

func Test_ListTweetsInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListTweetsInput
	}{
		{
			name:   "empty params",
			params: &types.ListTweetsInput{},
		},
		{
			name:   "some params",
			params: &types.ListTweetsInput{ID: "id"},
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
