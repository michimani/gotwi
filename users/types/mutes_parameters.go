package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type MutesMaxResults int

type MutesMutingGetParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	MaxResults      MutesMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var MutesMutingGetQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (m MutesMaxResults) Valid() bool {
	return m > 0 && m <= 1000
}

func (m MutesMaxResults) String() string {
	return strconv.Itoa(int(m))
}

func (p *MutesMutingGetParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *MutesMutingGetParams) AccessToken() string {
	return p.accessToken
}

func (p *MutesMutingGetParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, MutesMutingGetQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *MutesMutingGetParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *MutesMutingGetParams) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.TweetFields, p.UserFields)

	return m
}

type MutesMutingPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // The authenticated user ID

	// JSON body parameter
	TargetUserID *string `json:"target_user_id,omitempty"`
}

func (p *MutesMutingPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *MutesMutingPostParams) AccessToken() string {
	return p.accessToken
}

func (p *MutesMutingPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *MutesMutingPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *MutesMutingPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type MutesMutingDeleteParams struct {
	accessToken string

	// Path parameters
	SourceUserID string // The authenticated user ID
	TargetUserID string // The user ID for unfollow
}

func (p *MutesMutingDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *MutesMutingDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *MutesMutingDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.SourceUserID == "" || p.TargetUserID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.SourceUserID)
	endpoint := strings.Replace(endpointBase, ":source_user_id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.TargetUserID)
	endpoint = strings.Replace(endpoint, ":target_user_id", escapedTID, 1)

	return endpoint
}

func (p *MutesMutingDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *MutesMutingDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
