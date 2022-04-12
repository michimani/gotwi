package types

type UpdateOutput struct {
	Data struct {
		Hidden bool `json:"hidden"`
	} `json:"data"`
}

func (r *UpdateOutput) HasPartialError() bool {
	return false
}
