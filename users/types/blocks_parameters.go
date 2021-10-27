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

type BlocksMaxResults int

type BlocksBlockingGetParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	MaxResults      BlocksMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var BlocksBlockingGetQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (m BlocksMaxResults) Valid() bool {
	return m > 0 && m <= 1000
}

func (m BlocksMaxResults) String() string {
	return strconv.Itoa(int(m))
}

func (p *BlocksBlockingGetParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BlocksBlockingGetParams) AccessToken() string {
	return p.accessToken
}

func (p *BlocksBlockingGetParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, BlocksBlockingGetQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *BlocksBlockingGetParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *BlocksBlockingGetParams) ParameterMap() map[string]string {
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

type BlocksBlockingPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // The authenticated user ID

	// JSON body parameter
	TargetUserID *string `json:"target_user_id,omitempty"`
}

func (p *BlocksBlockingPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BlocksBlockingPostParams) AccessToken() string {
	return p.accessToken
}

func (p *BlocksBlockingPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *BlocksBlockingPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *BlocksBlockingPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type BlocksBlockingDeleteParams struct {
	accessToken string

	// Path parameters
	SourceUserID string // The authenticated user ID
	TargetUserID string // The user ID for unfollow
}

func (p *BlocksBlockingDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BlocksBlockingDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *BlocksBlockingDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.SourceUserID == "" || p.TargetUserID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.SourceUserID)
	endpoint := strings.Replace(endpointBase, ":source_user_id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.TargetUserID)
	endpoint = strings.Replace(endpoint, ":target_user_id", escapedTID, 1)

	return endpoint
}

func (p *BlocksBlockingDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *BlocksBlockingDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
