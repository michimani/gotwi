package types_test

import (
	"testing"

	"github.com/michimani/gotwi/list/managelist/types"
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
