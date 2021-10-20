package types

import (
	"io"
	"strconv"
	"time"

	"github.com/michimani/gotwi/internal/util"
)

type SearchTweetsMaxResults int

type SearchTweetsRecentParams struct {
	accessToken string

	// Path parameters
	Query       string
	StartTime   *time.Time
	EndTime     *time.Time
	SinceID     string
	UntilID     string
	Expansions  []string
	MediaFields []string
	PlaceFields []string
	PollFields  []string
	TweetFields []string
	UserFields  []string
	NextToken   string
	MaxResults  SearchTweetsMaxResults
}

var SearchTweetsRecentQueryParams = map[string]struct{}{
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

func (m SearchTweetsMaxResults) Valid() bool {
	return m > 10 && m <= 100
}

func (m SearchTweetsMaxResults) String() string {
	return strconv.Itoa(int(m))
}

func (p *SearchTweetsRecentParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SearchTweetsRecentParams) AccessToken() string {
	return p.accessToken
}

func (p *SearchTweetsRecentParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Query == "" {
		return ""
	}

	pm := p.ParameterMap()
	qs := util.QueryString(pm, SearchTweetsRecentQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *SearchTweetsRecentParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *SearchTweetsRecentParams) ParameterMap() map[string]string {
	m := map[string]string{}

	m["query"] = p.Query

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
	} else {
		m["max_results"] = "10"
	}

	if p.NextToken != "" {
		m["next_token"] = p.NextToken
	}

	return m
}

type SearchTweetsAllParams struct {
	accessToken string

	// Path parameters
	Query       string
	StartTime   *time.Time
	EndTime     *time.Time
	SinceID     string
	UntilID     string
	Expansions  []string
	MediaFields []string
	PlaceFields []string
	PollFields  []string
	TweetFields []string
	UserFields  []string
	NextToken   string
	MaxResults  SearchTweetsMaxResults
}

var SearchTweetsAllQueryParams = map[string]struct{}{
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

func (p *SearchTweetsAllParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SearchTweetsAllParams) AccessToken() string {
	return p.accessToken
}

func (p *SearchTweetsAllParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Query == "" {
		return ""
	}

	pm := p.ParameterMap()
	qs := util.QueryString(pm, SearchTweetsAllQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *SearchTweetsAllParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *SearchTweetsAllParams) ParameterMap() map[string]string {
	m := map[string]string{}

	m["query"] = p.Query

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
	} else {
		m["max_results"] = "10"
	}

	if p.NextToken != "" {
		m["next_token"] = p.NextToken
	}

	return m
}
