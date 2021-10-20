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

type MutesMutingPostResponse struct {
	Data struct {
		Muting bool `json:"muting"`
	} `json:"data"`
}

func (r *MutesMutingPostResponse) HasPartialError() bool {
	return false
}

type MutesMutingDeleteResponse struct {
	Data struct {
		Muting bool `json:"muting"`
	} `json:"data"`
}

func (r *MutesMutingDeleteResponse) HasPartialError() bool {
	return false
}
