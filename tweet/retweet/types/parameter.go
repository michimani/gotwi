package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type ListUsersMaxResults int

func (m ListUsersMaxResults) Valid() bool {
	return m > 0 && m <= 100
}

func (m ListUsersMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListUsersInput struct {
	accessToken string

	// Path parameter
	ID string // Tweet ID

	// Query parameters
	Expansions      fields.ExpansionList
	MaxResults      ListUsersMaxResults // default 100
	PaginationToken string
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var listUsersQueryParameters = map[string]struct{}{
	"id":               {},
	"expansions":       {},
	"max_results":      {},
	"pagination_token": {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *ListUsersInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListUsersInput) AccessToken() string {
	return p.accessToken
}

func (p *ListUsersInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listUsersQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListUsersInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListUsersInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

type CreateInput struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // The authenticated user ID

	// JSON body parameter
	TweetID string `json:"tweet_id"` // required
}

func (p *CreateInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *CreateInput) AccessToken() string {
	return p.accessToken
}

func (p *CreateInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *CreateInput) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *CreateInput) ParameterMap() map[string]string {
	return map[string]string{}
}

type DeleteInput struct {
	accessToken string

	// Path parameter
	ID            string // The authenticated user ID
	SourceTweetID string
}

func (p *DeleteInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *DeleteInput) AccessToken() string {
	return p.accessToken
}

func (p *DeleteInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.SourceTweetID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.SourceTweetID)
	endpoint = strings.Replace(endpoint, ":source_tweet_id", escapedTID, 1)

	return endpoint
}

func (p *DeleteInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *DeleteInput) ParameterMap() map[string]string {
	return map[string]string{}
}
