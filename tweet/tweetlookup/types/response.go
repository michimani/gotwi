package types

import "github.com/michimani/gotwi/resources"

type ListOutput struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type GetOutput struct {
	Data     resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *GetOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
