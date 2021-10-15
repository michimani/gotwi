package types

import "github.com/michimani/gotwi/resources"

type TweetCountsRecentResponse struct {
	Data   []resources.TweetCount         `json:"data"`
	Meta   resources.TweetCountRecentMeta `json:"meta"`
	Errors []resources.PartialError       `json:"errors"`
}

func (r *TweetCountsRecentResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type TweetCountsAllResponse struct {
	Data   []resources.TweetCount      `json:"data"`
	Meta   resources.TweetCountAllMeta `json:"meta"`
	Errors []resources.PartialError    `json:"errors"`
}

func (r *TweetCountsAllResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
