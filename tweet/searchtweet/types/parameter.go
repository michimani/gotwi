package types

import (
	"io"
	"strconv"
	"time"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type ListMaxResults int

type ListRecentInput struct {
	accessToken string

	// Path parameters
	Query       string
	StartTime   *time.Time
	EndTime     *time.Time
	SinceID     string
	UntilID     string
	Expansions  fields.ExpansionList
	MediaFields fields.MediaFieldList
	PlaceFields fields.PlaceFieldList
	PollFields  fields.PollFieldList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
	NextToken   string
	MaxResults  ListMaxResults
}

var listRecentQueryParameters = map[string]struct{}{
	"query":        {},
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
	"start_time":   {},
	"end_time":     {},
	"since_id":     {},
	"until_id":     {},
	"max_results":  {},
	"next_token":   {},
}

func (m ListMaxResults) Valid() bool {
	return m >= 10 && m <= 100
}

func (m ListMaxResults) String() string {
	return strconv.Itoa(int(m))
}

func (p *ListRecentInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListRecentInput) AccessToken() string {
	return p.accessToken
}

func (p *ListRecentInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Query == "" {
		return ""
	}

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listRecentQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListRecentInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListRecentInput) ParameterMap() map[string]string {
	m := map[string]string{}

	m["query"] = p.Query

	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	if p.StartTime != nil {
		m["start_time"] = p.StartTime.Format(time.RFC3339)
	}

	if p.EndTime != nil {
		m["end_time"] = p.EndTime.Format(time.RFC3339)
	}

	if p.SinceID != "" {
		m["since_id"] = p.SinceID
	}

	if p.UntilID != "" {
		m["until_id"] = p.UntilID
	}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.NextToken != "" {
		m["next_token"] = p.NextToken
	}

	return m
}

type ListAllInput struct {
	accessToken string

	// Path parameters
	Query       string
	StartTime   *time.Time
	EndTime     *time.Time
	SinceID     string
	UntilID     string
	Expansions  fields.ExpansionList
	MediaFields fields.MediaFieldList
	PlaceFields fields.PlaceFieldList
	PollFields  fields.PollFieldList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
	NextToken   string
	MaxResults  ListMaxResults
}

var listAllQueryParameters = map[string]struct{}{
	"query":        {},
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
	"start_time":   {},
	"end_time":     {},
	"since_id":     {},
	"until_id":     {},
	"max_results":  {},
	"next_token":   {},
}

func (p *ListAllInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListAllInput) AccessToken() string {
	return p.accessToken
}

func (p *ListAllInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Query == "" {
		return ""
	}

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listAllQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListAllInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListAllInput) ParameterMap() map[string]string {
	m := map[string]string{}

	m["query"] = p.Query

	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	if p.StartTime != nil {
		m["start_time"] = p.StartTime.Format(time.RFC3339)
	}

	if p.EndTime != nil {
		m["end_time"] = p.EndTime.Format(time.RFC3339)
	}

	if p.SinceID != "" {
		m["since_id"] = p.SinceID
	}

	if p.UntilID != "" {
		m["until_id"] = p.UntilID
	}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.NextToken != "" {
		m["next_token"] = p.NextToken
	}

	return m
}
