package types_test

import (
	"testing"

	"github.com/michimani/gotwi/media/upload/types"
	"github.com/michimani/gotwi/resources"
	"github.com/stretchr/testify/assert"
)

func Test_InitializeOutput_HasPartialError(t *testing.T) {
	var errorTitle string = "test partial error"
	cases := []struct {
		name   string
		res    *types.InitializeOutput
		expect bool
	}{
		{
			name: "has partial error",
			res: &types.InitializeOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				},
			},
			expect: true,
		},
		{
			name: "has no partial error",
			res: &types.InitializeOutput{
				Errors: []resources.PartialError{},
			},
			expect: false,
		},
		{
			name:   "partial error is nil",
			res:    &types.InitializeOutput{},
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

func Test_AppendOutput_HasPartialError(t *testing.T) {
	var errorTitle string = "test partial error"
	cases := []struct {
		name   string
		res    *types.AppendOutput
		expect bool
	}{
		{
			name: "has partial error",
			res: &types.AppendOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				},
			},
			expect: true,
		},
		{
			name: "has no partial error",
			res: &types.AppendOutput{
				Errors: []resources.PartialError{},
			},
			expect: false,
		},
		{
			name:   "partial error is nil",
			res:    &types.AppendOutput{},
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

func Test_FinalizeOutput_HasPartialError(t *testing.T) {
	var errorTitle string = "test partial error"
	cases := []struct {
		name   string
		res    *types.FinalizeOutput
		expect bool
	}{
		{
			name: "has partial error",
			res: &types.FinalizeOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				},
			},
			expect: true,
		},
		{
			name: "has no partial error",
			res: &types.FinalizeOutput{
				Errors: []resources.PartialError{},
			},
			expect: false,
		},
		{
			name:   "partial error is nil",
			res:    &types.FinalizeOutput{},
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
