package tweets

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
)

const FilteredStreamRulesGetEndpoint = "https://api.twitter.com/2/tweets/search/stream/rules"

// Return a list of rules currently active on the streaming endpoint, either as a list or individually.
// https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream-rules
func FilteredStreamRulesGet(ctx context.Context, c *gotwi.GotwiClient, p *types.FilteredStreamRulesGetParams) (*types.FilteredStreamRulesGetResponse, error) {
	res := &types.FilteredStreamRulesGetResponse{}
	if err := c.CallAPI(ctx, FilteredStreamRulesGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
