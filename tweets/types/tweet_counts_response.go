package types

import "github.com/michimani/gotwi/resources"

type TweetCountsTweetsCountsRecentResponse struct {
	Data   []resources.TweetCount         `json:"data"`
	Meta   resources.TweetCountRecentMeta `json:"meta"`
	Errors []resources.PartialError       `json:"errors"`
}

func (r *TweetCountsTweetsCountsRecentResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type TweetCountsTweetsCountsAllResponse struct {
	Data   []resources.TweetCount      `json:"data"`
	Meta   resources.TweetCountAllMeta `json:"meta"`
	Errors []resources.PartialError    `json:"errors"`
}

func (r *TweetCountsTweetsCountsAllResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
