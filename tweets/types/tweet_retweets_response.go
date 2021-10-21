package types

import "github.com/michimani/gotwi/resources"

type TweetRetweetsRetweetedByResponse struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes,omitempty"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *TweetRetweetsRetweetedByResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type TweetRetweetsPostResponse struct {
	Data struct {
		Retweeted bool `json:"retweeted"`
	} `json:"data"`
}

func (r *TweetRetweetsPostResponse) HasPartialError() bool {
	return false
}

type TweetRetweetsDeleteResponse struct {
	Data struct {
		Retweeted bool `json:"retweeted"`
	} `json:"data"`
}

func (r *TweetRetweetsDeleteResponse) HasPartialError() bool {
	return false
}
