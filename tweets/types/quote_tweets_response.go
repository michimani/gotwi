package types

import "github.com/michimani/gotwi/resources"

type QuoteTweetsResponse struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users,omitempty"`
	} `json:"includes,omitempty"`
	Meta   resources.QuoteTweetsMeta `json:"meta"`
	Errors []resources.PartialError  `json:"errors,omitempty"`
}

func (r *QuoteTweetsResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
