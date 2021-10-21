gotwi
===

This is a library for using the Twitter API v2 in the Go language. (It is still under development).

# Supported APIs

[What's New with Twitter API v2 | Docs | Twitter Developer Platform](https://developer.twitter.com/en/docs/twitter-api/early-access)

Progress of supporting API: 26/43

- **Tweets** (12/20)
  - Tweet lookup
    - [x] `GET /2/tweets`
    - [x] `GET /2/tweets/:id`
  - Search Tweets
    - [x] `GET /2/tweets/search/recent`
    - [x] `GET /2/tweets/search/all`
  - Tweet counts
    - [x] `GET /2/tweets/counts/recent`
    - [x] `GET /2/tweets/counts/all`
  - Timelines
    - [x] `GET /2/users/:id/tweets`
    - [x] `GET /2/users/:id/mentions` 
  - Filtered stream
    - [ ] `POST /2/tweets/search/stream/rules`
    - [ ] `GET /2/tweets/search/stream/rules`
    - [ ] `GET /2/tweets/search/stream`
  - Sampled stream
    - [ ] `GET /2/tweets/sample/stream`
  - Retweets
    - [x] `GET /2/users/:id/retweeted_by`
    - [ ] `POST /2/users/:id/retweets`
    - [ ] `DELETE /2/users/:id/retweets/:source_tweet_id`
  - Likes
    - [x] `GET /2/tweets/:id/liking_users`
    - [x] `GET /2/tweets/:id/liked_tweets`
    - [x] `POST /2/users/:id/likes`
    - [ ] `DELETE /2/users/:id/likes/:tweet_id`
  - Hide replies
    - [ ] `PUT /2/tweets/:id/hidden`
- ✅  **Users** (14/14)
  - User lookup
    - [x] `GET /2/users`
    - [x] `GET /2/users/:id`
    - [x] `GET /2/users/by`
    - [x] `GET /2/users/by/username`
  - Follows
    - [x] `GET /2/users/:id/following`
    - [x] `GET /2/users/:id/followers`
    - [x] `POST /2/users/:id/following`
    - [x] `DELETE /2/users/:source_user_id/following/:target_user_id`
  - Blocks
    - [x] `GET /2/users/:id/blocking`
    - [x] `POST /2/users/:id/blocking`
    - [x] `DELETE /2/users/:source_user_id/blocking/:target_user_id`
  - Mutes
    - [x] `GET /2/users/:id/muting`
    - [x] `POST /2/users/:id/muting`
    - [x] `DELETE /2/users/:source_user_id/muting/:target_user_id`
- **Lists** (0/9)
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

## Request with OAuth 2.0 Bearer Token

- Get a user by user name.

```go
package main

import (
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/users"
	"github.com/michimani/gotwi/users/types"
)

func main() {
	in := &gotwi.NewGotwiClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
	}

	c, err := gotwi.NewGotwiClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.UserLookupByUsernameParams{
		Username: "michimani210",
		Expansions: []string{
			string(fields.UserPinnedTweetID),
		},
		UserFields: []string{
			string(fields.UserCreatedAt),
		},
		TweetFields: []string{
			string(fields.TweetCreatedAt),
		},
	}
	res, err := users.UserLookupByUsername(c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID: ", res.Data.ID)
	fmt.Println("Name: ", res.Data.Name)
	fmt.Println("Username: ", res.Data.Username)
	fmt.Println("CreatedAt: ", res.Data.CreatedAt)
	for _, t := range res.Includes.Tweets {
		fmt.Println("PinnedTweet: ", t.Text)
	}
}
```

```
go run main.go
```

You will get the following output.

```
ID:  581780917
Name:  michimani Lv.859
Username:  michimani210
CreatedAt:  2012-05-16 12:07:04 +0000 UTC
PinnedTweet:  pinned tweet
```

## Request with OAuth 1.0a User Context

- Get blocking users.

```go
func main() {
	in := &gotwi.NewGotwiClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           "your-twitter-acount-oauth-token",
		OAuthTokenSecret:     "your-twitter-acount-oauth-token-secret",
	}

	c, err := gotwi.NewGotwiClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &types.BlocksBlockingGetParams{
		ID:        "your-twitter-acount-id",
		MaxResults: 5,
	}

	res, err := users.BlocksBlockingGet(c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, b := range res.Data {
		fmt.Println(b.Name)
	}
}
```

```
go run main.go
```

You will get the output like following.

```
blockingUser1
blockingUser2
blockingUser3
blockingUser4
blockingUser5
```