package types

import (
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type ListTweetsLookupMaxResults int

func (m ListTweetsLookupMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListTweetsLookupMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListTweetsLookupParams struct {
	accessToken string

	// Path parameter
	ID string // List ID

	// Query parameters
	MaxResults      ListTweetsLookupMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var ListTweetsLookupQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *ListTweetsLookupParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListTweetsLookupParams) AccessToken() string {
	return p.accessToken
}

func (p *ListTweetsLookupParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, ListTweetsLookupQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *ListTweetsLookupParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListTweetsLookupParams) ParameterMap() map[string]string {
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
