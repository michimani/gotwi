package types

import "github.com/michimani/gotwi/resources"

type GetOutput struct {
	Data     resources.Space `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *GetOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListOutput struct {
	Data     []resources.Space `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListByCreatorIDsOutput struct {
	Data     []resources.Space `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Meta   resources.SpacesLookupByCreatorsIDsMeta `json:"meta"`
	Errors []resources.PartialError                `json:"errors"`
}

func (r *ListByCreatorIDsOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListBuyersOutput struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListBuyersOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListTweetsOutput struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Meta   resources.SpacesLookupTweetsMeta `json:"meta"`
	Errors []resources.PartialError         `json:"errors"`
}

func (r *ListTweetsOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
