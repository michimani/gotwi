package types

type PinnedListsPostResponse struct {
	Data struct {
		Pinned bool `json:"pinned"`
	} `json:"data"`
}

func (r *PinnedListsPostResponse) HasPartialError() bool {
	return false
}

type PinnedListsDeleteResponse struct {
	Data struct {
		Pinned bool `json:"pinned"`
	} `json:"data"`
}

func (r *PinnedListsDeleteResponse) HasPartialError() bool {
	return false
}
