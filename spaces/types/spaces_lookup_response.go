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

type SpacesLookupBuyersResponse struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *SpacesLookupBuyersResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type SpacesLookupTweetsResponse struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Meta   resources.SpacesLookupTweetsMeta `json:"meta"`
	Errors []resources.PartialError         `json:"errors"`
}

func (r *SpacesLookupTweetsResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
