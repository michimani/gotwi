package resources

import "time"

type TweetCount struct {
	End        *time.Time `json:"end"`
	Start      *time.Time `json:"start"`
	TweetCount int        `json:"tweet_count"`
}

type TweetCountRecentMeta struct {
	TotalTweetCount int `json:"total_tweet_count"`
}

type TweetCountAllMeta struct {
	TotalTweetCount int    `json:"total_tweet_count"`
	NextToken       string `json:"next_token"`
}
