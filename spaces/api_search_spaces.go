package spaces

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/spaces/types"
)

const (
	SearchSpacesEndpoint = "https://api.twitter.com/2/spaces/search"
)

// Return live or scheduled Spaces matching your specified search terms.
// This endpoint performs a keyword search, meaning that it will return Spaces
// that are an exact case-insensitive match of the specified search term. The search term will match the original title of the Space.
// https://developer.twitter.com/en/docs/twitter-api/spaces/search/api-reference/get-spaces-search
func SearchSpaces(ctx context.Context, c *gotwi.GotwiClient, p *types.SearchSpacesParams) (*types.SearchSpacesResponse, error) {
	res := &types.SearchSpacesResponse{}
	if err := c.CallAPI(ctx, SearchSpacesEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
