package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

// ListUsersInput is struct for requesting `GET /2/users`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
type ListUsersInput struct {
	accessToken string

	// Query parameters
	IDs         []string // required
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var listUsersQueryParameters = map[string]struct{}{
	"ids":          {},
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *ListUsersInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListUsersInput) AccessToken() string {
	return p.accessToken
}

func (p *ListUsersInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.IDs == nil || len(p.IDs) == 0 {
		return ""
	}

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listUsersQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListUsersInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListUsersInput) ParameterMap() map[string]string {
	m := map[string]string{}

	m["ids"] = util.QueryValue(p.IDs)

	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

// GetUserInput is struct for requesting `GET /2/users/:id`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
type GetUserInput struct {
	accessToken string

	// Path parameters
	ID string // required

	// Query parameters
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var getUserQueryParameters = map[string]struct{}{
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *GetUserInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *GetUserInput) AccessToken() string {
	return p.accessToken
}

func (p *GetUserInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, getUserQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *GetUserInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *GetUserInput) ParameterMap() map[string]string {
	m := map[string]string{}

	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

// ListUsersByUsernamesInput is struct for requesting `GET /2/users/by`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by
type ListUsersByUsernamesInput struct {
	accessToken string

	// Query parameters
	Usernames   []string // required
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var listUsersByUsernamesQueryParameters = map[string]struct{}{
	"usernames":    {},
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *ListUsersByUsernamesInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListUsersByUsernamesInput) AccessToken() string {
	return p.accessToken
}

func (p *ListUsersByUsernamesInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Usernames == nil || len(p.Usernames) == 0 {
		return ""
	}

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listUsersByUsernamesQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListUsersByUsernamesInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListUsersByUsernamesInput) ParameterMap() map[string]string {
	m := map[string]string{}

	m["usernames"] = util.QueryValue(p.Usernames)

	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

// GetUserByUsernameInput is struct for requesting `GET /2/users/by/username/:username`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
type GetUserByUsernameInput struct {
	accessToken string

	// Path parameters
	Username string // required

	// Query parameters
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var getUserByUsernameQueryParameters = map[string]struct{}{
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *GetUserByUsernameInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *GetUserByUsernameInput) AccessToken() string {
	return p.accessToken
}

func (p *GetUserByUsernameInput) ResolveEndpoint(endpointBase string) string {
	if p.Username == "" {
		return ""
	}

	encoded := url.QueryEscape(p.Username)
	endpoint := strings.Replace(endpointBase, ":username", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, getUserByUsernameQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *GetUserByUsernameInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *GetUserByUsernameInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

// GetMeInput is struct for requesting `GET /2/users/me`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
type GetMeInput struct {
	accessToken string

	// Query parameters
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var getMeQueryParameters = map[string]struct{}{
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *GetMeInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *GetMeInput) AccessToken() string {
	return p.accessToken
}

func (p *GetMeInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, getMeQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *GetMeInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *GetMeInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)
	return m
}
