package types_test

import (
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetLikesLikingUsers_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.TweetLikesLikingUsersResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.TweetLikesLikingUsersResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.TweetLikesLikingUsersResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.TweetLikesLikingUsersResponse{
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

func Test_TweetLikesLikedTweets_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.TweetLikesLikedTweetsResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.TweetLikesLikedTweetsResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.TweetLikesLikedTweetsResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.TweetLikesLikedTweetsResponse{
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

func Test_TweetLikesPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.TweetLikesPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.TweetLikesPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.TweetLikesPostResponse{
				Data: struct {
					Liked bool "json:\"liked\""
				}{
					Liked: *gotwi.Bool(true),
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

func Test_TweetLikesDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.TweetLikesDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.TweetLikesDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.TweetLikesDeleteResponse{
				Data: struct {
					Liked bool "json:\"liked\""
				}{
					Liked: *gotwi.Bool(true),
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
