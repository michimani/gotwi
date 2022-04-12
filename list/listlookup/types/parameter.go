package types

import (
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type GetInput struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	Expansions fields.ExpansionList
	ListFields fields.ListFieldList
	UserFields fields.UserFieldList
}

var getQueryParameters = map[string]struct{}{
	"expansions":  {},
	"list.fields": {},
	"user.fields": {},
}

func (p *GetInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *GetInput) AccessToken() string {
	return p.accessToken
}

func (p *GetInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, getQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *GetInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *GetInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)
	return m
}

type ListOwnedMaxResults int

func (m ListOwnedMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListOwnedMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListOwnedInput struct {
	accessToken string

	// Path parameter
	ID string // User ID

	// Query parameters
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
	MaxResults      ListOwnedMaxResults
	PaginationToken string
}

var listOwnedQueryParameters = map[string]struct{}{
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
	"max_results":      {},
	"pagination_token": {},
}

func (p *ListOwnedInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListOwnedInput) AccessToken() string {
	return p.accessToken
}

func (p *ListOwnedInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listOwnedQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListOwnedInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListOwnedInput) ParameterMap() map[string]string {
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
