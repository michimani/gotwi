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
		fmt.Println("The 1st parameter is required for your account id.")
		os.Exit(1)
	}

	accountID := args[1]

	oauth1Client, err := newOAuth1Client()
	if err != nil {
		panic(err)
	}

	onlyFollowsRecentActivity(oauth1Client, accountID)
}

func newOAuth1Client() (*gotwi.Client, error) {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}

	return gotwi.NewClient(in)
}
