package batchcompliance

import (
	"context"
	"errors"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/compliance/batchcompliance/types"
)

const (
	listJobsEndpoint  = "https://api.twitter.com/2/compliance/jobs"
	GetJobEndpoint    = "https://api.twitter.com/2/compliance/jobs/:id"
	createJobEndpoint = "https://api.twitter.com/2/compliance/jobs"
)

// Returns a list of recent compliance jobs.
// https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/get-compliance-jobs
func ListJobs(ctx context.Context, c gotwi.IClient, p *types.ListJobsInput) (*types.ListJobsOutput, error) {
	if p == nil {
		return nil, errors.New("ListJobsInput is nil")
	}
	res := &types.ListJobsOutput{}
	if err := c.CallAPI(ctx, listJobsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Get a single compliance job with the specified ID.
// https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/get-compliance-jobs-id
func GetJob(ctx context.Context, c gotwi.IClient, p *types.GetJobInput) (*types.GetJobOutput, error) {
	if p == nil {
		return nil, errors.New("GetJobInput is nil")
	}
	res := &types.GetJobOutput{}
	if err := c.CallAPI(ctx, GetJobEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Creates a new compliance job for Tweet IDs or user IDs.
// A compliance job will contain an ID and a destination URL.
// The destination URL represents the location that contains the list of IDs consumed by your App.
// You can run one batch job at a time.
// https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/api-reference/post-compliance-jobs
func CreateJob(ctx context.Context, c gotwi.IClient, p *types.CreateJobInput) (*types.CreateJobOutput, error) {
	if p == nil {
		return nil, errors.New("CreateJobInput is nil")
	}
	res := &types.CreateJobOutput{}
	if err := c.CallAPI(ctx, createJobEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
