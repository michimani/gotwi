package types

import "github.com/michimani/gotwi/resources"

type ListMembershipsOutput struct {
	Data     []resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Meta   resources.ListMembersListMembershipsMeta `json:"meta"`
	Errors []resources.PartialError                 `json:"errors"`
}

func (r *ListMembershipsOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListOutput struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Meta   resources.ListMembersGetMeta `json:"meta"`
	Errors []resources.PartialError     `json:"errors"`
}

func (r *ListOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type CreateOutput struct {
	Data struct {
		IsMember bool `json:"is_member"`
	} `json:"data"`
}

func (r *CreateOutput) HasPartialError() bool {
	return false
}

type DeleteOutput struct {
	Data struct {
		IsMember bool `json:"is_member"`
	} `json:"data"`
}

func (r *DeleteOutput) HasPartialError() bool {
	return false
}
