package types

import "github.com/michimani/gotwi/resources"

type ListOutput struct {
	Data     []resources.User         `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type CreateOutput struct {
	Data struct {
		Blocking bool `json:"blocking"`
	} `json:"data"`
}

func (r *CreateOutput) HasPartialError() bool {
	return false
}

type DeleteOutput struct {
	Data struct {
		Blocking bool `json:"blocking"`
	} `json:"data"`
}

func (r *DeleteOutput) HasPartialError() bool {
	return false
}
