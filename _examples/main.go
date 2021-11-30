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
	// postSampleTweet()

	// recentActivity()
}

func newOAuth1Client() (*gotwi.GotwiClient, error) {
	in := &gotwi.NewGotwiClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}

	return gotwi.NewGotwiClient(in)
}

func newOAuth2Client() (*gotwi.GotwiClient, error) {
	in2 := &gotwi.NewGotwiClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
	}

	return gotwi.NewGotwiClient(in2)
}

func postSampleTweet() {
	oauth1Client, err := newOAuth1Client()
	if err != nil {
		panic(err)
	}

	tweetText := "This is a sample tweet."
	tweetID, err := SimpleTweet(oauth1Client, tweetText)
	if err != nil {
		panic(err)
	}

	fmt.Println("Posted tweet ID is ", tweetID)
}

func recentActivity() {
	oauth1Client, err := newOAuth1Client()
	if err != nil {
		panic(err)
	}

	onlyFollowsRecentActivity(oauth1Client, "your-account-id")
}
