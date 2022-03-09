package types

import (
	"io"

	"github.com/michimani/gotwi/internal/util"
)

type FilteredStreamRulesGetParams struct {
	accessToken string

	// Query parameters
	IDs []string
}

var FilteredStreamRulesGetQueryParams = map[string]struct{}{
	"ids": {},
}

func (p *FilteredStreamRulesGetParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *FilteredStreamRulesGetParams) AccessToken() string {
	return p.accessToken
}

func (p *FilteredStreamRulesGetParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, FilteredStreamRulesGetQueryParams)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *FilteredStreamRulesGetParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *FilteredStreamRulesGetParams) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.IDs != nil && len(p.IDs) > 0 {
		m["ids"] = util.QueryValue(p.IDs)
	}

	return m
}
