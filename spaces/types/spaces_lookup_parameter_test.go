package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/spaces/types"
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
			p := &types.SpacesLookupIDParams{}
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
		params *types.SpacesLookupIDParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SpacesLookupIDParams{
				ID: "sid",
			},
			expect: endpointRoot + "sid",
		},
		{
			name: "with expansions",
			params: &types.SpacesLookupIDParams{
				ID:         "sid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with space.fields",
			params: &types.SpacesLookupIDParams{
				ID:          "sid",
				SpaceFields: fields.SpaceFieldList{"sf1", "sf2"},
			},
			expect: endpointRoot + "sid" + "?space.fields=sf1%2Csf2",
		},
		{
			name: "with users.fields",
			params: &types.SpacesLookupIDParams{
				ID:         "sid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "sid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SpacesLookupIDParams{
				Expansions:  fields.ExpansionList{"ex"},
				ID:          "sid",
				SpaceFields: fields.SpaceFieldList{"sf"},
				UserFields:  fields.UserFieldList{"uf"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex&space.fields=sf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.SpacesLookupIDParams{
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
		params *types.SpacesLookupIDParams
	}{
		{
			name:   "empty params",
			params: &types.SpacesLookupIDParams{},
		},
		{
			name:   "some params",
			params: &types.SpacesLookupIDParams{ID: "sid"},
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
			p := &types.SpacesLookupParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SpacesLookup_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"

	cases := []struct {
		name   string
		params *types.SpacesLookupParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SpacesLookupParams{
				IDs: []string{"sid1", "sid2"},
			},
			expect: endpointBase + "?ids=sid1%2Csid2",
		},
		{
			name: "with expansions",
			params: &types.SpacesLookupParams{
				IDs:        []string{"sid1", "sid2"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&ids=sid1%2Csid2",
		},
		{
			name: "with space.fields",
			params: &types.SpacesLookupParams{
				IDs:         []string{"sid1", "sid2"},
				SpaceFields: fields.SpaceFieldList{"sf1", "sf2"},
			},
			expect: endpointBase + "?ids=sid1%2Csid2&space.fields=sf1%2Csf2",
		},
		{
			name: "with users.fields",
			params: &types.SpacesLookupParams{
				IDs:        []string{"sid1", "sid2"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?ids=sid1%2Csid2&user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SpacesLookupParams{
				Expansions:  fields.ExpansionList{"ex"},
				IDs:         []string{"sid1", "sid2"},
				SpaceFields: fields.SpaceFieldList{"sf"},
				UserFields:  fields.UserFieldList{"uf"},
			},
			expect: endpointBase + "?expansions=ex&ids=sid1%2Csid2&space.fields=sf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.SpacesLookupParams{
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
		params *types.SpacesLookupParams
	}{
		{
			name:   "empty params",
			params: &types.SpacesLookupParams{},
		},
		{
			name:   "some params",
			params: &types.SpacesLookupParams{IDs: []string{"sid"}},
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
			p := &types.SpacesLookupByCreatorIDsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SpacesLookupByCreatorIDs_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint/"

	cases := []struct {
		name   string
		params *types.SpacesLookupByCreatorIDsParams
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.SpacesLookupByCreatorIDsParams{
				UserIDs: []string{"uid1", "uid2"},
			},
			expect: endpointBase + "?user_ids=uid1%2Cuid2",
		},
		{
			name: "with expansions",
			params: &types.SpacesLookupByCreatorIDsParams{
				UserIDs:    []string{"uid1", "uid2"},
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&user_ids=uid1%2Cuid2",
		},
		{
			name: "with space.fields",
			params: &types.SpacesLookupByCreatorIDsParams{
				UserIDs:     []string{"uid1", "uid2"},
				SpaceFields: fields.SpaceFieldList{"sf1", "sf2"},
			},
			expect: endpointBase + "?space.fields=sf1%2Csf2&user_ids=uid1%2Cuid2",
		},
		{
			name: "with users.fields",
			params: &types.SpacesLookupByCreatorIDsParams{
				UserIDs:    []string{"uid1", "uid2"},
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?user.fields=uf1%2Cuf2&user_ids=uid1%2Cuid2",
		},
		{
			name: "all query parameters",
			params: &types.SpacesLookupByCreatorIDsParams{
				Expansions:  fields.ExpansionList{"ex"},
				UserIDs:     []string{"uid1", "uid2"},
				SpaceFields: fields.SpaceFieldList{"sf"},
				UserFields:  fields.UserFieldList{"uf"},
			},
			expect: endpointBase + "?expansions=ex&space.fields=sf&user.fields=uf&user_ids=uid1%2Cuid2",
		},
		{
			name: "has no required parameter",
			params: &types.SpacesLookupByCreatorIDsParams{
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
		params *types.SpacesLookupByCreatorIDsParams
	}{
		{
			name:   "empty params",
			params: &types.SpacesLookupByCreatorIDsParams{},
		},
		{
			name:   "some params",
			params: &types.SpacesLookupByCreatorIDsParams{UserIDs: []string{"uid"}},
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

func Test_SpacesLookupBuyersParams_SetAccessToken(t *testing.T) {
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
			p := &types.SpacesLookupBuyersParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SpacesLookupBuyersParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.SpacesLookupBuyersParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.SpacesLookupBuyersParams{ID: "sid"},
			expect: endpointRoot + "sid",
		},
		{
			name: "with expansions",
			params: &types.SpacesLookupBuyersParams{
				ID:         "sid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.SpacesLookupBuyersParams{
				ID:          "sid",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.SpacesLookupBuyersParams{
				ID:          "sid",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.SpacesLookupBuyersParams{
				ID:         "sid",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.SpacesLookupBuyersParams{
				ID:          "sid",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.SpacesLookupBuyersParams{
				ID:         "sid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "sid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SpacesLookupBuyersParams{
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
			params: &types.SpacesLookupBuyersParams{
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

func Test_SpacesLookupBuyersParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SpacesLookupBuyersParams
	}{
		{
			name:   "empty params",
			params: &types.SpacesLookupBuyersParams{},
		},
		{
			name:   "some params",
			params: &types.SpacesLookupBuyersParams{ID: "id"},
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

func Test_SpacesLookupTweetsParams_SetAccessToken(t *testing.T) {
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
			p := &types.SpacesLookupTweetsParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SpacesLookupTweetsParams_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"

	cases := []struct {
		name   string
		params *types.SpacesLookupTweetsParams
		expect string
	}{
		{
			name:   "only required parameter",
			params: &types.SpacesLookupTweetsParams{ID: "sid"},
			expect: endpointRoot + "sid",
		},
		{
			name: "with expansions",
			params: &types.SpacesLookupTweetsParams{
				ID:         "sid",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointRoot + "sid" + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.SpacesLookupTweetsParams{
				ID:          "sid",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.SpacesLookupTweetsParams{
				ID:          "sid",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.SpacesLookupTweetsParams{
				ID:         "sid",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.SpacesLookupTweetsParams{
				ID:          "sid",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointRoot + "sid" + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.SpacesLookupTweetsParams{
				ID:         "sid",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointRoot + "sid" + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SpacesLookupTweetsParams{
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
			params: &types.SpacesLookupTweetsParams{
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

func Test_SpacesLookupTweetsParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SpacesLookupTweetsParams
	}{
		{
			name:   "empty params",
			params: &types.SpacesLookupTweetsParams{},
		},
		{
			name:   "some params",
			params: &types.SpacesLookupTweetsParams{ID: "id"},
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
