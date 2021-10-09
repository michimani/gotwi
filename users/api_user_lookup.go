package users

import (
	"encoding/json"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/users/types"
)

const (
	UserLookupUsersEndpoint           = "https://api.twitter.com/2/users"
	UserLookupUsersIDEndpoint         = "https://api.twitter.com/2/users/:id"
	UserLookupUsersByEndpoint         = "https://api.twitter.com/2/users/by"
	UserLookupUsersByUsernameEndpoint = "https://api.twitter.com/2/users/by/username/:username"
)

func UserLookupUsers(c *gotwi.TwitterClient, p *types.UserLookupUsersParams) (*types.UserLookupUsersResponse, error) {
	res := &types.UserLookupUsersResponse{}
	if err := execUserLookupAPI(c, UserLookupUsersEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func UserLookupUsersID(c *gotwi.TwitterClient, p *types.UserLookupUsersIDParams) (*types.UserLookupUsersIDResponse, error) {
	res := &types.UserLookupUsersIDResponse{}
	if err := execUserLookupAPI(c, UserLookupUsersIDEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func UserLookupUsersBy(c *gotwi.TwitterClient, p *types.UserLookupUsersByParams) (*types.UserLookupUsersByResponse, error) {
	res := &types.UserLookupUsersByResponse{}
	if err := execUserLookupAPI(c, UserLookupUsersByEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func UserLookupUsersByUsername(c *gotwi.TwitterClient, p *types.UserLookupUsersByUsernameParams) (*types.UserLookupUsersByUsernameResponse, error) {
	res := &types.UserLookupUsersByUsernameResponse{}
	if err := execUserLookupAPI(c, UserLookupUsersByUsernameEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func execUserLookupAPI(c *gotwi.TwitterClient, endpoint, method string, p util.Parameters, i util.Response) error {
	req, err := c.Prepare(endpoint, method, p)
	if err != nil {
		return err
	}

	res, err := c.Exec(req)
	if err != nil {
		return err
	}

	fmt.Println(string(res.Body))
	if err := json.Unmarshal(res.Body, &i); err != nil {
		return err
	}

	return nil
}
