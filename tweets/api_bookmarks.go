package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	BookmarksEndpoint = "https://api.twitter.com/2/users/:id/bookmarks"
)

// Allows you to get information about a authenticated user’s 800 most recent bookmarked Tweets
// https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/get-users-id-bookmarks
func Bookmarks(ctx context.Context, c *gotwi.GotwiClient, p *types.BookmarksParams) (*types.BookmarksResponse, error) {
	res := &types.BookmarksResponse{}
	if err := c.CallAPI(ctx, BookmarksEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
