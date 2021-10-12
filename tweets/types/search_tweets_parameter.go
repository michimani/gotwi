package types

import (
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/michimani/gotwi/internal/util"
)

type SearchTweetsMaxResult int

type SearchTweetsTweetsSearchRecentParams struct {
	accessToken string

	// Path parameters
	Query       string
	StartTime   *time.Time
	EndTime     *time.Time
	SinceID     string
	UntileID    string
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

func (p *SearchTweetsTweetsSearchRecentParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SearchTweetsTweetsSearchRecentParams) AccessToken() string {
	return p.accessToken
}

func (p *SearchTweetsTweetsSearchRecentParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Query == "" {
		return ""
	}

	query := url.Values{}
	query.Add("query", p.Query)
	return endpoint + resolveSearchTweetsQuery(query,
		p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields,
		p.StartTime, p.EndTime,
		p.SinceID, p.UntileID,
		p.NextToken, p.MaxResults,
	)
}

func (p *SearchTweetsTweetsSearchRecentParams) Body() io.Reader {
	return nil
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
