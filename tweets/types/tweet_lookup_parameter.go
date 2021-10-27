package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type TweetLookupParams struct {
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

var TweetLookupQueryParams = map[string]struct{}{
	"ids":          {},
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *TweetLookupParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetLookupParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetLookupParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.IDs == nil || len(p.IDs) == 0 {
		return ""
	}

	pm := p.ParameterMap()
	qs := util.QueryString(pm, TweetLookupQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TweetLookupParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetLookupParams) ParameterMap() map[string]string {
	m := map[string]string{}

	m["ids"] = util.QueryValue(p.IDs)

	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}

type TweetLookupIDParams struct {
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

var TweetLookupIDQueryParams = map[string]struct{}{
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *TweetLookupIDParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetLookupIDParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetLookupIDParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, TweetLookupIDQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TweetLookupIDParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetLookupIDParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}
