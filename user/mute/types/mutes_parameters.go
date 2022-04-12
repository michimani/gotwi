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

type ListMaxResults int

// ListsInput is struct for requesting `GET /2/users/:id/muting`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
type ListsInput struct {
	accessToken string

	// Path parameter
	ID string // required: The authenticated user ID

	// Query parameters
	MaxResults      ListMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var listsQueryParameters = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (m ListMaxResults) Valid() bool {
	return m > 0 && m <= 1000
}

func (m ListMaxResults) String() string {
	return strconv.Itoa(int(m))
}

func (p *ListsInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListsInput) AccessToken() string {
	return p.accessToken
}

func (p *ListsInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listsQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListsInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListsInput) ParameterMap() map[string]string {
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

// CreateInput is struct for requesting `POST /2/users/:id/muting`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/post-users-user_id-muting
type CreateInput struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // required: The authenticated user ID

	// JSON body parameter
	TargetID string `json:"target_user_id"` // required: target user ID to mute
}

func (p *CreateInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *CreateInput) AccessToken() string {
	return p.accessToken
}

func (p *CreateInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.TargetID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *CreateInput) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *CreateInput) ParameterMap() map[string]string {
	return map[string]string{}
}

// DeleteInput is struct for requesting `DELETE /2/users/:source_user_id/muting/:target_user_id`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/delete-users-user_id-muting
type DeleteInput struct {
	accessToken string

	// Path parameters
	SourceUserID string // required: The authenticated user ID
	TargetID     string // required: The user ID to unmute
}

func (p *DeleteInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *DeleteInput) AccessToken() string {
	return p.accessToken
}

func (p *DeleteInput) ResolveEndpoint(endpointBase string) string {
	if p.SourceUserID == "" || p.TargetID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.SourceUserID)
	endpoint := strings.Replace(endpointBase, ":source_user_id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.TargetID)
	endpoint = strings.Replace(endpoint, ":target_user_id", escapedTID, 1)

	return endpoint
}

func (p *DeleteInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *DeleteInput) ParameterMap() map[string]string {
	return map[string]string{}
}
