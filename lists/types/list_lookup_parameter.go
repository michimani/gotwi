package types

import (
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type ListLookupIDParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	Expansions fields.ExpansionList
	ListFields fields.ListFieldList
	UserFields fields.UserFieldList
}

var ListLookupIDQueryParams = map[string]struct{}{
	"expansions":  {},
	"list.fields": {},
	"user.fields": {},
}

func (p *ListLookupIDParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListLookupIDParams) AccessToken() string {
	return p.accessToken
}

func (p *ListLookupIDParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, ListLookupIDQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *ListLookupIDParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListLookupIDParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)
	return m
}

type ListLookupOwnedListsMaxResults int

func (m ListLookupOwnedListsMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListLookupOwnedListsMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListLookupOwnedListsParams struct {
	accessToken string

	// Path parameter
	ID string // User ID

	// Query parameters
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
	MaxResults      ListLookupOwnedListsMaxResults
	PaginationToken string
}

var ListLookupOwnedListsQueryParams = map[string]struct{}{
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
	"max_results":      {},
	"pagination_token": {},
}

func (p *ListLookupOwnedListsParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListLookupOwnedListsParams) AccessToken() string {
	return p.accessToken
}

func (p *ListLookupOwnedListsParams) ResolveEndpoint(endpointBase string) string {
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

func (p *ListLookupOwnedListsParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListLookupOwnedListsParams) ParameterMap() map[string]string {
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
