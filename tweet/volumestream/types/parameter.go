package types

import (
	"io"
	"strconv"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type SampleStreamBackfillMinutes int

func (v SampleStreamBackfillMinutes) Valid() bool {
	return int(v) > 0 && int(v) <= 5
}

func (v SampleStreamBackfillMinutes) String() string {
	return strconv.Itoa(int(v))
}

type SampleStreamInput struct {
	accessToken string

	// Query parameters
	BackfillMinutes SampleStreamBackfillMinutes
	Expansions      fields.ExpansionList
	MediaFields     fields.MediaFieldList
	PlaceFields     fields.PlaceFieldList
	PollFields      fields.PollFieldList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var getQueryParameters = map[string]struct{}{
	"backfill_minutes": {},
	"expansions":       {},
	"media.fields":     {},
	"place.fields":     {},
	"poll.fields":      {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *SampleStreamInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SampleStreamInput) AccessToken() string {
	return p.accessToken
}

func (p *SampleStreamInput) ResolveEndpoint(endpointBase string) string {
	if p == nil {
		return ""
	}

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

	if p.BackfillMinutes.Valid() {
		m["backfill_minutes"] = p.BackfillMinutes.String()
	}

	return m
}
