package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/michimani/gotwi/internal/util"
)

type FollowsMaxResults int

func (m FollowsMaxResults) Valid() bool {
	return m > 0 && m <= 1000
}

func (m FollowsMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type FollowsFollowingGetParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	MaxResults      FollowsMaxResults
	PaginationToken string
	Expansions      []string
	TweetFields     []string
	UserFields      []string
}

var FollowsFollowingGetQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
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

	pm := p.ParameterMap()
	qs := util.QueryString(pm, FollowsFollowingGetQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *FollowsFollowingGetParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *FollowsFollowingGetParams) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	if p.Expansions != nil && len(p.Expansions) > 0 {
		m["expansions"] = util.QueryValue(p.Expansions)
	}

	if p.TweetFields != nil && len(p.TweetFields) > 0 {
		m["tweet.fields"] = util.QueryValue(p.TweetFields)
	}

	if p.UserFields != nil && len(p.UserFields) > 0 {
		m["user.fields"] = util.QueryValue(p.UserFields)
	}

	return m
}

type FollowsFollowersParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	MaxResults      FollowsMaxResults
	PaginationToken string
	Expansions      []string
	TweetFields     []string
	UserFields      []string
}

var FollowsFollowersQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
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

	pm := p.ParameterMap()
	qs := util.QueryString(pm, FollowsFollowersQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *FollowsFollowersParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *FollowsFollowersParams) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	if p.Expansions != nil && len(p.Expansions) > 0 {
		m["expansions"] = util.QueryValue(p.Expansions)
	}

	if p.TweetFields != nil && len(p.TweetFields) > 0 {
		m["tweet.fields"] = util.QueryValue(p.TweetFields)
	}

	if p.UserFields != nil && len(p.UserFields) > 0 {
		m["user.fields"] = util.QueryValue(p.UserFields)
	}

	return m
}

type FollowsFollowingPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // The authenticated user ID

	// JSON body parameter
	TargetUserID string `json:"target_user_id"`
}

var FollowsFollowingPostQueryParams = map[string]struct{}{}

func (p *FollowsFollowingPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *FollowsFollowingPostParams) AccessToken() string {
	return p.accessToken
}

func (p *FollowsFollowingPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	return endpoint
}

func (p *FollowsFollowingPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *FollowsFollowingPostParams) ParameterMap() map[string]string {
	m := map[string]string{}
	return m
}
