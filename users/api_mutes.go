package users

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	MutesMutingGetEndpoint = "https://api.twitter.com/2/users/:id/muting"
)

// Returns a list of users who are muted by the specified user ID.
// https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
func MutesMutingGet(c *gotwi.GotwiClient, p *types.MutesMutingGetParams) (*types.MutesMutingGetResponse, error) {
	res := &types.MutesMutingGetResponse{}
	if err := c.CallAPI(MutesMutingGetEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}
