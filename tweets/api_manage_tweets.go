package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const ManageTweetsPostEndpoint = "https://api.twitter.com/2/tweets"

// Creates a Tweet on behalf of an authenticated user.
// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets
func ManageTweetsPost(ctx context.Context, c *gotwi.GotwiClient, p *types.ManageTweetsPostParams) (*types.ManageTweetsPostResponse, error) {
	res := &types.ManageTweetsPostResponse{}
	if err := c.CallAPI(ctx, ManageTweetsPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
