package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

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
