package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/tweet/timeline/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetTimelinesTweets_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListTweetsOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListTweetsOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListTweetsOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListTweetsOutput{
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

func Test_ListMentionsOutput_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListMentionsOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListMentionsOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListMentionsOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListMentionsOutput{
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

func Test_ListReverseChronologicalOutput_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListReverseChronologicalOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListReverseChronologicalOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListReverseChronologicalOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListReverseChronologicalOutput{
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
