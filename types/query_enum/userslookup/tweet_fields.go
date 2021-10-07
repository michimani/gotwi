package userslookup

type UsersTweetField string

const (
	Attachments        UsersTweetField = "attachments"
	AuthorID           UsersTweetField = "author_id"
	ContextAnnotations UsersTweetField = "context_annotations"
	ConversationID     UsersTweetField = "conversation_id"
	TweetCreatedAt     UsersTweetField = "created_at"
	TweetEntities      UsersTweetField = "entities"
	Geo                UsersTweetField = "geo"
	TweetID            UsersTweetField = "id"
	InReplyToUserID    UsersTweetField = "in_reply_to_user_id"
	Lang               UsersTweetField = "lang"
	NonPublicMetrics   UsersTweetField = "non_public_metrics"
	TweetPublicMetrics UsersTweetField = "public_metrics"
	OrganicMetrics     UsersTweetField = "organic_metrics"
	PromotedMetrics    UsersTweetField = "promoted_metrics"
	PossiblySensitive  UsersTweetField = "possibly_sensitive"
	ReferencedTweets   UsersTweetField = "referenced_tweets"
	ReplySettings      UsersTweetField = "reply_settings"
	Source             UsersTweetField = "source"
	Text               UsersTweetField = "text"
	TweetWithheld      UsersTweetField = "withheld"
)
