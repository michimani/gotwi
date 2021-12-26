package types_test

import (
	"testing"

	"github.com/michimani/gotwi/lists/types"
	"github.com/stretchr/testify/assert"
)

func Test_ManageListsPutResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.ManageListsPutResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.ManageListsPutResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.ManageListsPutResponse{
				Updated: false,
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

func Test_ManageListsPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.ManageListsPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.ManageListsPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.ManageListsPostResponse{
				Data: struct {
					ID   string "json:\"id\""
					Name string "json:\"name\""
				}{
					ID: "id", Name: "name",
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

func Test_ManageListsDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.ManageListsDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.ManageListsDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.ManageListsDeleteResponse{
				Data: struct {
					Deleted bool "json:\"deleted\""
				}{
					Deleted: false,
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
