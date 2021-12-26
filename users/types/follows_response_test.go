package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/users/types"
	"github.com/stretchr/testify/assert"
)

func Test_FollowsFollowingGet_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.FollowsFollowingGetResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.FollowsFollowingGetResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.FollowsFollowingGetResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.FollowsFollowingGetResponse{
				Errors: nil},
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

func Test_FollowsFollowers_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.FollowsFollowersResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.FollowsFollowersResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.FollowsFollowersResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.FollowsFollowersResponse{
				Errors: nil},
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

func Test_FollowsFollowingPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.FollowsFollowingPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.FollowsFollowingPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.FollowsFollowingPostResponse{
				Data: struct {
					Following     bool "json:\"following\""
					PendingFollow bool "json:\"pending_follow\""
				}{
					Following:     true,
					PendingFollow: false,
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

func Test_FollowsFollowingDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.FollowsFollowingDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.FollowsFollowingDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.FollowsFollowingDeleteResponse{
				Data: struct {
					Following bool "json:\"following\""
				}{
					Following: false,
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
