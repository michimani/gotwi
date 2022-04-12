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

type ListFollowersMaxResults int

func (m ListFollowersMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListFollowersMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListFollowersInput struct {
	accessToken string

	// Path parameter
	ID string // List ID

	// Query parameters
	MaxResults      ListFollowersMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var listFollowersQueryParameters = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *ListFollowersInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListFollowersInput) AccessToken() string {
	return p.accessToken
}

func (p *ListFollowersInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listFollowersQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListFollowersInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListFollowersInput) ParameterMap() map[string]string {
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

type ListFollowedMaxResults int

func (m ListFollowedMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListFollowedMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListFollowedInput struct {
	accessToken string

	// Path parameter
	ID string // User ID

	// Query parameters
	MaxResults      ListFollowedMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
}

var listFollowedQueryParameters = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
}

func (p *ListFollowedInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListFollowedInput) AccessToken() string {
	return p.accessToken
}

func (p *ListFollowedInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listFollowedQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListFollowedInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListFollowedInput) ParameterMap() map[string]string {
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

type CreateInput struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // User ID

	// JSON body parameter
	ListID string `json:"list_id"` // required
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
	ID     string // User ID
	ListID string
}

func (p *DeleteInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *DeleteInput) AccessToken() string {
	return p.accessToken
}

func (p *DeleteInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.ListID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedListID := url.QueryEscape(p.ListID)
	endpoint = strings.Replace(endpoint, ":list_id", escapedListID, 1)

	return endpoint
}

func (p *DeleteInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *DeleteInput) ParameterMap() map[string]string {
	return map[string]string{}
}
