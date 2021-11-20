package types

import "github.com/michimani/gotwi/resources"

type PinnedListsGetResponse struct {
	Data     []resources.List `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
}

func (r *PinnedListsGetResponse) HasPartialError() bool {
	return false
}

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
