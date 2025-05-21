package types

import "github.com/michimani/gotwi/resources"

// InitializeOutput is the output for the Initialize endpoint.
type InitializeOutput struct {
	Data   resources.UploadedMedia  `json:"data"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *InitializeOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type AppendOutput struct {
	Data struct {
		ExpiresAt int `json:"expires_at"`
	} `json:"data"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *AppendOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
