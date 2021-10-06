gotwi
===

This is a library for using the Twitter API v2 in the Go language. (It is still under development).

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