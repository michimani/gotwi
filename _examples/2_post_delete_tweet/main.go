package main

import (
	"fmt"
	"os"

	"github.com/michimani/gotwi"
)

const (
	OAuthTokenEnvKeyName       = "GOTWI_ACCESS_TOKEN"
	OAuthTokenSecretEnvKeyName = "GOTWI_ACCESS_TOKEN_SECRET"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("The 1st parameter is required for command. (post|delete)")
		os.Exit(1)
	}
	cmd := args[1]

	if len(args) < 3 {
		fmt.Println("The 1st parameter is required. (tweet_text|tweet_id)")
		os.Exit(1)
	}
	p := args[2]

	oauth1Client, err := newOAuth1Client()
	if err != nil {
		panic(err)
	}

	switch cmd {
	case "post":
		tweetID, err := SimpleTweet(oauth1Client, p)
		if err != nil {
			panic(err)
		}

		fmt.Println("Posted tweet ID is ", tweetID)
	case "delete":
		b, err := DeleteTweet(oauth1Client, p)
		if err != nil {
			panic(err)
		}

		fmt.Println("Delete tweet result:  ", b)
	default:
		fmt.Println("Unsupported command. Supported commands are 'post' and 'delete'.")
	}
}

func newOAuth1Client() (*gotwi.Client, error) {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}

	return gotwi.NewClient(in)
}
