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

type ListMembersListMembershipsMaxResults int

func (m ListMembersListMembershipsMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListMembersListMembershipsMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListMembersListMembershipsParams struct {
	accessToken string

	// Path parameter
	ID string // User ID

	// Query parameters
	MaxResults      ListMembersListMembershipsMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
}

var ListMembersListMembershipsQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
}

func (p *ListMembersListMembershipsParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListMembersListMembershipsParams) AccessToken() string {
	return p.accessToken
}

func (p *ListMembersListMembershipsParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, ListMembersListMembershipsQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *ListMembersListMembershipsParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListMembersListMembershipsParams) ParameterMap() map[string]string {
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

type ListMembersGetMaxResults int

func (m ListMembersGetMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListMembersGetMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListMembersGetParams struct {
	accessToken string

	// Path parameter
	ID string // List ID

	// Query parameters
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
	MaxResults      ListMembersGetMaxResults
	PaginationToken string
}

func (p *ListMembersGetParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListMembersGetParams) AccessToken() string {
	return p.accessToken
}

func (p *ListMembersGetParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, ListLookupOwnedListsQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *ListMembersGetParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListMembersGetParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	return m
}

type ListMembersPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // List ID

	// JSON body parameter
	UserID *string `json:"user_id,omitempty"`
}

func (p *ListMembersPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListMembersPostParams) AccessToken() string {
	return p.accessToken
}

func (p *ListMembersPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *ListMembersPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *ListMembersPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type ListMembersDeleteParams struct {
	accessToken string

	// Path parameter
	ID     string // List ID
	UserID string
}

func (p *ListMembersDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListMembersDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *ListMembersDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.UserID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedUserID := url.QueryEscape(p.UserID)
	endpoint = strings.Replace(endpoint, ":user_id", escapedUserID, 1)

	return endpoint
}

func (p *ListMembersDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListMembersDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
