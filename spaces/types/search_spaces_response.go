package types

import "github.com/michimani/gotwi/resources"

type SearchSpacesResponse struct {
	Data     []resources.Space `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *SearchSpacesResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
