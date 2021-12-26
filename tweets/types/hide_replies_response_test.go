package types_test

import (
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_HideRepliesResponse_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.HideRepliesResponse
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.HideRepliesResponse{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.HideRepliesResponse{
				Data: struct {
					Hidden bool "json:\"hidden\""
				}{
					Hidden: *gotwi.Bool(true),
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
