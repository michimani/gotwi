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

type TweetCountsRecentParams struct {
	accessToken string

	// Path parameters
	Query       string
	StartTime   *time.Time
	EndTime     *time.Time
	SinceID     string
	UntilID     string
	Granularity TweetCountsGranularity
}

var TweetCountsRecentQueryParams = map[string]struct{}{
	"query":       {},
	"start_time":  {},
	"end_time":    {},
	"since_id":    {},
	"until_id":    {},
	"granularity": {},
}

func (p *TweetCountsRecentParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetCountsRecentParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetCountsRecentParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Query == "" {
		return ""
	}

	pm := p.ParameterMap()
	qs := util.QueryString(pm, TweetCountsRecentQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TweetCountsRecentParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetCountsRecentParams) ParameterMap() map[string]string {
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

type TweetCountsAllParams struct {
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

var TweetCountsAllQueryParams = map[string]struct{}{
	"query":       {},
	"start_time":  {},
	"end_time":    {},
	"since_id":    {},
	"until_id":    {},
	"granularity": {},
	"next_token":  {},
}

func (p *TweetCountsAllParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetCountsAllParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetCountsAllParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Query == "" {
		return ""
	}

	pm := p.ParameterMap()
	qs := util.QueryString(pm, TweetCountsAllQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TweetCountsAllParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetCountsAllParams) ParameterMap() map[string]string {
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
