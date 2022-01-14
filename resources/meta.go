package resources

import "time"

type PaginationMeta struct {
	ResultCount   *int    `json:"result_count"`
	NextToken     *string `json:"next_token,omitempty"`
	PreviousToken *string `json:"previous_token,omitempty"`
}

type TweetCountRecentMeta struct {
	TotalTweetCount *int `json:"total_tweet_count"`
}

type TweetCountAllMeta struct {
	TotalTweetCount *int    `json:"total_tweet_count"`
	NextToken       *string `json:"next_token"`
}

type TweetTimelineMeta struct {
	Count         *int    `json:"count"`
	NewestID      *string `json:"newest_id"`
	OldestID      *string `json:"oldest_id"`
	NextToken     *string `json:"next_token"`
	PreviousToken *string `json:"previous_token"`
}

type SpacesLookupByCreatorsIDsMeta struct {
	ResultCount *int `json:"result_count"`
}

type SpacesLookupTweetsMeta struct {
	ResultCount *int `json:"result_count"`
}

type FilterdStreamRulesGetMeta struct {
	Sent *time.Time `json:"sent"`
}

type ListLookupOwnedListsMeta struct {
	ResultCount   *int    `json:"result_count"`
	NextToken     *string `json:"next_token,omitempty"`
	PreviousToken *string `json:"previous_token,omitempty"`
}

type ListMembersListMembershipsMeta struct {
	ResultCount   *int    `json:"result_count"`
	NextToken     *string `json:"next_token,omitempty"`
	PreviousToken *string `json:"previous_token,omitempty"`
}

type ListMembersGetMeta struct {
	ResultCount   *int    `json:"result_count"`
	NextToken     *string `json:"next_token,omitempty"`
	PreviousToken *string `json:"previous_token,omitempty"`
}

type ListTweetsLookupMeta struct {
	ResultCount   *int    `json:"result_count"`
	NextToken     *string `json:"next_token,omitempty"`
	PreviousToken *string `json:"previous_token,omitempty"`
}

type ListFollowsFollowersMeta struct {
	ResultCount   *int    `json:"result_count"`
	NextToken     *string `json:"next_token,omitempty"`
	PreviousToken *string `json:"previous_token,omitempty"`
}

type ListFollowsFollowedListsMeta struct {
	ResultCount   *int    `json:"result_count"`
	NextToken     *string `json:"next_token,omitempty"`
	PreviousToken *string `json:"previous_token,omitempty"`
}
