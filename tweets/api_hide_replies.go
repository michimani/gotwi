package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const HideRepliesEndpoint = "https://api.twitter.com/2/tweets/:id/hidden"

// Hides or unhides a reply to a Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/hide-replies/api-reference/put-tweets-id-hidden
func HideReplies(ctx context.Context, c *gotwi.GotwiClient, p *types.HideRepliesParams) (*types.HideRepliesResponse, error) {
	res := &types.HideRepliesResponse{}
	if err := c.CallAPI(ctx, HideRepliesEndpoint, "PUT", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
