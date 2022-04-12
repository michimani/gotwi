package types

import "github.com/michimani/gotwi/resources"

type ListOutput struct {
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

func (r *ListOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type CreateOutput struct {
	Data struct {
		Bookmarked bool `json:"bookmarked"`
	} `json:"data"`
}

func (r *CreateOutput) HasPartialError() bool {
	return false
}

type DeleteOutput struct {
	Data struct {
		Bookmarked bool `json:"bookmarked"`
	} `json:"data"`
}

func (r *DeleteOutput) HasPartialError() bool {
	return false
}
