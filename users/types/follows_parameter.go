package types

import (
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/michimani/gotwi/internal/util"
)

type FollowsMaxResult int

type FollowsFollowingGetParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	MaxResult       FollowsMaxResult
	PaginationToken string
	Expansions      []string
	TweetFields     []string
	UserFields      []string
}

func (m FollowsMaxResult) Valid() bool {
	return m > 0 && m <= 1000
}

func (m FollowsMaxResult) String() string {
	return strconv.Itoa(int(m))
}

func (p *FollowsFollowingGetParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *FollowsFollowingGetParams) AccessToken() string {
	return p.accessToken
}

func (p *FollowsFollowingGetParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	query := url.Values{}
	return endpoint + resolveFollowsQuery(query, p.MaxResult, p.PaginationToken, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *FollowsFollowingGetParams) Body() io.Reader {
	return nil
}

type FollowsFollowersParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	MaxResult       FollowsMaxResult
	PaginationToken string
	Expansions      []string
	TweetFields     []string
	UserFields      []string
}

func (p *FollowsFollowersParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *FollowsFollowersParams) AccessToken() string {
	return p.accessToken
}

func (p *FollowsFollowersParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	query := url.Values{}
	return endpoint + resolveFollowsQuery(query, p.MaxResult, p.PaginationToken, p.Expansions, p.TweetFields, p.UserFields)
}

func (p *FollowsFollowersParams) Body() io.Reader {
	return nil
}

func resolveFollowsQuery(q url.Values, max FollowsMaxResult, paginationToken string, expansions, tweetFields, userFields []string) string {
	if max.Valid() {
		q.Add("max_results", max.String())
	}

	if paginationToken != "" {
		q.Add("pagination_token", paginationToken)
	}

	if expansions != nil {
		q.Add("expansions", util.QueryValue(expansions))
	}

	if tweetFields != nil {
		q.Add("tweet.fields", util.QueryValue(tweetFields))
	}

	if userFields != nil {
		q.Add("user.fields", util.QueryValue(userFields))
	}

	encoded := q.Encode()
	if encoded == "" {
		return ""
	}

	return "?" + encoded
}
