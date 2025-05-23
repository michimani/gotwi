package tweetcount

import (
	"context"
	"errors"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/tweetcount/types"
)

const (
	listRecentEndpoint = "https://api.twitter.com/2/tweets/counts/recent"
	listAllEndpoint    = "https://api.twitter.com/2/tweets/counts/all"
)

// The recent Tweet counts endpoint returns count of Tweets from the last seven days that match a search query.
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-recent
func ListRecent(ctx context.Context, c gotwi.IClient, p *types.ListRecentInput) (*types.ListRecentOutput, error) {
	if p == nil {
		return nil, errors.New("ListRecentInput is nil")
	}
	res := &types.ListRecentOutput{}
	if err := c.CallAPI(ctx, listRecentEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

// This endpoint is only available to those users who have been approved for the Academic Research product track.
// The full-archive search endpoint returns the complete history of public Tweets matching a search query; since the first Tweet was created March 26, 2006.
// https://developer.twitter.com/en/docs/twitter-api/tweets/counts/api-reference/get-tweets-counts-all
func ListAll(ctx context.Context, c gotwi.IClient, p *types.ListAllInput) (*types.ListAllOutput, error) {
	if p == nil {
		return nil, errors.New("ListAllInput is nil")
	}
	res := &types.ListAllOutput{}
	if err := c.CallAPI(ctx, listAllEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
