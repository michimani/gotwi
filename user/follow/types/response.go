package types

import "github.com/michimani/gotwi/resources"

type ListFollowingsOutput struct {
	Data     []resources.User         `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListFollowingsOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListFollowersOutput struct {
	Data     []resources.User         `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListFollowersOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type CreateFollowingOutput struct {
	Data struct {
		Following     bool `json:"following"`
		PendingFollow bool `json:"pending_follow"`
	} `json:"data"`
}

func (r *CreateFollowingOutput) HasPartialError() bool {
	return false
}

type DeleteFollowingOutput struct {
	Data struct {
		Following bool `json:"following"`
	} `json:"data"`
}

func (r *DeleteFollowingOutput) HasPartialError() bool {
	return false
}
