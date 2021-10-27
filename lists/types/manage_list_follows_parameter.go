package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

type ManageListFollowsPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // User ID

	// JSON body parameter
	ListID *string `json:"list_id,omitempty"`
}

func (p *ManageListFollowsPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManageListFollowsPostParams) AccessToken() string {
	return p.accessToken
}

func (p *ManageListFollowsPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *ManageListFollowsPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *ManageListFollowsPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type ManageListFollowsDeleteParams struct {
	accessToken string

	// Path parameter
	ID     string // User ID
	ListID string
}

func (p *ManageListFollowsDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManageListFollowsDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *ManageListFollowsDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.ListID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedListID := url.QueryEscape(p.ListID)
	endpoint = strings.Replace(endpoint, ":list_id", escapedListID, 1)

	return endpoint
}

func (p *ManageListFollowsDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ManageListFollowsDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
