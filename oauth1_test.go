package gotwi_test

import (
	"testing"

	"github.com/michimani/gotwi"
	"github.com/stretchr/testify/assert"
)

func Test_EndpointDetail(t *testing.T) {
	cases := []struct {
		name     string
		endpoint gotwi.Endpoint
		expect   *gotwi.EndpointInfo
	}{
		{
			name:     "ok",
			endpoint: "endpoint",
			expect: &gotwi.EndpointInfo{
				Raw:                      "endpoint",
				Base:                     "endpoint",
				EncodedQueryParameterMap: map[string]string{},
			},
		},
		{
			name:     "ok with some parameters",
			endpoint: "endpoint?key1=value1&key2=value2",
			expect: &gotwi.EndpointInfo{
				Raw:  "endpoint?key1=value1&key2=value2",
				Base: "endpoint",
				EncodedQueryParameterMap: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			name:     "ok with encoded parameter",
			endpoint: "endpoint?key1=value1&key2=value2&key3=value%20value3",
			expect: &gotwi.EndpointInfo{
				Raw:  "endpoint?key1=value1&key2=value2&key3=value%20value3",
				Base: "endpoint",
				EncodedQueryParameterMap: map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value value3",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ed, err := c.endpoint.Detail()
			assert.NoError(tt, err)
			assert.Equal(tt, c.expect.Raw, ed.Raw)
			assert.Equal(tt, c.expect.Base, ed.Base)
			assert.Equal(tt, len(c.expect.EncodedQueryParameterMap), len(ed.EncodedQueryParameterMap))
			for k, v := range c.expect.EncodedQueryParameterMap {
				assert.Equal(tt, v, ed.EncodedQueryParameterMap[k])
			}
		})
	}
}
