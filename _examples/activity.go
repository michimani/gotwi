package main

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweets"
	tweetsTypes "github.com/michimani/gotwi/tweets/types"
	"github.com/michimani/gotwi/users"
	"github.com/michimani/gotwi/users/types"
)

type twitterUser struct {
	ID       string
	Name     string
	Username string
}

func (f twitterUser) displayName() string {
	return fmt.Sprintf("%s@%s", f.Name, f.Username)
}

// onlyFollowsRecentActivity will output the accounts that are unilaterally following
// the specified user ID, along with up to three most recent tweets.
func onlyFollowsRecentActivity(c *gotwi.GotwiClient, userID string) {
	// list follows
	followings := map[string]twitterUser{}

	paginationToken := "init"
	for paginationToken != "" {
		p := &types.FollowsFollowingGetParams{
			ID:         userID,
			MaxResults: 1000,
		}

		if paginationToken != "init" && paginationToken != "" {
			p.PaginationToken = paginationToken
		}

		res, err := users.FollowsFollowingGet(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		for _, u := range res.Data {
			followings[gotwi.StringValue(u.ID)] = twitterUser{
				ID:       gotwi.StringValue(u.ID),
				Name:     gotwi.StringValue(u.Name),
				Username: gotwi.StringValue(u.Username),
			}
		}

		if res.Meta.NextToken != nil {
			paginationToken = gotwi.StringValue(res.Meta.NextToken)
		} else {
			paginationToken = ""
		}
	}

	// list followers
	followers := map[string]twitterUser{}

	paginationToken = "init"
	for paginationToken != "" {
		p := &types.FollowsFollowersParams{
			ID:         userID,
			MaxResults: 1000,
		}

		if paginationToken != "init" && paginationToken != "" {
			p.PaginationToken = paginationToken
		}

		res, err := users.FollowsFollowers(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		for _, u := range res.Data {
			followers[gotwi.StringValue(u.ID)] = twitterUser{
				ID:       gotwi.StringValue(u.ID),
				Name:     gotwi.StringValue(u.Name),
				Username: gotwi.StringValue(u.Username),
			}
		}

		if res.Meta.NextToken != nil {
			paginationToken = gotwi.StringValue(res.Meta.NextToken)
		} else {
			paginationToken = ""
		}
	}

	// only following
	onlyFollowings := map[string]twitterUser{}
	for fid, u := range followings {
		if _, ok := followers[fid]; ok {
			continue
		}

		onlyFollowings[fid] = u
	}

	// get recent tweets
	for _, onlyFollow := range onlyFollowings {
		p := &tweetsTypes.SearchTweetsRecentParams{
			MaxResults:  10,
			Query:       "from:" + onlyFollow.Username + " -is:retweet -is:reply",
			TweetFields: fields.TweetFieldList{fields.TweetFieldCreatedAt},
		}
		res, err := tweets.SearchTweetsRecent(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		fmt.Printf("----- %s's recent Tweets -----\n", onlyFollow.displayName())
		c := 0
		for _, t := range res.Data {
			if c > 3 {
				break
			}
			fmt.Printf("[%s] %s\n", t.CreatedAt, gotwi.StringValue(t.Text))
			c++
		}

		fmt.Println()
	}
}
