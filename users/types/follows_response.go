package types

import "github.com/michimani/gotwi/resources"

type FollowsFollowingGetResponse struct {
	Data     []resources.User `json:"data"`
	Meta     resources.Meta   `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *FollowsFollowingGetResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type FollowsFollowersResponse struct {
	Data     []resources.User `json:"data"`
	Meta     resources.Meta   `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *FollowsFollowersResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
