package fields

type UserField string

const (
	UserCreatedAt     UserField = "created_at"
	Description       UserField = "description"
	Entities          UserField = "entities"
	UserID            UserField = "id"
	Location          UserField = "location"
	Name              UserField = "name"
	UserPinnedTweetID UserField = "pinned_tweet_id"
	ProfileImageURL   UserField = "profile_image_url"
	Protected         UserField = "protected"
	PublicMetrics     UserField = "public_metrics"
	UserURL           UserField = "url"
	Username          UserField = "username"
	Verified          UserField = "verified"
	Withheld          UserField = "withheld"
)
