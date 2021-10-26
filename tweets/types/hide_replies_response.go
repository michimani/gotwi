package types

type HideRepliesResponse struct {
	Data struct {
		Hidden bool `json:"hidden"`
	} `json:"data"`
}

func (r *HideRepliesResponse) HasPartialError() bool {
	return false
}
