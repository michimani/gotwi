package types_test

import (
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/hidereply/types"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateOutput_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.UpdateOutput
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.UpdateOutput{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.UpdateOutput{
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
