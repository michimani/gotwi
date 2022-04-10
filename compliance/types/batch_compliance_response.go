package types

import "github.com/michimani/gotwi/resources"

type BatchComplianceJobsResponse struct {
	Data   []resources.Compliance   `json:"data"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *BatchComplianceJobsResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type BatchComplianceJobsPostResponse struct {
	Data   resources.Compliance     `json:"data"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *BatchComplianceJobsPostResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
