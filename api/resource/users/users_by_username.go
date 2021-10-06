package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/types"
	"github.com/michimani/gotwi/types/params"
	"github.com/michimani/gotwi/types/response"
)

const (
	UsersByUserNameEndpoint = "https://api.twitter.com/2/users/by/username/:username"
)

func UsersByUserName(c *types.TwitterClient, p *params.ByUserNameParams) (*response.UsersByUsername, error) {
	if p == nil {
		return nil, fmt.Errorf("ByUserNameParams is nil")
	}

	if !c.IsReady() {
		return nil, fmt.Errorf("Twitter client is not ready")
	}

	req, err := newRequest(c, p)
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

	var r response.UsersByUsername
	if err := json.Unmarshal(res.Body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

func newRequest(c *types.TwitterClient, p *params.ByUserNameParams) (*http.Request, error) {
	ep := resolveEndpoint(p)
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	return req, nil
}

func resolveEndpoint(p *params.ByUserNameParams) string {
	encoded := url.QueryEscape(p.UserName)
	endpoint := strings.Replace(UsersByUserNameEndpoint, ":username", encoded, 1)
	query := url.Values{}
	if p.Expansions != nil {
		query.Add("expansions", p.Expansions.QueryValue())
	}

	if p.TweetFields != nil {
		query.Add("tweet.fields", p.TweetFields.QueryValue())
	}

	if p.UserFields != nil {
		query.Add("user.fields", p.UserFields.QueryValue())
	}

	if query.Has("expansions") || query.Has("tweet.fields") || query.Has("user.fields") {
		endpoint = endpoint + "?" + query.Encode()
	}

	return endpoint
}
