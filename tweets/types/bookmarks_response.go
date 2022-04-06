package types

import "github.com/michimani/gotwi/resources"

type BookmarksResponse struct {
	Data     []resources.Tweet `json:"data"`
	Meta     resources.PaginationMeta
	Includes struct {
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes,omitempty"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *BookmarksResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
