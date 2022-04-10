package compliance

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/compliance/types"
)

const (
	BatchComplianceJobsEndpoint     = "https://api.twitter.com/2/compliance/jobs"
	BatchComplianceJobsIDEndpoint   = "https://api.twitter.com/2/compliance/jobs/:id"
	BatchComplianceJobsPostEndpoint = "https://api.twitter.com/2/compliance/jobs"
)

// Returns a list of recent compliance jobs.
// https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/get-compliance-jobs
func BatchComplianceJobs(ctx context.Context, c *gotwi.GotwiClient, p *types.BatchComplianceJobsParams) (*types.BatchComplianceJobsResponse, error) {
	res := &types.BatchComplianceJobsResponse{}
	if err := c.CallAPI(ctx, BatchComplianceJobsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Get a single compliance job with the specified ID.
// https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/get-compliance-jobs-id
func BatchComplianceJobsID(ctx context.Context, c *gotwi.GotwiClient, p *types.BatchComplianceJobsIDParams) (*types.BatchComplianceJobsIDResponse, error) {
	res := &types.BatchComplianceJobsIDResponse{}
	if err := c.CallAPI(ctx, BatchComplianceJobsIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Creates a new compliance job for Tweet IDs or user IDs.
// A compliance job will contain an ID and a destination URL.
// The destination URL represents the location that contains the list of IDs consumed by your App.
// You can run one batch job at a time.
// https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/post-compliance-jobs
func BatchComplianceJobsPost(ctx context.Context, c *gotwi.GotwiClient, p *types.BatchComplianceJobsPostParams) (*types.BatchComplianceJobsPostResponse, error) {
	res := &types.BatchComplianceJobsPostResponse{}
	if err := c.CallAPI(ctx, BatchComplianceJobsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
