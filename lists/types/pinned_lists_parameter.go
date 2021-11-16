package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

type PinnedListsPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // User ID

	// JSON body parameter
	ListID *string `json:"list_id,omitempty"`
}

func (p *PinnedListsPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *PinnedListsPostParams) AccessToken() string {
	return p.accessToken
}

func (p *PinnedListsPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *PinnedListsPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *PinnedListsPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type PinnedListsDeleteParams struct {
	accessToken string

	// Path parameter
	ID     string // User ID
	ListID string
}

func (p *PinnedListsDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *PinnedListsDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *PinnedListsDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.ListID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedListID := url.QueryEscape(p.ListID)
	endpoint = strings.Replace(endpoint, ":list_id", escapedListID, 1)

	return endpoint
}

func (p *PinnedListsDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *PinnedListsDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
