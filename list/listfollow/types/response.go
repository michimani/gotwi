package types

import "github.com/michimani/gotwi/resources"

type ListFollowersOutput struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Meta   resources.ListFollowsFollowersMeta `json:"meta"`
	Errors []resources.PartialError           `json:"errors"`
}

func (r *ListFollowersOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListFollowedOutput struct {
	Data     []resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Meta   resources.ListFollowsFollowedListsMeta `json:"meta"`
	Errors []resources.PartialError               `json:"errors"`
}

func (r *ListFollowedOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type CreateOutput struct {
	Data struct {
		Following bool `json:"following"`
	} `json:"data"`
}

func (r *CreateOutput) HasPartialError() bool {
	return false
}

type DeleteOutput struct {
	Data struct {
		Following bool `json:"following"`
	} `json:"data"`
}

func (r *DeleteOutput) HasPartialError() bool {
	return false
}
