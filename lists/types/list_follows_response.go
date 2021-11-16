package types

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
