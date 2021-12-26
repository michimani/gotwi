package types_test

import (
	"testing"

	"github.com/michimani/gotwi/lists/types"
	"github.com/michimani/gotwi/resources"
	"github.com/stretchr/testify/assert"
)

func Test_ListMembersListMemberships_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListMembersListMembershipsResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListMembersListMembershipsResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListMembersListMembershipsResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListMembersListMembershipsResponse{
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

func Test_ListMembersGet_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListMembersGetResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListMembersGetResponse{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListMembersGetResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListMembersGetResponse{
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

func Test_ListMembersPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.ListMembersPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.ListMembersPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.ListMembersPostResponse{
				Data: struct {
					IsMember bool "json:\"is_member\""
				}{
					IsMember: false,
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

func Test_ListMembersDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.ListMembersDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.ListMembersDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.ListMembersDeleteResponse{
				Data: struct {
					IsMember bool "json:\"is_member\""
				}{
					IsMember: false,
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
