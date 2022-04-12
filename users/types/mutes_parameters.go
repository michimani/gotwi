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

// ListMutedUsersInput is struct for requesting `GET /2/users/:id/muting`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
type ListMutedUsersInput struct {
	accessToken string

	// Path parameter
	ID string // required: The authenticated user ID

	// Query parameters
	MaxResults      MutesMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var listMutedUsersQueryParameters = map[string]struct{}{
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

func (p *ListMutedUsersInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListMutedUsersInput) AccessToken() string {
	return p.accessToken
}

func (p *ListMutedUsersInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listMutedUsersQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListMutedUsersInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListMutedUsersInput) ParameterMap() map[string]string {
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

// CreateMutedUserInput is struct for requesting `POST /2/users/:id/muting`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/post-users-user_id-muting
type CreateMutedUserInput struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // required: The authenticated user ID

	// JSON body parameter
	TargetUserID string `json:"target_user_id,omitempty"` // required: target user ID to mute
}

func (p *CreateMutedUserInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *CreateMutedUserInput) AccessToken() string {
	return p.accessToken
}

func (p *CreateMutedUserInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.TargetUserID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *CreateMutedUserInput) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *CreateMutedUserInput) ParameterMap() map[string]string {
	return map[string]string{}
}

// DeleteMutedUserInput is struct for requesting `DELETE /2/users/:source_user_id/muting/:target_user_id`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/delete-users-user_id-muting
type DeleteMutedUserInput struct {
	accessToken string

	// Path parameters
	SourceUserID string // required: The authenticated user ID
	TargetUserID string // required: The user ID to unmute
}

func (p *DeleteMutedUserInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *DeleteMutedUserInput) AccessToken() string {
	return p.accessToken
}

func (p *DeleteMutedUserInput) ResolveEndpoint(endpointBase string) string {
	if p.SourceUserID == "" || p.TargetUserID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.SourceUserID)
	endpoint := strings.Replace(endpointBase, ":source_user_id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.TargetUserID)
	endpoint = strings.Replace(endpoint, ":target_user_id", escapedTID, 1)

	return endpoint
}

func (p *DeleteMutedUserInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *DeleteMutedUserInput) ParameterMap() map[string]string {
	return map[string]string{}
}
