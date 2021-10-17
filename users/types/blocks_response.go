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
