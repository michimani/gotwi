package types

import (
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/michimani/gotwi/internal/util"
)

type SearchTweetsMaxResult int

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
	MaxResults  SearchTweetsMaxResult
}

func (m SearchTweetsMaxResult) Valid() bool {
	return m > 10 && m <= 100
}

func (m SearchTweetsMaxResult) String() string {
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

	query := url.Values{}
	query.Add("query", p.Query)
	return endpoint + resolveSearchTweetsQuery(query,
		p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields,
		p.StartTime, p.EndTime,
		p.SinceID, p.UntilID,
		p.NextToken, p.MaxResults,
	)
}

func (p *SearchTweetsRecentParams) Body() io.Reader {
	return nil
}

func (p *SearchTweetsRecentParams) ParameterMap() map[string]string {
	m := map[string]string{}

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
	MaxResults  SearchTweetsMaxResult
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

	query := url.Values{}
	query.Add("query", p.Query)
	return endpoint + resolveSearchTweetsQuery(query,
		p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields,
		p.StartTime, p.EndTime,
		p.SinceID, p.UntilID,
		p.NextToken, p.MaxResults,
	)
}

func (p *SearchTweetsAllParams) Body() io.Reader {
	return nil
}

func (p *SearchTweetsAllParams) ParameterMap() map[string]string {
	m := map[string]string{}

	return m
}

func resolveSearchTweetsQuery(q url.Values,
	expansions, mediaFields, placeFields, pollFields, tweetFields, userFields []string,
	start, end *time.Time,
	since, until string,
	nextToken string, max SearchTweetsMaxResult,
) string {
	if expansions != nil {
		q.Add("expansions", util.QueryValue(expansions))
	}

	if mediaFields != nil {
		q.Add("media.fields", util.QueryValue(mediaFields))
	}

	if placeFields != nil {
		q.Add("place.fields", util.QueryValue(placeFields))
	}

	if pollFields != nil {
		q.Add("poll.fields", util.QueryValue(pollFields))
	}

	if tweetFields != nil {
		q.Add("tweet.fields", util.QueryValue(tweetFields))
	}

	if userFields != nil {
		q.Add("user.fields", util.QueryValue(userFields))
	}

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

	if max.Valid() {
		q.Add("max_results", max.String())
	} else {
		q.Add("max_results", "10")
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
