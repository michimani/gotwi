package filteredstream

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/filteredstream/types"
)

const listRulesEndpoint = "https://api.twitter.com/2/tweets/search/stream/rules"

// Return a list of rules currently active on the streaming endpoint, either as a list or individually.
// https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/api-reference/get-tweets-search-stream-rules
func ListRules(ctx context.Context, c *gotwi.Client, p *types.ListRulesInput) (*types.ListRulesOutput, error) {
	res := &types.ListRulesOutput{}
	if err := c.CallAPI(ctx, listRulesEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
