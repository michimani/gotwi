package hidereply

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/hidereply/types"
)

const updateEndpoint = "https://api.twitter.com/2/tweets/:id/hidden"

// Hides or unhides a reply to a Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/hide-replies/api-reference/put-tweets-id-hidden
func Update(ctx context.Context, c *gotwi.Client, p *types.UpdateInput) (*types.UpdateOutput, error) {
	res := &types.UpdateOutput{}
	if err := c.CallAPI(ctx, updateEndpoint, "PUT", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
