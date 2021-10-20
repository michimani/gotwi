package types

import "github.com/michimani/gotwi/resources"

type BlocksBlockingGetResponse struct {
	Data     []resources.User         `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *BlocksBlockingGetResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type BlocksBlockingPostResponse struct {
	Data struct {
		Blocking bool `json:"blocking"`
	} `json:"data"`
}

func (r *BlocksBlockingPostResponse) HasPartialError() bool {
	return false
}

type BlocksBlockingDeleteResponse struct {
	Data struct {
		Blocking bool `json:"blocking"`
	} `json:"data"`
}

func (r *BlocksBlockingDeleteResponse) HasPartialError() bool {
	return false
}
