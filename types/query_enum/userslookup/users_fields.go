package userslookup

type UsersUserField string

const (
	UserCreatedAt     UsersUserField = "created_at"
	Description       UsersUserField = "description"
	Entities          UsersUserField = "entities"
	UserID            UsersUserField = "id"
	Location          UsersUserField = "location"
	Name              UsersUserField = "name"
	UserPinnedTweetID UsersUserField = "pinned_tweet_id"
	ProfileImageURL   UsersUserField = "profile_image_url"
	Protected         UsersUserField = "protected"
	PublicMetrics     UsersUserField = "public_metrics"
	URL               UsersUserField = "url"
	Username          UsersUserField = "username"
	Verified          UsersUserField = "verified"
	Withheld          UsersUserField = "withheld"
)
