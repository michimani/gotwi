package types

type ManagePinnedListPostResponse struct {
	Data struct {
		Pinned bool `json:"pinned"`
	} `json:"data"`
}

func (r *ManagePinnedListPostResponse) HasPartialError() bool {
	return false
}

type ManagePinnedListDeleteResponse struct {
	Data struct {
		Pinned bool `json:"pinned"`
	} `json:"data"`
}

func (r *ManagePinnedListDeleteResponse) HasPartialError() bool {
	return false
}
