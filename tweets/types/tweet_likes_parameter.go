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

type TweetLikesLikingUsersParams struct {
	accessToken string

	// Path parameter
	ID string // Tweet ID

	// Query parameters
	Expansions  fields.ExpansionList
	MediaFields fields.MediaFieldList
	PlaceFields fields.PlaceFieldList
	PollFields  fields.PollFieldList
	TweetFields fields.TweetFieldList
	UserFields  fields.UserFieldList
}

var TweetLikesLikingUsersQueryParams = map[string]struct{}{
	"expansions":   {},
	"media.fields": {},
	"place.fields": {},
	"poll.fields":  {},
	"tweet.fields": {},
	"user.fields":  {},
}

func (p *TweetLikesLikingUsersParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetLikesLikingUsersParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetLikesLikingUsersParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, TweetLikesLikingUsersQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TweetLikesLikingUsersParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetLikesLikingUsersParams) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}

type TweetLikesMaxResults int

func (m TweetLikesMaxResults) Valid() bool {
	return m >= 10 && m <= 100
}

func (m TweetLikesMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type TweetLikesLikedTweetsParams struct {
	accessToken string

	// Path parameter
	ID string // Tweet ID

	// Query parameters
	MaxResults      TweetLikesMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	MediaFields     fields.MediaFieldList
	PlaceFields     fields.PlaceFieldList
	PollFields      fields.PollFieldList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var TweetLikesLikedTweetsQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"media.fields":     {},
	"place.fields":     {},
	"poll.fields":      {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *TweetLikesLikedTweetsParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetLikesLikedTweetsParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetLikesLikedTweetsParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, TweetLikesLikedTweetsQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *TweetLikesLikedTweetsParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetLikesLikedTweetsParams) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}

type TweetLikesPostParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // The authenticated user ID

	// JSON body parameter
	TweetID *string `json:"tweet_id,omitempty"`
}

func (p *TweetLikesPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetLikesPostParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetLikesPostParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *TweetLikesPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *TweetLikesPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type TweetLikesDeleteParams struct {
	accessToken string

	// Path parameter
	ID      string // The authenticated user ID
	TweetID string
}

func (p *TweetLikesDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *TweetLikesDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *TweetLikesDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.TweetID == "" {
		return ""
	}

	escapedSID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedSID, 1)
	escapedTID := url.QueryEscape(p.TweetID)
	endpoint = strings.Replace(endpoint, ":tweet_id", escapedTID, 1)

	return endpoint
}

func (p *TweetLikesDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *TweetLikesDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
