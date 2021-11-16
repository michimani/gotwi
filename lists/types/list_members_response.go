package types

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
