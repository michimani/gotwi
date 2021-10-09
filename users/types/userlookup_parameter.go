package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/internal/util"
)

type UserLookupUsersParams struct {
	accessToken string

	// Query parameters
	IDs         []string
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UserLookupUsersParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UserLookupUsersParams) AccessToken() string {
	return p.accessToken
}

func (p *UserLookupUsersParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.IDs == nil || len(p.IDs) == 0 {
		return ""
	}

	query := url.Values{}
	query.Add("ids", util.QueryValue(p.IDs))
	return endpoint + resolveUserLookupQuery(query, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *UserLookupUsersParams) Body() io.Reader {
	return nil
}

type UserLookupUsersIDParams struct {
	accessToken string

	// Path parameters
	ID string

	// Query parameters
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UserLookupUsersIDParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UserLookupUsersIDParams) AccessToken() string {
	return p.accessToken
}

func (p *UserLookupUsersIDParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	query := url.Values{}
	return endpoint + resolveUserLookupQuery(query, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *UserLookupUsersIDParams) Body() io.Reader {
	return nil
}

type UserLookupUsersByParams struct {
	accessToken string

	// Query parameters
	Usernames   []string
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UserLookupUsersByParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UserLookupUsersByParams) AccessToken() string {
	return p.accessToken
}

func (p *UserLookupUsersByParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Usernames == nil || len(p.Usernames) == 0 {
		return ""
	}

	query := url.Values{}
	query.Add("usernames", util.QueryValue(p.Usernames))
	return endpoint + resolveUserLookupQuery(query, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *UserLookupUsersByParams) Body() io.Reader {
	return nil
}

type UserLookupUsersByUsernameParams struct {
	accessToken string

	// Path parameters
	Username string

	// Query parameters
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *UserLookupUsersByUsernameParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UserLookupUsersByUsernameParams) AccessToken() string {
	return p.accessToken
}

func (p *UserLookupUsersByUsernameParams) ResolveEndpoint(endpointBase string) string {
	if p.Username == "" {
		return ""
	}

	encoded := url.QueryEscape(p.Username)
	endpoint := strings.Replace(endpointBase, ":username", encoded, 1)

	query := url.Values{}
	return endpoint + resolveUserLookupQuery(query, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *UserLookupUsersByUsernameParams) Body() io.Reader {
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
