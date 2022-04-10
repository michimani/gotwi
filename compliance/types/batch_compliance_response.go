package types

import "github.com/michimani/gotwi/resources"

type BatchComplianceJobsPostResponse struct {
	Data   resources.Compliance     `json:"data"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *BatchComplianceJobsPostResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
