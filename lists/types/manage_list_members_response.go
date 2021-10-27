package types

type ManageListMembersPostResponse struct {
	Data struct {
		IsMember bool `json:"is_member"`
	} `json:"data"`
}

func (r *ManageListMembersPostResponse) HasPartialError() bool {
	return false
}

type ManageListMembersDeleteResponse struct {
	Data struct {
		IsMember bool `json:"is_member"`
	} `json:"data"`
}

func (r *ManageListMembersDeleteResponse) HasPartialError() bool {
	return false
}
