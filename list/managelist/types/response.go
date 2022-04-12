package types

type CreateOutput struct {
	Data struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

func (r *CreateOutput) HasPartialError() bool {
	return false
}

type UpdateOutput struct {
	Updated bool `json:"updated"`
}

func (r *UpdateOutput) HasPartialError() bool {
	return false
}

type DeleteOutput struct {
	Data struct {
		Deleted bool `json:"deleted"`
	} `json:"data"`
}

func (r *DeleteOutput) HasPartialError() bool {
	return false
}
