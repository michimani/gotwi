package types

import (
	"io"

	"github.com/michimani/gotwi/internal/util"
)

type ListRulesInput struct {
	accessToken string

	// Query parameters
	IDs []string
}

var listRulesQueryParameters = map[string]struct{}{
	"ids": {},
}

func (p *ListRulesInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListRulesInput) AccessToken() string {
	return p.accessToken
}

func (p *ListRulesInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listRulesQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListRulesInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListRulesInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.IDs != nil && len(p.IDs) > 0 {
		m["ids"] = util.QueryValue(p.IDs)
	}

	return m
}
