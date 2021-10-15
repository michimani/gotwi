package resources

type PaginationMeta struct {
	ResultCount   int    `json:"result_count"`
	NextToken     string `json:"next_token,omitempty"`
	PreviousToken string `json:"previous_token,omitempty"`
}

type TweetCountRecentMeta struct {
	TotalTweetCount int `json:"total_tweet_count"`
}

type TweetCountAllMeta struct {
	TotalTweetCount int    `json:"total_tweet_count"`
	NextToken       string `json:"next_token"`
}
