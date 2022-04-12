package types

import "github.com/michimani/gotwi/resources"

type ListJobsOutput struct {
	Data   []resources.Compliance   `json:"data"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListJobsOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type GetJobOutput struct {
	Data   resources.Compliance     `json:"data"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *GetJobOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type CreateJobOutput struct {
	Data   resources.Compliance     `json:"data"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *CreateJobOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
