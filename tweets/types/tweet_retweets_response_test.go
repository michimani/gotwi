package types_test

import (
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetRetweetsRetweetedBy_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.TweetRetweetsRetweetedByResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.TweetRetweetsRetweetedByResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.TweetRetweetsRetweetedByResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.TweetRetweetsRetweetedByResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}

func Test_TweetRetweetsPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.TweetRetweetsPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.TweetRetweetsPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.TweetRetweetsPostResponse{
				Data: struct {
					Retweeted bool "json:\"retweeted\""
				}{
					Retweeted: *gotwi.Bool(true),
				},
			},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}

func Test_TweetRetweetsDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.TweetRetweetsDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.TweetRetweetsDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.TweetRetweetsDeleteResponse{
				Data: struct {
					Retweeted bool "json:\"retweeted\""
				}{
					Retweeted: *gotwi.Bool(true),
				},
			},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}
