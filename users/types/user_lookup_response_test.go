package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/users/types"
	"github.com/stretchr/testify/assert"
)

func Test_UserLookupUsers_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.UserLookupUsersResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.UserLookupUsersResponse{
				Errors: []resources.PartialError{
					{Title: "test partical error"},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.UserLookupUsersResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.UserLookupUsersResponse{
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

func Test_UserLookupUsersID_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.UserLookupUsersIDResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.UserLookupUsersIDResponse{
				Errors: []resources.PartialError{
					{Title: "test partical error"},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.UserLookupUsersIDResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.UserLookupUsersIDResponse{
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

func Test_UserLookupUsersBy_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.UserLookupUsersByResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.UserLookupUsersByResponse{
				Errors: []resources.PartialError{
					{Title: "test partical error"},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.UserLookupUsersByResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.UserLookupUsersByResponse{
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

func Test_UserLookupUsersByUsername_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.UserLookupUsersByUsernameResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.UserLookupUsersByUsernameResponse{
				Errors: []resources.PartialError{
					{Title: "test partical error"},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.UserLookupUsersByUsernameResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.UserLookupUsersByUsernameResponse{
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
