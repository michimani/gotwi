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
	UsersIDEndpoint = "https://api.twitter.com/2/users/:id"
)

func UsersID(c *gotwi.TwitterClient, p *params.UsersIDParams) (*response.UsersID, error) {
	req, err := c.Prepare(UsersIDEndpoint, "GET", p)
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

	var r response.UsersID
	if err := json.Unmarshal(res.Body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
