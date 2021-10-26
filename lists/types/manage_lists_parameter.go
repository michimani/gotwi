package types

import (
	"encoding/json"
	"io"
	"net/url"
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

type ManageListsPutParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // List ID

	// JSON body parameter
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Private     *bool   `json:"private,omitempty"`
}

func (p *ManageListsPutParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManageListsPutParams) AccessToken() string {
	return p.accessToken
}

func (p *ManageListsPutParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *ManageListsPutParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *ManageListsPutParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type ManageListsDeleteParams struct {
	accessToken string

	// Path parameter
	ID string // List ID
}

func (p *ManageListsDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManageListsDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *ManageListsDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *ManageListsDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ManageListsDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
