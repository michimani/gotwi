package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

type CreateInput struct {
	accessToken string

	// JSON body parameter
	Name        string  `json:"name,"` // required
	Description *string `json:"description,omitempty"`
	Private     *bool   `json:"private,omitempty"`
}

func (p *CreateInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *CreateInput) AccessToken() string {
	return p.accessToken
}

func (p *CreateInput) ResolveEndpoint(endpointBase string) string {
	if p.Name == "" {
		return ""
	}

	return endpointBase
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

type UpdateInput struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // List ID

	// JSON body parameter
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Private     *bool   `json:"private,omitempty"`
}

func (p *UpdateInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *UpdateInput) AccessToken() string {
	return p.accessToken
}

func (p *UpdateInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *UpdateInput) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *UpdateInput) ParameterMap() map[string]string {
	return map[string]string{}
}

type DeleteInput struct {
	accessToken string

	// Path parameter
	ID string // List ID
}

func (p *DeleteInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *DeleteInput) AccessToken() string {
	return p.accessToken
}

func (p *DeleteInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *DeleteInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *DeleteInput) ParameterMap() map[string]string {
	return map[string]string{}
}
