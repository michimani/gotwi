package tweets

import (
	"encoding/json"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	TweetLookupTweetsEndpoint = "https://api.twitter.com/2/tweets"
)

func TweetLookupTweets(c *gotwi.TwitterClient, p *types.TweetLookupTweetsParams) (*types.TweetLookupTweetsResponse, error) {
	res := &types.TweetLookupTweetsResponse{}
	if err := execTweetLookupAPI(c, TweetLookupTweetsEndpoint, "GET", p, res); err != nil {
		return nil, err
	}

	return res, nil
}

func execTweetLookupAPI(c *gotwi.TwitterClient, endpoint, method string, p util.Parameters, i util.Response) error {
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
