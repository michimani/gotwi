package tweets

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetLikesLikingUsersEndpoint = "https://api.twitter.com/2/tweets/:id/liking_users"
)

// Allows you to get information about a Tweet’s liking users.
// You will receive the most recent 100 users who liked the specified Tweet.
// https://developer.twitter.com/en/docs/twitter-api/tweets/likes/api-reference/get-tweets-id-liking_users
func TweetLikesLikingUsers(c *gotwi.GotwiClient, p *types.TweetLikesLikingUsersParams) (*types.TweetLikesLikingUsersResponse, error) {
	res := &types.TweetLikesLikingUsersResponse{}
	if err := c.CallAPI(TweetLikesLikingUsersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
