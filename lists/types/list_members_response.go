package types

import "github.com/michimani/gotwi/resources"

type ListMembersListMembershipsResponse struct {
	Data     []resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Meta   resources.ListMembersListMembershipsMeta `json:"meta"`
	Errors []resources.PartialError                 `json:"errors"`
}

func (r *ListMembersListMembershipsResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListMembersGetResponse struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Meta   resources.ListMembersGetMeta `json:"meta"`
	Errors []resources.PartialError     `json:"errors"`
}

func (r *ListMembersGetResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListMembersPostResponse struct {
	Data struct {
		IsMember bool `json:"is_member"`
	} `json:"data"`
}

func (r *ListMembersPostResponse) HasPartialError() bool {
	return false
}

type ListMembersDeleteResponse struct {
	Data struct {
		IsMember bool `json:"is_member"`
	} `json:"data"`
}

func (r *ListMembersDeleteResponse) HasPartialError() bool {
	return false
}
