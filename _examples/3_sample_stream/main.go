package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/michimani/gotwi"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("The 1st parameter is required for sampling count.")
		os.Exit(1)
	}

	count, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	client, err := newOAuth2Client()
	if err != nil {
		panic(err)
	}

	samplingTweets(client, count)
}

func newOAuth2Client() (*gotwi.Client, error) {
	in2 := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth2BearerToken,
	}

	return gotwi.NewClient(in2)
}
