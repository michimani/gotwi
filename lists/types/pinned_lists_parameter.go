package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type PinnedListsGetParams struct {
	accessToken string

	// Path parameter
	ID string // User ID

	// Query parameter
	Expansions fields.ExpansionList
	ListFields fields.ListFieldList
	UserFields fields.UserFieldList
}

var PinnedListsGetQueryParams = map[string]struct{}{
	"expansions":  {},
	"list.fields": {},
	"user.fields": {},
}

func (p *PinnedListsGetParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *PinnedListsGetParams) AccessToken() string {
	return p.accessToken
}

func (p *PinnedListsGetParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, PinnedListsGetQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *PinnedListsGetParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *PinnedListsGetParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)
	return m
}

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
