package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

type ManagePinnedListPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // User ID

	// JSON body parameter
	ListID *string `json:"list_id,omitempty"`
}

func (p *ManagePinnedListPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManagePinnedListPostParams) AccessToken() string {
	return p.accessToken
}

func (p *ManagePinnedListPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *ManagePinnedListPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *ManagePinnedListPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type ManagePinnedListDeleteParams struct {
	accessToken string

	// Path parameter
	ID     string // User ID
	ListID string
}

func (p *ManagePinnedListDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManagePinnedListDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *ManagePinnedListDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.ListID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedListID := url.QueryEscape(p.ListID)
	endpoint = strings.Replace(endpoint, ":list_id", escapedListID, 1)

	return endpoint
}

func (p *ManagePinnedListDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ManagePinnedListDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
