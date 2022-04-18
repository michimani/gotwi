package types

import (
	"io"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type SampleStreamInput struct {
	accessToken string

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

func (p *SampleStreamInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SampleStreamInput) AccessToken() string {
	return p.accessToken
}

func (p *SampleStreamInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, getQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *SampleStreamInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *SampleStreamInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}
