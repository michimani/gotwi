gotwi
===

This is a library for using the Twitter API v2 in the Go language. (It is still under development).

# Supported APIs

- Tweets
  - Tweet lookup
    - [ ] `GET /2/tweets`
    - [ ] `GET /2/tweets/:id`
  - Search Tweets
    - [ ] `GET /2/tweets/search/recent`
    - [ ] `GET /2/tweets/search/all`
  - Tweet counts
    - [ ] `GET /2/tweets/counts/recent`
    - [ ] `GET /2/tweets/counts/all`
  - Timelines
    - [ ] `GET /2/users/:id/tweets`
    - [ ] `GET /2/users/:id/mentions` 
  - Filtered stream
    - [ ] `POST /2/tweets/search/stream/rules`
    - [ ] `GET /2/tweets/search/stream/rules`
    - [ ] `GET /2/tweets/search/stream`
  - Sampled stream
    - [ ] `GET /2/tweets/sample/stream`
  - Retweets
    - [ ] `GET /2/users/:id/retweeted_by`
    - [ ] `POST /2/users/:id/retweets`
    - [ ] `DELETE /2/users/:id/retweets/:source_tweet_id`
  - Likes
    - [ ] `GET /2/tweets/:id/liking_users`
    - [ ] `GET /2/tweets/:id/liked_tweets`
    - [ ] `POST /2/users/:id/likes`
    - [ ] `DELETE /2/users/:id/likes/:tweet_id`
  - Hide replies
    - [ ] `PUT /2/tweets/:id/hidden`
- Users
  - Users lookup
    - [ ] `GET /2/users`
    - [ ] `GET /2/users/:id`
    - [ ] `GET /2/users/by`
    - [x] `GET /2/users/by/username`
  - Follows
    - [ ] `GET /2/users/:id/following`
    - [ ] `GET /2/users/:id/followers`
    - [ ] `POST /2/users/:id/following`
    - [ ] `DELETE /2/users/:source_user_id/following/:target_user_id`
  - Blocks
    - [ ] `GET /2/users/:id/blocking`
    - [ ] `POST /2/users/:id/blocking`
    - [ ] `DELETE /2/users/:source_user_id/blocking/:target_user_id`
  - Mutes
    - [ ] `GET /2/users/:id/muting`
    - [ ] `POST /2/users/:id/muting`
    - [ ] `DELETE /2/users/:source_user_id/muting/:target_user_id`
- Lists
  - Manage Lists
    - [ ] `POST /2/lists`
    - [ ] `DELETE /2/lists/:id`
    - [ ] `PUT /2/lists/:id`
  - Manage List members
    - [ ] `POST /2/lists/members`
    - [ ] `DELETE /2/lists/:id/members/:user_id`
  - Manage List follows
    - [ ] `POST /2/users/:id/followed_lists`
    - [ ] `DELETE /2/users/:id/followed_lists/:list_id`
  - Manage pinned Lists
    - [ ] `POST /2/users/:id/pinned_lists`
    - [ ] `DELETE /2/users/:id/pinned_lists/:list_id`

# Sample

## Prepare

Set the API key and API key secret to environment variables.

```
export GOTWI_API_KEY=your-api-key
export GOTWI_API_KEY_SECRET=your-api-key-secret
```

## Get a user by user name

```go
package main

import (
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/api/resource/users"
	"github.com/michimani/gotwi/types/params"
)

func main() {
	c, err := gotwi.NewAuthorizedClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &params.ByUserNameParams{
		UserName: "michimani210",
		Expansions: &params.UsersExpansions{
			params.ExpansionPinnedTweetID,
		},
		UserFields: &params.UsersUserFields{
			params.PublicMetrics,
			params.UserCreatedAt,
		},
	}

	u, err := users.UsersByUserName(c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u.Data.ID)
	fmt.Println(u.Data.Name)
	fmt.Println(u.Data.UserName)
	fmt.Println(u.Data.CreatedAt)
	fmt.Println(u.Data.PublicMetrics.FollowersCount)
	fmt.Println(u.Data.PublicMetrics.FollowingCount)
	fmt.Println(u.Data.PublicMetrics.TweetCount)
	fmt.Println(u.Includes.Tweets[0].Text)
}
```

```
go run main.go
```

You will get the following output.

```
581780917
michimani Lv.859
michimani210
2012-05-16 12:07:04 +0000 UTC
724
709
35096
pinned tweet
```