package util_test

import (
	"net/http"
	"testing"

	"github.com/michimani/gotwi/internal/util"
	"github.com/stretchr/testify/assert"
)

func Test_HeaderValue(t *testing.T) {
	cases := []struct {
		name   string
		header http.Header
		key    string
		expect []string
	}{
		{
			name: "normal",
			header: http.Header{
				"key1": []string{"value1-1", "value1-2"},
				"key2": []string{"value2-1", "value2-2"},
			},
			key:    "key1",
			expect: []string{"value1-1", "value1-2"},
		},
		{
			name: "normal: not exists key",
			header: http.Header{
				"key1": []string{"value1-1", "value1-2"},
				"key2": []string{"value2-1", "value2-2"},
			},
			key:    "key0",
			expect: []string{},
		},
		{
			name:   "normal: empty header",
			header: http.Header{},
			key:    "key",
			expect: []string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			vs := util.HeaderValues(c.key, c.header)

			assert.Len(tt, vs, len(c.expect))
			assert.Equal(tt, c.expect, vs)
		})
	}
}
