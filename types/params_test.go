package types_test

import (
	"testing"

	"github.com/michimani/gotwi/types"
	"github.com/stretchr/testify/assert"
)

func TestQueryString(t *testing.T) {
	cases := []struct {
		name   string
		params []string
		expect string
	}{
		{
			name:   "normal 1",
			params: []string{"param1", "param2", "param3"},
			expect: "param1,param2,param3",
		},
		{
			name:   "normal 2: only one param",
			params: []string{"param1"},
			expect: "param1",
		},
		{
			name:   "normal 3: empty params",
			params: []string{},
			expect: "",
		},
		{
			name:   "normal 4: nil params",
			params: nil,
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			q := types.QueryValue(c.params)
			assert.Equal(tt, c.expect, q)
		})
	}
}
