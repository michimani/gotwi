package types

import "github.com/michimani/gotwi/resources"

type ListTweetsLookupResponse struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users  []resources.User  `json:"users"`
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Meta   resources.ListTweetsLookupMeta `json:"meta"`
	Errors []resources.PartialError       `json:"errors"`
}

func (r *ListTweetsLookupResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
