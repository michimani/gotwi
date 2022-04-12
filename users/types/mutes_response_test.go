package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/users/types"
	"github.com/stretchr/testify/assert"
)

func Test_MutesMutingGet_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListMutedUsersOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListMutedUsersOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListMutedUsersOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListMutedUsersOutput{
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

func Test_CreateMutedUserOutput_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.CreateMutedUserOutput
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.CreateMutedUserOutput{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.CreateMutedUserOutput{
				Data: struct {
					Muting bool "json:\"muting\""
				}{
					Muting: false,
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

func Test_DeleteMutedUserOutput_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.DeleteMutedUserOutput
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.DeleteMutedUserOutput{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.DeleteMutedUserOutput{
				Data: struct {
					Muting bool "json:\"muting\""
				}{
					Muting: false,
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
