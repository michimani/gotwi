package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/params"
	"github.com/michimani/gotwi/types/response"
)

const (
	UsersEndpoint = "https://api.twitter.com/2/users"
)

func Users(c *gotwi.TwitterClient, p *params.UsersParams) (*response.Users, error) {
	req, err := c.Prepare(UsersEndpoint, "GET", p)
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

	var r response.Users
	if err := json.Unmarshal(res.Body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
