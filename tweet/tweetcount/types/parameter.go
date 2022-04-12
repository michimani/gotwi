package types

import (
	"io"
	"time"

	"github.com/michimani/gotwi/internal/util"
)

type TweetCountsGranularity string

const (
	TweetCountsGranularityMinute TweetCountsGranularity = "minute"
	TweetCountsGranularityHour   TweetCountsGranularity = "hour" // default
	TweetCountsGranularityDay    TweetCountsGranularity = "day"
)

func (g TweetCountsGranularity) String() string {
	return string(g)
}

func (g TweetCountsGranularity) Valid() bool {
	return g == TweetCountsGranularityMinute || g == TweetCountsGranularityHour || g == TweetCountsGranularityDay
}

type ListRecentInput struct {
	accessToken string

	// Path parameters
	Query       string
	StartTime   *time.Time
	EndTime     *time.Time
	SinceID     string
	UntilID     string
	Granularity TweetCountsGranularity
}

var listRecentQueryParameters = map[string]struct{}{
	"query":       {},
	"start_time":  {},
	"end_time":    {},
	"since_id":    {},
	"until_id":    {},
	"granularity": {},
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

	if p.Granularity.Valid() {
		m["granularity"] = p.Granularity.String()
	} else {
		m["granularity"] = TweetCountsGranularityHour.String()
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
	Granularity TweetCountsGranularity
	NextToken   string
}

var listAllQueryParameters = map[string]struct{}{
	"query":       {},
	"start_time":  {},
	"end_time":    {},
	"since_id":    {},
	"until_id":    {},
	"granularity": {},
	"next_token":  {},
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

	if p.Granularity.Valid() {
		m["granularity"] = p.Granularity.String()
	} else {
		m["granularity"] = TweetCountsGranularityHour.String()
	}

	if p.NextToken != "" {
		m["next_token"] = p.NextToken
	}

	return m
}
