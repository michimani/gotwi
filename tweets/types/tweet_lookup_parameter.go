package types

import (
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/internal/util"
)

type TweetLookupParams struct {
	accessToken string

	// Query parameters
	IDs         []string
	Expansions  []string
	MediaFields []string
	PlaceFields []string
	PollFields  []string
	TweetFields []string
	UserFields  []string
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

	if p.Expansions != nil && len(p.Expansions) > 0 {
		m["expansions"] = util.QueryValue(p.Expansions)
	}

	if p.MediaFields != nil && len(p.MediaFields) > 0 {
		m["media.fields"] = util.QueryValue(p.MediaFields)
	}

	if p.PlaceFields != nil && len(p.PlaceFields) > 0 {
		m["place.fields"] = util.QueryValue(p.PlaceFields)
	}

	if p.PollFields != nil && len(p.PollFields) > 0 {
		m["poll.fields"] = util.QueryValue(p.PollFields)
	}

	if p.TweetFields != nil && len(p.TweetFields) > 0 {
		m["tweet.fields"] = util.QueryValue(p.TweetFields)
	}

	if p.UserFields != nil && len(p.UserFields) > 0 {
		m["user.fields"] = util.QueryValue(p.UserFields)
	}

	return m
}

type TweetLookupIDParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	Expansions  []string
	MediaFields []string
	PlaceFields []string
	PollFields  []string
	TweetFields []string
	UserFields  []string
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

	if p.Expansions != nil && len(p.Expansions) > 0 {
		m["expansions"] = util.QueryValue(p.Expansions)
	}

	if p.MediaFields != nil && len(p.MediaFields) > 0 {
		m["media.fields"] = util.QueryValue(p.MediaFields)
	}

	if p.PlaceFields != nil && len(p.PlaceFields) > 0 {
		m["place.fields"] = util.QueryValue(p.PlaceFields)
	}

	if p.PollFields != nil && len(p.PollFields) > 0 {
		m["poll.fields"] = util.QueryValue(p.PollFields)
	}

	if p.TweetFields != nil && len(p.TweetFields) > 0 {
		m["tweet.fields"] = util.QueryValue(p.TweetFields)
	}

	if p.UserFields != nil && len(p.UserFields) > 0 {
		m["user.fields"] = util.QueryValue(p.UserFields)
	}

	return m
}
