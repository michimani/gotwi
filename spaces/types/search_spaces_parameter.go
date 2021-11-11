package types

import (
	"io"
	"strconv"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type SearchSpacesMaxResults int

type SearchSpacesParams struct {
	accessToken string

	// Path parameters
	Query       string
	Expansions  fields.ExpansionList
	MaxResults  SearchSpacesMaxResults
	SpaceFields fields.SpaceFieldList
	UserFields  fields.UserFieldList
	State       fields.State
}

var SearchSpacesQueryParams = map[string]struct{}{
	"query":        {},
	"expansions":   {},
	"space.fields": {},
	"user.fields":  {},
	"max_results":  {},
	"state":        {},
}

func (m SearchSpacesMaxResults) Valid() bool {
	return m > 0 && m <= 100
}

func (m SearchSpacesMaxResults) String() string {
	return strconv.Itoa(int(m))
}

func (p *SearchSpacesParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SearchSpacesParams) AccessToken() string {
	return p.accessToken
}

func (p *SearchSpacesParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Query == "" {
		return ""
	}

	pm := p.ParameterMap()
	qs := util.QueryString(pm, SearchSpacesQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *SearchSpacesParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *SearchSpacesParams) ParameterMap() map[string]string {
	m := map[string]string{}

	m["query"] = p.Query
	m = fields.SetFieldsParams(m, p.Expansions, p.SpaceFields, p.UserFields)

	if p.State.Valid() {
		m["state"] = p.State.String()
	}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	return m
}
