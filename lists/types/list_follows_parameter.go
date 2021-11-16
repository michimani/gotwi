package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

type ListFollowsPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // User ID

	// JSON body parameter
	ListID *string `json:"list_id,omitempty"`
}

func (p *ListFollowsPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListFollowsPostParams) AccessToken() string {
	return p.accessToken
}

func (p *ListFollowsPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *ListFollowsPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *ListFollowsPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type ListFollowsDeleteParams struct {
	accessToken string

	// Path parameter
	ID     string // User ID
	ListID string
}

func (p *ListFollowsDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListFollowsDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *ListFollowsDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.ListID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedListID := url.QueryEscape(p.ListID)
	endpoint = strings.Replace(endpoint, ":list_id", escapedListID, 1)

	return endpoint
}

func (p *ListFollowsDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListFollowsDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
