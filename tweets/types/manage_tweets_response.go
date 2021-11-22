package types

type ManageTweetsPostResponse struct {
	Data struct {
		ID   *string `json:"id"`
		Text *string `json:"text"`
	} `json:"data"`
}

func (r *ManageTweetsPostResponse) HasPartialError() bool {
	return false
}
