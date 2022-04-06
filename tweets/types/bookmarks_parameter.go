package types

import (
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