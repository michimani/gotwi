package util_test

import (
	"testing"

	"github.com/michimani/gotwi/internal/util"
	"github.com/stretchr/testify/assert"
)

func Test_QueryValue(t *testing.T) {
	cases := []struct {
		name   string
		params []string
		expect string
	}{
		{
			name:   "normal",
			params: []string{"param1", "param2", "param3"},
			expect: "param1,param2,param3",
		},
		{
			name:   "normal: only one param",
			params: []string{"param1"},
			expect: "param1",
		},
		{
			name:   "normal: empty params",
			params: []string{},
			expect: "",
		},
		{
			name:   "normal: nil params",
			params: nil,
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			q := util.QueryValue(c.params)
			assert.Equal(tt, c.expect, q)
		})
	}
}

func Test_QueryString(t *testing.T) {
	cases := []struct {
		name     string
		includes map[string]struct{}
		params   map[string]string
		expect   string
	}{
		{
			name: "ok",
			includes: map[string]struct{}{
				"key1": {},
			},
			params: map[string]string{
				"key1": "value1",
			},
			expect: "key1=value1",
		},
		{
			name: "ok: some params",
			includes: map[string]struct{}{
				"key1": {},
				"key2": {},
			},
			params: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			expect: "key1=value1&key2=value2",
		},
		{
			name:     "ok: empty includes",
			includes: map[string]struct{}{},
			params: map[string]string{
				"key1": "value1",
			},
			expect: "",
		},
		{
			name: "ok: empty params",
			includes: map[string]struct{}{
				"key1": {},
			},
			params: map[string]string{},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			q := util.QueryString(c.params, c.includes)
			assert.Equal(tt, c.expect, q)
		})
	}
}
