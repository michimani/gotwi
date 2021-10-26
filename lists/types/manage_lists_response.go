package types

type ManageListsPostResponse struct {
	Data struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

func (r *ManageListsPostResponse) HasPartialError() bool {
	return false
}
