package params

import (
	"io"
	"net/url"
	"strings"
)

type UsersExpansion string
type UsersTweetField string
type UsersUserField string

const (
	ExpansionPinnedTweetID UsersExpansion  = "pinned_tweet_id"
	Attachments            UsersTweetField = "attachments"
	AuthorID               UsersTweetField = "author_id"
	ContextAnnotations     UsersTweetField = "context_annotations"
	ConversationID         UsersTweetField = "conversation_id"
	TweetCreatedAt         UsersTweetField = "created_at"
	TweetEntities          UsersTweetField = "entities"
	Geo                    UsersTweetField = "geo"
	TweetID                UsersTweetField = "id"
	InReplyToUserID        UsersTweetField = "in_reply_to_user_id"
	Lang                   UsersTweetField = "lang"
	NonPublicMetrics       UsersTweetField = "non_public_metrics"
	TweetPublicMetrics     UsersTweetField = "public_metrics"
	OrganicMetrics         UsersTweetField = "organic_metrics"
	PromotedMetrics        UsersTweetField = "promoted_metrics"
	PossiblySensitive      UsersTweetField = "possibly_sensitive"
	ReferencedTweets       UsersTweetField = "referenced_tweets"
	ReplySettings          UsersTweetField = "reply_settings"
	Source                 UsersTweetField = "source"
	Text                   UsersTweetField = "text"
	TweetWithheld          UsersTweetField = "withheld"
	UserCreatedAt          UsersUserField  = "created_at"
	Description            UsersUserField  = "description"
	Entities               UsersUserField  = "entities"
	UserID                 UsersUserField  = "id"
	Location               UsersUserField  = "location"
	Name                   UsersUserField  = "name"
	UserPinnedTweetID      UsersUserField  = "pinned_tweet_id"
	ProfileImageURL        UsersUserField  = "profile_image_url"
	Protected              UsersUserField  = "protected"
	PublicMetrics          UsersUserField  = "public_metrics"
	URL                    UsersUserField  = "url"
	Username               UsersUserField  = "username"
	Verified               UsersUserField  = "verified"
	Withheld               UsersUserField  = "withheld"
)

type ByUsernameParams struct {
	accessToken string

	// Path parameters
	Username string

	// Query parameters
	Expansions  []string
	TweetFields []string
	UserFields  []string
}

func (p *ByUsernameParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ByUsernameParams) AccessToken() string {
	return p.accessToken
}

func (p *ByUsernameParams) ResolveEndpoint(endpointBase string) string {
	encoded := url.QueryEscape(p.Username)
	endpoint := strings.Replace(endpointBase, ":username", encoded, 1)

	query := url.Values{}
	if p.Expansions != nil {
		query.Add("expansions", queryValue(p.Expansions))
	}

	if p.TweetFields != nil {
		query.Add("tweet.fields", queryValue(p.TweetFields))
	}

	if p.UserFields != nil {
		query.Add("user.fields", queryValue(p.UserFields))
	}

	if p.Expansions != nil && len(p.Expansions) > 0 {
	}

	if query.Has("expansions") || query.Has("tweet.fields") || query.Has("user.fields") {
		endpoint = endpoint + "?" + query.Encode()
	}

	return endpoint
}

func queryValue(params []string) string {
	if len(params) == 0 {
		return ""
	}

	return strings.Join(params, ",")
}

func (p *ByUsernameParams) Body() io.Reader {
	return nil
}
