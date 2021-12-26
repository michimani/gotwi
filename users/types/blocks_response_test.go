package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/users/types"
	"github.com/stretchr/testify/assert"
)

func Test_BlocksBlockingGet_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.BlocksBlockingGetResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.BlocksBlockingGetResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.BlocksBlockingGetResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.BlocksBlockingGetResponse{
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

func Test_BlocksBlockingPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.BlocksBlockingPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.BlocksBlockingPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.BlocksBlockingPostResponse{
				Data: struct {
					Blocking bool "json:\"blocking\""
				}{
					Blocking: false,
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

func Test_BlocksBlockingDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.BlocksBlockingDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.BlocksBlockingDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.BlocksBlockingDeleteResponse{
				Data: struct {
					Blocking bool "json:\"blocking\""
				}{
					Blocking: false,
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
