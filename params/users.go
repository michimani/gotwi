package params

import (
	"io"
	"net/url"

	"github.com/michimani/gotwi/types"
)

type UsersParams struct {
	accessToken string

	// Query parameters
	IDs         []string
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UsersParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UsersParams) AccessToken() string {
	return p.accessToken
}

func (p *UsersParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.IDs == nil || len(p.IDs) == 0 {
		return ""
	}

	query := url.Values{}
	query.Add("ids", types.QueryValue(p.IDs))

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

func (p *UsersParams) Body() io.Reader {
	return nil
}
