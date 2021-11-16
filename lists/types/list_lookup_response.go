package types

import "github.com/michimani/gotwi/resources"

type ListLookupIDResponse struct {
	Data     resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes,omitempty"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *ListLookupIDResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListLookupOwnedListsResponse struct {
	Data     []resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes,omitempty"`
	Meta   resources.ListLookupOwnedListsMeta
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *ListLookupOwnedListsResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
