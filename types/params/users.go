package params

type UsersExpansion string
type UsersTweetField string
type UsersUserField string

type UsersExpansions []UsersExpansion
type UsersTweetFields []UsersTweetField
type UsersUserFields []UsersUserField

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

type ByUserNameParams struct {
	// Path parameters
	UserName string

	// Query parameters
	Expansions  *UsersExpansions
	TweetFields *UsersTweetFields
	UserFields  *UsersUserFields
}

func (ue *UsersExpansions) QueryValue() string {
	v := ""
	if ue == nil {
		return v
	}
	for _, e := range *ue {
		if v != "" {
			v = v + ","
		}
		v = v + string(e)
	}

	return v
}

func (ue *UsersTweetFields) QueryValue() string {
	v := ""
	if ue == nil {
		return v
	}
	for _, e := range *ue {
		if v != "" {
			v = v + ","
		}
		v = v + string(e)
	}

	return v
}

func (ue *UsersUserFields) QueryValue() string {
	v := ""
	if ue == nil {
		return v
	}
	for _, e := range *ue {
		if v != "" {
			v = v + ","
		}
		v = v + string(e)
	}

	return v
}
