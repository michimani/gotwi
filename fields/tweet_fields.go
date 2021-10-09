package fields

type TweetField string

const (
	Attachments          TweetField = "attachments"
	TweetAuthorID        TweetField = "author_id"
	ContextAnnotations   TweetField = "context_annotations"
	ConversationID       TweetField = "conversation_id"
	TweetCreatedAt       TweetField = "created_at"
	TweetEntities        TweetField = "entities"
	TweetGeo             TweetField = "geo"
	TweetID              TweetField = "id"
	TweetInReplyToUserID TweetField = "in_reply_to_user_id"
	Lang                 TweetField = "lang"
	NonPublicMetrics     TweetField = "non_public_metrics"
	TweetPublicMetrics   TweetField = "public_metrics"
	OrganicMetrics       TweetField = "organic_metrics"
	PromotedMetrics      TweetField = "promoted_metrics"
	PossiblySensitive    TweetField = "possibly_sensitive"
	ReferencedTweets     TweetField = "referenced_tweets"
	ReplySettings        TweetField = "reply_settings"
	Source               TweetField = "source"
	Text                 TweetField = "text"
	TweetWithheld        TweetField = "withheld"
)
