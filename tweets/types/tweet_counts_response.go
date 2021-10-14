package types

import "github.com/michimani/gotwi/resources"

type TweetCountsTweetsCountsRecentResponse struct {
	Data   []resources.TweetCount   `json:"data"`
	Meta   resources.TweetCountMeta `json:"meta"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *TweetCountsTweetsCountsRecentResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
