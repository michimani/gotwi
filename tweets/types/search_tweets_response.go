package types

import "github.com/michimani/gotwi/resources"

type SearchTweetsRecentResponse struct {
	Data     []resources.Tweet        `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *SearchTweetsRecentResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type SearchTweetsAllResponse struct {
	Data     []resources.Tweet        `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *SearchTweetsAllResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
