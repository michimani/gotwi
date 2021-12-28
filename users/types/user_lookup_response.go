package types

import "github.com/michimani/gotwi/resources"

type UserLookupResponse struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *UserLookupResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type UserLookupIDResponse struct {
	Data     resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *UserLookupIDResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type UserLookupByResponse struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *UserLookupByResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type UserLookupByUsernameResponse struct {
	Data     resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *UserLookupByUsernameResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type UserLookupMeResponse struct {
	Data     resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *UserLookupMeResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
