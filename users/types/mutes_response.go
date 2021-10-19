package types

import "github.com/michimani/gotwi/resources"

type MutesMutingGetResponse struct {
	Data     []resources.User         `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *MutesMutingGetResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
