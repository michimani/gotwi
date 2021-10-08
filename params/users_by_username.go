package params

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/types"
)

type UsersByUsernameParams struct {
	accessToken string

	// Path parameters
	Username string

	// Query parameters
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UsersByUsernameParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UsersByUsernameParams) AccessToken() string {
	return p.accessToken
}

func (p *UsersByUsernameParams) ResolveEndpoint(endpointBase string) string {
	if p.Username == "" {
		return ""
	}

	encoded := url.QueryEscape(p.Username)
	endpoint := strings.Replace(endpointBase, ":username", encoded, 1)

	query := url.Values{}
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

func (p *UsersByUsernameParams) Body() io.Reader {
	return nil
}
