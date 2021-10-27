package types

type ManageListFollowsPostResponse struct {
	Data struct {
		Following bool `json:"following"`
	} `json:"data"`
}

func (r *ManageListFollowsPostResponse) HasPartialError() bool {
	return false
}

type ManageListFollowsDeleteResponse struct {
	Data struct {
		Following bool `json:"following"`
	} `json:"data"`
}

func (r *ManageListFollowsDeleteResponse) HasPartialError() bool {
	return false
}
