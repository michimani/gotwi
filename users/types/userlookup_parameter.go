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

	if p.Expansions != nil {
		query.Add("expansions", util.QueryValue(p.Expansions))
	}

	if p.TweetFields != nil {
		query.Add("tweet.fields", util.QueryValue(p.TweetFields))
	}

	if p.UserFields != nil {
		query.Add("user.fields", util.QueryValue(p.UserFields))
	}

	if query.Has("ids") || query.Has("expansions") || query.Has("tweet.fields") || query.Has("user.fields") {
		endpoint = endpoint + "?" + query.Encode()
	}

	return endpoint
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
	if p.Expansions != nil {
		query.Add("expansions", util.QueryValue(p.Expansions))
	}

	if p.TweetFields != nil {
		query.Add("tweet.fields", util.QueryValue(p.TweetFields))
	}

	if p.UserFields != nil {
		query.Add("user.fields", util.QueryValue(p.UserFields))
	}

	if p.Expansions != nil && len(p.Expansions) > 0 {
	}

	if query.Has("expansions") || query.Has("tweet.fields") || query.Has("user.fields") {
		endpoint = endpoint + "?" + query.Encode()
	}

	return endpoint
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

	if p.Expansions != nil {
		query.Add("expansions", util.QueryValue(p.Expansions))
	}

	if p.TweetFields != nil {
		query.Add("tweet.fields", util.QueryValue(p.TweetFields))
	}

	if p.UserFields != nil {
		query.Add("user.fields", util.QueryValue(p.UserFields))
	}

	if p.Expansions != nil && len(p.Expansions) > 0 {
	}

	if query.Has("usernames") || query.Has("expansions") || query.Has("tweet.fields") || query.Has("user.fields") {
		endpoint = endpoint + "?" + query.Encode()
	}

	return endpoint
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
	if p.Expansions != nil {
		query.Add("expansions", util.QueryValue(p.Expansions))
	}

	if p.TweetFields != nil {
		query.Add("tweet.fields", util.QueryValue(p.TweetFields))
	}

	if p.UserFields != nil {
		query.Add("user.fields", util.QueryValue(p.UserFields))
	}

	if p.Expansions != nil && len(p.Expansions) > 0 {
	}

	if query.Has("expansions") || query.Has("tweet.fields") || query.Has("user.fields") {
		endpoint = endpoint + "?" + query.Encode()
	}

	return endpoint
}

func (p *UserLookupUsersByUsernameParams) Body() io.Reader {
	return nil
}
