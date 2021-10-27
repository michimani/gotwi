package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

type ManageListMembersPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // List ID

	// JSON body parameter
	UserID *string `json:"user_id,omitempty"`
}

func (p *ManageListMembersPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManageListMembersPostParams) AccessToken() string {
	return p.accessToken
}

func (p *ManageListMembersPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *ManageListMembersPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *ManageListMembersPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type ManageListMembersDeleteParams struct {
	accessToken string

	// Path parameter
	ID     string // List ID
	UserID string
}

func (p *ManageListMembersDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManageListMembersDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *ManageListMembersDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.UserID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedUserID := url.QueryEscape(p.UserID)
	endpoint = strings.Replace(endpoint, ":user_id", escapedUserID, 1)

	return endpoint
}

func (p *ManageListMembersDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ManageListMembersDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
