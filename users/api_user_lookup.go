package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/users/types"
)

const (
	UserLookupUsersEndpoint           = "https://api.twitter.com/2/users"
	UserLookupUsersIDEndpoint         = "https://api.twitter.com/2/users/:id"
	UserLookupUsersByEndpoint         = "https://api.twitter.com/2/users/by"
	UserLookupUsersByUsernameEndpoint = "https://api.twitter.com/2/users/by/username/:username"
)

func UserLookupUsers(c *gotwi.TwitterClient, p *types.UserLookupUsersParams) (*types.UserLookupUsersResponse, error) {
	req, err := c.Prepare(UserLookupUsersEndpoint, "GET", p)
	if err != nil {
		return nil, err
	}

	res, err := c.Exec(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("An error occered. %s", res.Status)
	}

	var r types.UserLookupUsersResponse
	if err := json.Unmarshal(res.Body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func UserLookupUsersID(c *gotwi.TwitterClient, p *types.UserLookupUsersIDParams) (*types.UserLookupUsersIDResponse, error) {
	req, err := c.Prepare(UserLookupUsersIDEndpoint, "GET", p)
	if err != nil {
		return nil, err
	}

	res, err := c.Exec(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("An error occered. %s", res.Status)
	}

	var r types.UserLookupUsersIDResponse
	if err := json.Unmarshal(res.Body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func UserLookupUsersBy(c *gotwi.TwitterClient, p *types.UserLookupUsersByParams) (*types.UserLookupUsersByResponse, error) {
	req, err := c.Prepare(UserLookupUsersByEndpoint, "GET", p)
	if err != nil {
		return nil, err
	}

	res, err := c.Exec(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("An error occered. %s", res.Status)
	}

	var r types.UserLookupUsersByResponse
	if err := json.Unmarshal(res.Body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func UserLookupUsersByUsername(c *gotwi.TwitterClient, p *types.UserLookupUsersByUsernameParams) (*types.UserLookupUsersByUsernameResponse, error) {
	req, err := c.Prepare(UserLookupUsersByUsernameEndpoint, "GET", p)
	if err != nil {
		return nil, err
	}

	res, err := c.Exec(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("An error occered. %s", res.Status)
	}

	var r types.UserLookupUsersByUsernameResponse
	if err := json.Unmarshal(res.Body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
