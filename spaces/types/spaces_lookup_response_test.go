package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/spaces/types"
	"github.com/stretchr/testify/assert"
)

func Test_SpacesLookupID_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.SpacesLookupIDResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.SpacesLookupIDResponse{
				Errors: []resources.PartialError{
					{Title: "test partical error"},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.SpacesLookupIDResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.SpacesLookupIDResponse{
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

func Test_SpacesLookup_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.SpacesLookupResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.SpacesLookupResponse{
				Errors: []resources.PartialError{
					{Title: "test partical error"},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.SpacesLookupResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.SpacesLookupResponse{
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

func Test_SpacesLookupByCreatorIDs_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.SpacesLookupByCreatorIDsResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.SpacesLookupByCreatorIDsResponse{
				Errors: []resources.PartialError{
					{Title: "test partical error"},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.SpacesLookupByCreatorIDsResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.SpacesLookupByCreatorIDsResponse{
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
