package types

import "github.com/michimani/gotwi/resources"

type SearchTweetsTweetsSearchRecentResponse struct {
	Data     []resources.Tweet `json:"data"`
	Meta     resources.Meta    `json:"meta"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *SearchTweetsTweetsSearchRecentResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
