package types

import "github.com/michimani/gotwi/resources"

type SpacesLookupIDResponse struct {
	Data     resources.Space `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *SpacesLookupIDResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type SpacesLookupResponse struct {
	Data     []resources.Space `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *SpacesLookupResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type SpacesLookupByCreatorIDsResponse struct {
	Data     []resources.Space `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Meta   resources.SpacesLookupByCreatorsIDsMeta `json:"meta"`
	Errors []resources.PartialError                `json:"errors"`
}

func (r *SpacesLookupByCreatorIDsResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
