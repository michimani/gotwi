package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/media/upload/types"
)

const (
	OAuthTokenEnvKeyName       = "GOTWI_ACCESS_TOKEN"
	OAuthTokenSecretEnvKeyName = "GOTWI_ACCESS_TOKEN_SECRET"

	sampleFilePath = "sample.png"
)

func main() {
	args := os.Args
	isMulti := false
	isPost := false
	if len(args) > 1 && args[1] == "multi" {
		isMulti = true
	}

	if len(args) > 2 && args[2] == "post" {
		isPost = true
	}

	client, err := newOAuth1Client()
	if err != nil {
		panic(err)
	}

	fileBytes, err := os.ReadFile(sampleFilePath)
	if err != nil {
		panic(err)
	}

	// Initialize
	res, err := Initialize(client, &types.InitializeInput{
		MediaType:     types.MediaTypePNG,
		TotalBytes:    len(fileBytes),
		Shared:        false,
		MediaCategory: types.MediaCategoryTweetImage,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Initialized media ID: ", res.Data.MediaID)

	// Append
	mediaID := res.Data.MediaID

	if isMulti {
		// Append with some segments
		fmt.Println("Appending with some segments")
		chunkSize := len(fileBytes) / 10
		segmentIndex := 0
		for i := 0; i < len(fileBytes); i += chunkSize {
			end := min(i+chunkSize, len(fileBytes))

			chunk := fileBytes[i:end]
			appendRes, err := Append(client, &types.AppendInput{
				MediaID:      mediaID,
				Media:        bytes.NewReader(chunk),
				SegmentIndex: segmentIndex,
			})
			if err != nil {
				panic(err)
			}

			fmt.Printf("Appended segment %d response: %+v\n", segmentIndex, appendRes)
			segmentIndex++
		}
	} else {
		fmt.Println("Appending with single segment")
		appendRes, err := Append(client, &types.AppendInput{
			MediaID:      mediaID,
			Media:        bytes.NewReader(fileBytes),
			SegmentIndex: 0,
		})
		if err != nil {
			panic(err)
		}

		fmt.Printf("Appended response: %+v\n", appendRes)
	}

	// Finalize
	finalizeRes, err := Finalize(client, &types.FinalizeInput{
		MediaID: mediaID,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Finalized res: %+v\n", finalizeRes)

	// Post media
	if !isPost {
		return
	}

	postedID, err := PostWithMedia(client, "post with a media by using gotwi", mediaID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Posted ID: ", postedID)
}

func newOAuth1Client() (*gotwi.Client, error) {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}

	return gotwi.NewClient(in)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
