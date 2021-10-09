package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/internal/util"
)

type TweetLookupTweetsParams struct {
	accessToken string

	// Query parameters
	IDs         []string
	Expansions  []string
	MediaFields []string
	PlaceFields []string
	PollFields  []string
	TweetFields []string
	UserFields  []string
}

func (p *TweetLookupTweetsParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetLookupTweetsParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetLookupTweetsParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.IDs == nil || len(p.IDs) == 0 {
		return ""
	}

	query := url.Values{}
	query.Add("ids", util.QueryValue(p.IDs))
	return endpoint + resolveTweetLookupQuery(query,
		p.Expansions,
		p.MediaFields,
		p.PlaceFields,
		p.PollFields,
		p.TweetFields,
		p.UserFields,
	)
}

func (p *TweetLookupTweetsParams) Body() io.Reader {
	return nil
}

type TweetLookupTweetsIDParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	Expansions  []string
	MediaFields []string
	PlaceFields []string
	PollFields  []string
	TweetFields []string
	UserFields  []string
}

func (p *TweetLookupTweetsIDParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetLookupTweetsIDParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetLookupTweetsIDParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	query := url.Values{}
	return endpoint + resolveTweetLookupQuery(query,
		p.Expansions,
		p.MediaFields,
		p.PlaceFields,
		p.PollFields,
		p.TweetFields,
		p.UserFields,
	)
}

func (p *TweetLookupTweetsIDParams) Body() io.Reader {
	return nil
}

func resolveTweetLookupQuery(q url.Values, expansions, mediaFields, placeFields, pollFields, tweetFields, userFields []string) string {
	if expansions != nil {
		q.Add("expansions", util.QueryValue(expansions))
	}

	if mediaFields != nil {
		q.Add("media.fields", util.QueryValue(mediaFields))
	}

	if placeFields != nil {
		q.Add("place.fields", util.QueryValue(placeFields))
	}

	if pollFields != nil {
		q.Add("poll.fields", util.QueryValue(pollFields))
	}

	if tweetFields != nil {
		q.Add("tweet.fields", util.QueryValue(tweetFields))
	}

	if userFields != nil {
		q.Add("user.fields", util.QueryValue(userFields))
	}

	encoded := q.Encode()
	if encoded == "" {
		return ""
	}

	return "?" + encoded
}
