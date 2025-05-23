package quotetweet

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/quotetweet/types"
)

const (
	listEndpoint = "https://api.twitter.com/2/tweets/:id/quote_tweets"
)

// Returns Quote Tweets for a Tweet specified by the requested Tweet ID.
// https://developer.twitter.com/en/docs/twitter-api/tweets/quote-tweets/api-reference/get-tweets-id-quote_tweets
func List(ctx context.Context, c gotwi.IClient, p *types.ListInput) (*types.ListOutput, error) {
	if p == nil {
		return nil, fmt.Errorf("parameters is required")
	}

	res := &types.ListOutput{}
	if err := c.CallAPI(ctx, listEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
