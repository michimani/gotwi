package types

import "github.com/michimani/gotwi/resources"

type ListFollowsFollowersResponse struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Meta   resources.ListFollowsFollowersMeta `json:"meta"`
	Errors []resources.PartialError           `json:"errors"`
}

func (r *ListFollowsFollowersResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListFollowsFollowedListsResponse struct {
	Data     []resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Meta   resources.ListFollowsFollowedListsMeta `json:"meta"`
	Errors []resources.PartialError               `json:"errors"`
}

func (r *ListFollowsFollowedListsResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListFollowsPostResponse struct {
	Data struct {
		Following bool `json:"following"`
	} `json:"data"`
}

func (r *ListFollowsPostResponse) HasPartialError() bool {
	return false
}

type ListFollowsDeleteResponse struct {
	Data struct {
		Following bool `json:"following"`
	} `json:"data"`
}

func (r *ListFollowsDeleteResponse) HasPartialError() bool {
	return false
}
