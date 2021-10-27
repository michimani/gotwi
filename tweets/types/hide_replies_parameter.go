package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

type HideRepliesParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // The tweet ID to hide or unhide

	// JSON body parameter
	Hidden *bool `json:"hidden,omitempty"`
}

func (p *HideRepliesParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *HideRepliesParams) AccessToken() string {
	return p.accessToken
}

func (p *HideRepliesParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *HideRepliesParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *HideRepliesParams) ParameterMap() map[string]string {
	return map[string]string{}
}
