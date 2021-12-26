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
		res    *types.MutesMutingGetResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.MutesMutingGetResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.MutesMutingGetResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.MutesMutingGetResponse{
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

func Test_MutesMutingPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.MutesMutingPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.MutesMutingPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.MutesMutingPostResponse{
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

func Test_MutesMutingDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.MutesMutingDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.MutesMutingDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.MutesMutingDeleteResponse{
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
