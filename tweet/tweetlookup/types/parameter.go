package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type ListInput struct {
	accessToken string

	// Query parameters
	IDs         []string
	Expansions  fields.ExpansionList
	MediaFields fields.MediaFieldList
	PlaceFields fields.PlaceFieldList
	PollFields  fields.PollFieldList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var listQueryParameters = map[string]struct{}{
	"ids":          {},
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *ListInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListInput) AccessToken() string {
	return p.accessToken
}

func (p *ListInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.IDs == nil || len(p.IDs) == 0 {
		return ""
	}

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListInput) ParameterMap() map[string]string {
	m := map[string]string{}

	m["ids"] = util.QueryValue(p.IDs)

	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}

type GetInput struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	Expansions  fields.ExpansionList
	MediaFields fields.MediaFieldList
	PlaceFields fields.PlaceFieldList
	PollFields  fields.PollFieldList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var getQueryParameters = map[string]struct{}{
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *GetInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *GetInput) AccessToken() string {
	return p.accessToken
}

func (p *GetInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, getQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *GetInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *GetInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}
