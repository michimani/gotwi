package types_test

import (
	"testing"

	"github.com/michimani/gotwi/lists/types"
	"github.com/michimani/gotwi/resources"
	"github.com/stretchr/testify/assert"
)

func Test_ListFollowsFollowers_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListFollowsFollowersResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListFollowsFollowersResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListFollowsFollowersResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListFollowsFollowersResponse{
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

func Test_ListFollowsFollowedLists_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListFollowsFollowedListsResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListFollowsFollowedListsResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListFollowsFollowedListsResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListFollowsFollowedListsResponse{
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

func Test_ListFollowsPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.ListFollowsPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.ListFollowsPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.ListFollowsPostResponse{
				Data: struct {
					Following bool "json:\"following\""
				}{
					Following: false,
				}},
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

func Test_ListFollowsDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.ListFollowsDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.ListFollowsDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.ListFollowsDeleteResponse{
				Data: struct {
					Following bool "json:\"following\""
				}{
					Following: false,
				}},
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
