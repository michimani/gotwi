package types

import (
	"io"
	"net/url"

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
	return endpoint + p.resolveTweetLookupQuery(query)
}

func (p *TweetLookupTweetsParams) Body() io.Reader {
	return nil
}

func (p *TweetLookupTweetsParams) resolveTweetLookupQuery(q url.Values) string {
	if p.Expansions != nil {
		q.Add("expansions", util.QueryValue(p.Expansions))
	}

	if p.MediaFields != nil {
		q.Add("media.fields", util.QueryValue(p.MediaFields))
	}

	if p.PlaceFields != nil {
		q.Add("place.fields", util.QueryValue(p.PlaceFields))
	}

	if p.PollFields != nil {
		q.Add("poll.fields", util.QueryValue(p.PollFields))
	}

	if p.TweetFields != nil {
		q.Add("tweet.fields", util.QueryValue(p.TweetFields))
	}

	if p.UserFields != nil {
		q.Add("user.fields", util.QueryValue(p.UserFields))
	}

	encoded := q.Encode()
	if encoded == "" {
		return ""
	}

	return "?" + encoded
}
