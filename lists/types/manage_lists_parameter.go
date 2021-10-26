package types

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/michimani/gotwi"
)

type ManageListsPostParams struct {
	accessToken string

	// JSON body parameter
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Private     *bool   `json:"private,omitempty"`
}

func (p *ManageListsPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManageListsPostParams) AccessToken() string {
	return p.accessToken
}

func (p *ManageListsPostParams) ResolveEndpoint(endpointBase string) string {
	if p.Name == nil || gotwi.StringValue(p.Name) == "" {
		return ""
	}

	return endpointBase
}

func (p *ManageListsPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *ManageListsPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}
