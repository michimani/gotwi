package types_test

import (
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_ManageTweetsPostResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.ManageTweetsPostResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.ManageTweetsPostResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.ManageTweetsPostResponse{
				Data: struct {
					ID   *string "json:\"id\""
					Text *string "json:\"text\""
				}{
					ID:   gotwi.String("id"),
					Text: gotwi.String("text"),
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

func Test_ManageTweetsDeleteResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.ManageTweetsDeleteResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.ManageTweetsDeleteResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.ManageTweetsDeleteResponse{
				Data: struct {
					Deleted *bool "json:\"deleted\""
				}{
					Deleted: gotwi.Bool(true),
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
