package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	QuoteTweetsEndpoint = "https://api.twitter.com/2/tweets/:id/quote_tweets"
)

// Returns Quote Tweets for a Tweet specified by the requested Tweet ID.
// https://developer.twitter.com/en/docs/twitter-api/tweets/quote-tweets/api-reference/get-tweets-id-quote_tweets
func QuoteTweets(ctx context.Context, c *gotwi.GotwiClient, p *types.QuoteTweetsParams) (*types.QuoteTweetsResponse, error) {
	res := &types.QuoteTweetsResponse{}
	if err := c.CallAPI(ctx, QuoteTweetsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
