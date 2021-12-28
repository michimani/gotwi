package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type UserLookupParams struct {
	accessToken string

	// Query parameters
	IDs         []string
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var UserLookupQueryParams = map[string]struct{}{
	"ids":          {},
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
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

	pm := p.ParameterMap()
	qs := util.QueryString(pm, UserLookupQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *UserLookupParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *UserLookupParams) ParameterMap() map[string]string {
	m := map[string]string{}

	m["ids"] = util.QueryValue(p.IDs)

	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

type UserLookupIDParams struct {
	accessToken string

	// Path parameters
	ID string

	// Query parameters
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var UserLookupIDQueryParams = map[string]struct{}{
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
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

	pm := p.ParameterMap()
	qs := util.QueryString(pm, UserLookupIDQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *UserLookupIDParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *UserLookupIDParams) ParameterMap() map[string]string {
	m := map[string]string{}

	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

type UserLookupByParams struct {
	accessToken string

	// Query parameters
	Usernames   []string
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var UserLookupByQueryParams = map[string]struct{}{
	"usernames":    {},
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
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

	pm := p.ParameterMap()
	qs := util.QueryString(pm, UserLookupByQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *UserLookupByParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *UserLookupByParams) ParameterMap() map[string]string {
	m := map[string]string{}

	m["usernames"] = util.QueryValue(p.Usernames)

	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

type UserLookupByUsernameParams struct {
	accessToken string

	// Path parameters
	Username string

	// Query parameters
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var UserLookupByUsernameQueryParams = map[string]struct{}{
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
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

	pm := p.ParameterMap()
	qs := util.QueryString(pm, UserLookupByUsernameQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *UserLookupByUsernameParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *UserLookupByUsernameParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

type UserLookupMeParams struct {
	accessToken string

	// Query parameters
	Expansions  fields.ExpansionList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var UserLookupMeQueryParams = map[string]struct{}{
	"expansions":   {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *UserLookupMeParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UserLookupMeParams) AccessToken() string {
	return p.accessToken
}

func (p *UserLookupMeParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	pm := p.ParameterMap()
	qs := util.QueryString(pm, UserLookupMeQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *UserLookupMeParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *UserLookupMeParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)
	return m
}
