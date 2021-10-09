package types

import "github.com/michimani/gotwi/resources"

type UserLookupUsersResponse struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *UserLookupUsersResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type UserLookupUsersIDResponse struct {
	Data     resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *UserLookupUsersIDResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type UserLookupUsersByResponse struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *UserLookupUsersByResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type UserLookupUsersByUsernameResponse struct {
	Data     resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *UserLookupUsersByUsernameResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
