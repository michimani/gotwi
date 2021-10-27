package types

import (
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type TweetTimelinesMaxResults int

func (m TweetTimelinesMaxResults) Valid() bool {
	return m >= 5 && m <= 100
}

func (m TweetTimelinesMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type TweetTimelinesTweetsParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	StartTime       *time.Time
	EndTime         *time.Time
	SinceID         string
	UntilID         string
	Exclude         fields.ExcludeList
	Expansions      fields.ExpansionList
	MediaFields     fields.MediaFieldList
	PlaceFields     fields.PlaceFieldList
	PollFields      fields.PollFieldList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
	PaginationToken string
	MaxResults      TweetTimelinesMaxResults
}

var TweetTimelinesTweetsQueryParams = map[string]struct{}{
	"id":               {},
	"exclude":          {},
	"expansions":       {},
	"media.fields":     {},
	"place.fields":     {},
	"poll.fields":      {},
	"tweet.fields":     {},
	"user.fields":      {},
	"start_time":       {},
	"end_time":         {},
	"since_id":         {},
	"until_id":         {},
	"max_results":      {},
	"pagination_token": {},
}

func (p *TweetTimelinesTweetsParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetTimelinesTweetsParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetTimelinesTweetsParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, TweetTimelinesTweetsQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TweetTimelinesTweetsParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetTimelinesTweetsParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Exclude, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

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

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	return m
}

type TweetTimelinesMentionsParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	StartTime       *time.Time
	EndTime         *time.Time
	SinceID         string
	UntilID         string
	Exclude         fields.ExcludeList
	Expansions      fields.ExpansionList
	MediaFields     fields.MediaFieldList
	PlaceFields     fields.PlaceFieldList
	PollFields      fields.PollFieldList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
	PaginationToken string
	MaxResults      TweetTimelinesMaxResults
}

var TweetTimelinesMentionsQueryParams = map[string]struct{}{
	"id":               {},
	"exclude":          {},
	"expansions":       {},
	"media.fields":     {},
	"place.fields":     {},
	"poll.fields":      {},
	"tweet.fields":     {},
	"user.fields":      {},
	"start_time":       {},
	"end_time":         {},
	"since_id":         {},
	"until_id":         {},
	"max_results":      {},
	"pagination_token": {},
}

func (p *TweetTimelinesMentionsParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetTimelinesMentionsParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetTimelinesMentionsParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, TweetTimelinesMentionsQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TweetTimelinesMentionsParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetTimelinesMentionsParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Exclude, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

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

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	return m
}
