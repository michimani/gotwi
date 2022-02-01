gotwi
===

[![Go Reference](https://pkg.go.dev/badge/github.com/michimani/gotwi.svg)](https://pkg.go.dev/github.com/michimani/gotwi)
[![Twitter API v2](https://img.shields.io/endpoint?url=https%3A%2F%2Ftwbadges.glitch.me%2Fbadges%2Fv2)](https://developer.twitter.com/en/docs/twitter-api)
[![codecov](https://codecov.io/gh/michimani/gotwi/branch/main/graph/badge.svg?token=NA873TE6RV)](https://codecov.io/gh/michimani/gotwi)

This is a library for using the Twitter API v2 in the Go language. (It is still under development).

# Supported APIs

[Twitter API Documentation | Docs | Twitter Developer Platform  ](https://developer.twitter.com/en/docs/twitter-api)

Progress of supporting APIs:

- **Tweets**
  - Tweet lookup
    - [x] `GET /2/tweets`
    - [x] `GET /2/tweets/:id`
  - Manage Tweet
    - [x] `POST /2/tweets`
    - [x] `DELETE /2/tweets/:id`
  - Timelines
    - [x] `GET /2/users/:id/tweets`
    - [x] `GET /2/users/:id/mentions` 
  - Search Tweets
    - [x] `GET /2/tweets/search/recent`
    - [x] `GET /2/tweets/search/all`
  - Tweet counts
    - [x] `GET /2/tweets/counts/recent`
    - [x] `GET /2/tweets/counts/all`
  - Filtered stream
    - [ ] `POST /2/tweets/search/stream/rules`
    - [x] `GET /2/tweets/search/stream/rules`
    - [ ] `GET /2/tweets/search/stream`
  - Volume streams
    - [ ] `GET /2/tweets/sample/stream`
  - Retweets
    - [x] `GET /2/users/:id/retweeted_by`
    - [x] `POST /2/users/:id/retweets`
    - [x] `DELETE /2/users/:id/retweets/:source_tweet_id`
  - Likes
    - [x] `GET /2/tweets/:id/liking_users`
    - [x] `GET /2/tweets/:id/liked_tweets`
    - [x] `POST /2/users/:id/likes`
    - [x] `DELETE /2/users/:id/likes/:tweet_id`
  - Hide replies
    - [x] `PUT /2/tweets/:id/hidden`
- **Users**
  - User lookup
    - [x] `GET /2/users`
    - [x] `GET /2/users/:id`
    - [x] `GET /2/users/by`
    - [x] `GET /2/users/by/username`
    - [x] `GET /2/users/by/me`
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
- **Lists**
  - List lookup
    - [x] `GET /2/lists/:id`
    - [x] `GET /2/users/:id/owned_lists`
  - Manage Lists
    - [x] `POST /2/lists`
    - [x] `DELETE /2/lists/:id`
    - [x] `PUT /2/lists/:id`
  - List Tweets lookup
    - [x] `GET /2/lists/:id/tweets`
  - List members
    - [x] `GET /2/users/:id/list_memberships`
    - [x] `GET /2/lists/:id/members`
    - [x] `POST /2/lists/:id/members`
    - [x] `DELETE /2/lists/:id/members/:user_id`
  - List follows
    - [x] `GET /2/lists/:id/followers`
    - [x] `GET /2/users/:id/followed_lists`
    - [x] `POST /2/users/:id/followed_lists`
    - [x] `DELETE /2/users/:id/followed_lists/:list_id`
  - Pinned Lists
    - [x] `GET /2/users/:id/pinned_lists`
    - [x] `POST /2/users/:id/pinned_lists`
    - [x] `DELETE /2/users/:id/pinned_lists/:list_id`
- **Spaces**
  - Spaces Lookup
    - [x] `GET /2/spaces/:id`
    - [x] `GET /2/spaces`
    - [x] `GET /2/spaces/by/creator_ids`
    - [x] `GET /2/spaces/:id/buyers`
    - [x] `GET /2/spaces/:id/tweets`
  - Search Spaces
    - [x] `GET /2/spaces/search`
- **Compliance**
  - Batch compliance
    - [ ] `GET /2/compliance/jobs/:id`
    - [ ] `GET /2/compliance/jobs`
    - [ ] `POST /2/compliance/jobs`

# Sample

## Prepare

Set the API key and API key secret to environment variables.

```
export GOTWI_API_KEY=your-api-key
export GOTWI_API_KEY_SECRET=your-api-key-secret
```

## Request with OAuth 2.0 Bearer Token

This authentication method allows only read-only access to public information.

### Example: Get a user by user name.

```go
package main

import (
	"context"
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
		Expansions: fields.ExpansionList{
			fields.ExpansionPinnedTweetID,
		},
		UserFields: fields.UserFieldList{
			fields.UserFieldCreatedAt,
		},
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldCreatedAt,
		},
	}
	res, err := users.UserLookupByUsername(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID: ", gotwi.StringValue(u.Data.ID))
	fmt.Println("Name: ", gotwi.StringValue(u.Data.Name))
	fmt.Println("Username: ", gotwi.StringValue(u.Data.Username))
	fmt.Println("CreatedAt: ", u.Data.CreatedAt)
	if u.Includes.Tweets != nil {
		for _, t := range u.Includes.Tweets {
			fmt.Println("PinnedTweet: ", gotwi.StringValue(t.Text))
		}
	}
}
```

```
go run main.go
```

You will get the output like following.

```
ID:  581780917
Name:  michimani Lv.859
Username:  michimani210
CreatedAt:  2012-05-16 12:07:04 +0000 UTC
PinnedTweet:  pinned tweet
```

## Request with OAuth 1.0a User Context

With this authentication method, each operation will be performed as the authenticated Twitter account. For example, you can tweet as that account, or retrieve accounts that are blocked by that account.

### Example: Tweet with poll.

```go
package main

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets"
	"github.com/michimani/gotwi/tweets/types"
)

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

	p := &types.ManageTweetsPostParams{
		Text: gotwi.String("This is a test tweet with poll."),
		Poll: &types.ManageTweetsPostParamsPoll{
			DurationMinutes: gotwi.Int(5),
			Options: []string{
				"Cyan",
				"Magenta",
				"Yellow",
				"Key plate",
			},
		},
	}

	res, err := tweets.ManageTweetsPost(context.Background(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
}
```

```
go run main.go
```

You will get the output like following.

```
[1462813519607263236] This is a test tweet with poll.
```

## Error handling

Each function that calls the Twitter API (e.g. `TweetRetweetsRetweetedBy()`) may return an error for some reason.
If the error is caused by the Twitter API returning a status other than 2XX, you can check the details by doing the following.

```go
res, err := tweets.TweetRetweetsRetweetedBy(context.Background(), c, p)
if err != nil {
	fmt.Println(err)

	// more error information
	ge := err.(*gotwi.GotwiError)
	if ge.OnAPI {
		fmt.Println(ge.Title)
		fmt.Println(ge.Detail)
		fmt.Println(ge.Type)
		fmt.Println(ge.Status)
		fmt.Println(ge.StatusCode)

		for _, ae := range ge.APIErrors {
			fmt.Println(ae.Message)
			fmt.Println(ae.Label)
			fmt.Println(ae.Parameters)
			fmt.Println(ae.Code)
			fmt.Println(ae.Code.Detail())
		}

		if ge.RateLimitInfo != nil {
			fmt.Println(ge.RateLimitInfo.Limit)
			fmt.Println(ge.RateLimitInfo.Remaining)
			fmt.Println(ge.RateLimitInfo.ResetAt)
		}
	}
}
```



## More examples

See [_examples](https://github.com/michimani/gotwi/tree/main/_examples) directory.

# Licence

[MIT](https://github.com/michimani/gotwi/blob/main/LICENCE)

# Author

[michimani210](https://twitter.com/michimani210)

