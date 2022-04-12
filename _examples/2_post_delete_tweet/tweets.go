package main

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

// SimpleTweet posts a tweet with only text, and return posted tweet ID.
func SimpleTweet(c *gotwi.Client, text string) (string, error) {
	p := &types.CreateInput{
		Text: gotwi.String(text),
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		return "", err
	}

	return gotwi.StringValue(res.Data.ID), nil
}

// DeleteTweet deletes a tweet specified by tweet ID.
func DeleteTweet(c *gotwi.Client, id string) (bool, error) {
	p := &types.DeleteInput{
		ID: id,
	}

	res, err := managetweet.Delete(context.Background(), c, p)
	if err != nil {
		return false, err
	}

	return gotwi.BoolValue(res.Data.Deleted), nil
}
