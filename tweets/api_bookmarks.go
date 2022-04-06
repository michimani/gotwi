package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	BookmarksEndpoint     = "https://api.twitter.com/2/users/:id/bookmarks"
	BookmarksPostEndpoint = "https://api.twitter.com/2/users/:id/bookmarks"
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

// Causes the user ID of an authenticated user identified in the path parameter
// to Bookmark the target Tweet provided in the request body.
// https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/post-users-id-bookmarks
func BookmarksPost(ctx context.Context, c *gotwi.GotwiClient, p *types.BookmarksPostParams) (*types.BookmarksPostResponse, error) {
	res := &types.BookmarksPostResponse{}
	if err := c.CallAPI(ctx, BookmarksPostEndpoint, "POST", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
