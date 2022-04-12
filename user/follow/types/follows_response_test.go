package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/user/follow/types"
	"github.com/stretchr/testify/assert"
)

func Test_FollowsFollowingGet_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListFollowingsOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListFollowingsOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListFollowingsOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListFollowingsOutput{
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
		res    *types.ListFollowersOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListFollowersOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListFollowersOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListFollowersOutput{
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

func Test_CreateFollowingOutput_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.CreateFollowingOutput
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.CreateFollowingOutput{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.CreateFollowingOutput{
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

func Test_DeleteFollowingOutput_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.DeleteFollowingOutput
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.DeleteFollowingOutput{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.DeleteFollowingOutput{
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
