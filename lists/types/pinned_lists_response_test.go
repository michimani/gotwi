package types_test

import (
	"testing"

	"github.com/michimani/gotwi/lists/types"
	"github.com/michimani/gotwi/resources"
	"github.com/stretchr/testify/assert"
)

func Test_PinnedListsGetResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.PinnedListsGetResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.PinnedListsGetResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.PinnedListsGetResponse{
				Data: []resources.List{},
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

func Test_PinnedListsPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.PinnedListsPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.PinnedListsPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.PinnedListsPostResponse{
				Data: struct {
					Pinned bool "json:\"pinned\""
				}{
					Pinned: false,
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

func Test_PinnedListsDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.PinnedListsDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.PinnedListsDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.PinnedListsDeleteResponse{
				Data: struct {
					Pinned bool "json:\"pinned\""
				}{
					Pinned: false,
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
