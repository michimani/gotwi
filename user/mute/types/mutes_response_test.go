package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/user/mute/types"
	"github.com/stretchr/testify/assert"
)

func Test_MutesMutingGet_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListsOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListsOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListsOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListsOutput{
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

func Test_CreateOutput_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.CreateOutput
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.CreateOutput{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.CreateOutput{
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

func Test_DeleteOutput_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.DeleteOutput
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.DeleteOutput{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.DeleteOutput{
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
