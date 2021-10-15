package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/internal/util"
)

type UserLookupParams struct {
	accessToken string

	// Query parameters
	IDs         []string
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UserLookupParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UserLookupParams) AccessToken() string {
	return p.accessToken
}

func (p *UserLookupParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.IDs == nil || len(p.IDs) == 0 {
		return ""
	}

	query := url.Values{}
	query.Add("ids", util.QueryValue(p.IDs))
	return endpoint + resolveUserLookupQuery(query, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *UserLookupParams) Body() io.Reader {
	return nil
}

type UserLookupIDParams struct {
	accessToken string

	// Path parameters
	ID string

	// Query parameters
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UserLookupIDParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UserLookupIDParams) AccessToken() string {
	return p.accessToken
}

func (p *UserLookupIDParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	query := url.Values{}
	return endpoint + resolveUserLookupQuery(query, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *UserLookupIDParams) Body() io.Reader {
	return nil
}

type UserLookupByParams struct {
	accessToken string

	// Query parameters
	Usernames   []string
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UserLookupByParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UserLookupByParams) AccessToken() string {
	return p.accessToken
}

func (p *UserLookupByParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Usernames == nil || len(p.Usernames) == 0 {
		return ""
	}

	query := url.Values{}
	query.Add("usernames", util.QueryValue(p.Usernames))
	return endpoint + resolveUserLookupQuery(query, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *UserLookupByParams) Body() io.Reader {
	return nil
}

type UserLookupByUsernameParams struct {
	accessToken string

	// Path parameters
	Username string

	// Query parameters
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UserLookupByUsernameParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UserLookupByUsernameParams) AccessToken() string {
	return p.accessToken
}

func (p *UserLookupByUsernameParams) ResolveEndpoint(endpointBase string) string {
	if p.Username == "" {
		return ""
	}

	encoded := url.QueryEscape(p.Username)
	endpoint := strings.Replace(endpointBase, ":username", encoded, 1)

	query := url.Values{}
	return endpoint + resolveUserLookupQuery(query, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *UserLookupByUsernameParams) Body() io.Reader {
	return nil
}

func resolveUserLookupQuery(q url.Values, expansions, tweetFields, userFields []string) string {
	if expansions != nil {
		q.Add("expansions", util.QueryValue(expansions))
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
