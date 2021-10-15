package types

import (
	"io"
	"net/url"
	"time"
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

	query := url.Values{}
	query.Add("query", p.Query)
	return endpoint + resolveTweetCountsQuery(query,
		p.StartTime, p.EndTime,
		p.SinceID, p.UntilID,
		p.Granularity,
		"",
	)
}

func (p *TweetCountsRecentParams) Body() io.Reader {
	return nil
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

	query := url.Values{}
	query.Add("query", p.Query)
	return endpoint + resolveTweetCountsQuery(query,
		p.StartTime, p.EndTime,
		p.SinceID, p.UntilID,
		p.Granularity,
		p.NextToken,
	)
}

func (p *TweetCountsAllParams) Body() io.Reader {
	return nil
}

func resolveTweetCountsQuery(q url.Values, start, end *time.Time, since, until string, granularity TweetCountsGranularity, nextToken string) string {
	if start != nil {
		q.Add("start_time", start.Format(time.RFC3339))
	}

	if end != nil {
		q.Add("end_time", end.Format(time.RFC3339))
	}

	if since != "" {
		q.Add("since_id", since)
	}

	if until != "" {
		q.Add("until_id", until)
	}

	if granularity.Valid() {
		q.Add("granularity", granularity.String())
	} else {
		q.Add("granularity", TweetCountsGranularityHour.String())
	}

	if nextToken != "" {
		q.Add("next_token", nextToken)
	}

	encoded := q.Encode()
	if encoded == "" {
		return ""
	}

	return "?" + encoded
}
