package types

import "github.com/michimani/gotwi/resources"

type ListRecentOutput struct {
	Data     []resources.Tweet        `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Users  []resources.User  `json:"users,omitempty"`
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListRecentOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListAllOutput struct {
	Data     []resources.Tweet        `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Users  []resources.User  `json:"users,omitempty"`
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListAllOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
