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

type BookmarksMaxResults int

func (m BookmarksMaxResults) Valid() bool {
	return m >= 10 && m <= 100
}

func (m BookmarksMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type BookmarksParams struct {
	accessToken string

	// Path parameter
	ID string // Tweet ID

	// Query parameters
	MaxResults      BookmarksMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	MediaFields     fields.MediaFieldList
	PlaceFields     fields.PlaceFieldList
	PollFields      fields.PollFieldList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var BookmarksQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"media.fields":     {},
	"place.fields":     {},
	"poll.fields":      {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *BookmarksParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BookmarksParams) AccessToken() string {
	return p.accessToken
}

func (p *BookmarksParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, BookmarksQueryParams)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *BookmarksParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *BookmarksParams) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}

type BookmarksPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // The authenticated user ID

	// JSON body parameter
	TweetID *string `json:"tweet_id,omitempty"`
}

func (p *BookmarksPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BookmarksPostParams) AccessToken() string {
	return p.accessToken
}

func (p *BookmarksPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *BookmarksPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *BookmarksPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type BookmarksDeleteParams struct {
	accessToken string

	// Path parameter
	ID      string // The authenticated user ID
	TweetID string
}

func (p *BookmarksDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BookmarksDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *BookmarksDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.TweetID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.TweetID)
	endpoint = strings.Replace(endpoint, ":tweet_id", escapedTID, 1)

	return endpoint
}

func (p *BookmarksDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *BookmarksDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
