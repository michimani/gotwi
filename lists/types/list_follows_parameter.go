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

type ListFollowsFollowersMaxResults int

func (m ListFollowsFollowersMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListFollowsFollowersMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListFollowsFollowersParams struct {
	accessToken string

	// Path parameter
	ID string // List ID

	// Query parameters
	MaxResults      ListFollowsFollowersMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var ListFollowsFollowersQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *ListFollowsFollowersParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListFollowsFollowersParams) AccessToken() string {
	return p.accessToken
}

func (p *ListFollowsFollowersParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, ListFollowsFollowersQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *ListFollowsFollowersParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListFollowsFollowersParams) ParameterMap() map[string]string {
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

type ListFollowsFollowedListsMaxResults int

func (m ListFollowsFollowedListsMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListFollowsFollowedListsMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListFollowsFollowedListsParams struct {
	accessToken string

	// Path parameter
	ID string // User ID

	// Query parameters
	MaxResults      ListFollowsFollowedListsMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
}

var ListFollowsFollowedListsQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
}

func (p *ListFollowsFollowedListsParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListFollowsFollowedListsParams) AccessToken() string {
	return p.accessToken
}

func (p *ListFollowsFollowedListsParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, ListFollowsFollowedListsQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *ListFollowsFollowedListsParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListFollowsFollowedListsParams) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)

	return m
}

type ListFollowsPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // User ID

	// JSON body parameter
	ListID *string `json:"list_id,omitempty"`
}

func (p *ListFollowsPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListFollowsPostParams) AccessToken() string {
	return p.accessToken
}

func (p *ListFollowsPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *ListFollowsPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *ListFollowsPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type ListFollowsDeleteParams struct {
	accessToken string

	// Path parameter
	ID     string // User ID
	ListID string
}

func (p *ListFollowsDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListFollowsDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *ListFollowsDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.ListID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedListID := url.QueryEscape(p.ListID)
	endpoint = strings.Replace(endpoint, ":list_id", escapedListID, 1)

	return endpoint
}

func (p *ListFollowsDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListFollowsDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
