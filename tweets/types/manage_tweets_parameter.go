package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

type ManageTweetsPostParams struct {
	accessToken string

	// JSON body parameter
	DirectMessageDeepLink *string                      `json:"direct_message_deep_link,omitempty"`
	ForSuperFollowersOnly *bool                        `json:"for_super_followers_only,omitempty"`
	Geo                   *ManageTweetsPostParamsGeo   `json:"geo,omitempty"`
	Media                 *ManageTweetsPostParamsMedia `json:"media,omitempty"`
	Poll                  *ManageTweetsPostParamsPoll  `json:"poll,omitempty"`
	Reply                 *ManageTweetsPostParamsReply `json:"reply,omitempty"`
	ReplySettings         *string                      `json:"reply_settings,omitempty"`
	Text                  *string                      `json:"text,omitempty"`
}

type ManageTweetsPostParamsGeo struct {
	PlaceID *string `json:"place_id,omitempty"`
}

type ManageTweetsPostParamsMedia struct {
	MediaIDs     []string `json:"media_ids,omitempty"`
	TaggedUserID *string  `json:"tagged_user_ids,omitempty"`
}

type ManageTweetsPostParamsPoll struct {
	DurationMinutes *int     `json:"duration_minutes,omitempty"`
	Options         []string `json:"options,omitempty"`
}

type ManageTweetsPostParamsReply struct {
	ExcludeReplyUserIDs []string `json:"exclude_reply_user_ids,omitempty"`
	InReplyToTweetID    string   `json:"in_reply_to_tweet_id,omitempty"`
}

func (p *ManageTweetsPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManageTweetsPostParams) AccessToken() string {
	return p.accessToken
}

func (p *ManageTweetsPostParams) ResolveEndpoint(endpointBase string) string {
	return endpointBase
}

func (p *ManageTweetsPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *ManageTweetsPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}

type ManageTweetsDeleteParams struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // The tweet ID to delete
}

func (p *ManageTweetsDeleteParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ManageTweetsDeleteParams) AccessToken() string {
	return p.accessToken
}

func (p *ManageTweetsDeleteParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *ManageTweetsDeleteParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ManageTweetsDeleteParams) ParameterMap() map[string]string {
	return map[string]string{}
}
