package fields

type Expansion string

const (
	ExpansionPinnedTweetID     Expansion = "pinned_tweet_id"
	AttachmentsPollIDs         Expansion = "attachments.poll_ids"
	AttachmentsMediaKeys       Expansion = "attachments.media_keys"
	ExpansionAuthorID          Expansion = "author_id"
	EntitiesMentionsUsername   Expansion = "entities.mentions.username"
	GeoPlaceID                 Expansion = "geo.place_id"
	ExpansionInReplyToUserID   Expansion = "in_reply_to_user_id"
	ReferencedTweetsID         Expansion = "referenced_tweets.id"
	ReferencedTweetsIDAuthorID Expansion = "referenced_tweets.id.author_id"
)
