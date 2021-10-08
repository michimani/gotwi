package params

import (
	"io"
	"net/url"

	"github.com/michimani/gotwi/types"
)

type UsersByParams struct {
	accessToken string

	// Query parameters
	Usernames   []string
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UsersByParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UsersByParams) AccessToken() string {
	return p.accessToken
}

func (p *UsersByParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Usernames == nil || len(p.Usernames) == 0 {
		return ""
	}

	query := url.Values{}
	query.Add("usernames", types.QueryValue(p.Usernames))

	if p.Expansions != nil {
		query.Add("expansions", types.QueryValue(p.Expansions))
	}

	if p.TweetFields != nil {
		query.Add("tweet.fields", types.QueryValue(p.TweetFields))
	}

	if p.UserFields != nil {
		query.Add("user.fields", types.QueryValue(p.UserFields))
	}

	if p.Expansions != nil && len(p.Expansions) > 0 {
	}

	if query.Has("expansions") || query.Has("tweet.fields") || query.Has("user.fields") {
		endpoint = endpoint + "?" + query.Encode()
	}

	return endpoint
}

func (p *UsersByParams) Body() io.Reader {
	return nil
}
