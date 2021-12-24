package util_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/michimani/gotwi/internal/util"
	"github.com/stretchr/testify/assert"
)

func Test_GetRateLimitInformation(t *testing.T) {
	resetTime := time.Unix(int64(100000000), 0)

	cases := []struct {
		name    string
		res     *http.Response
		wantErr bool
		expect  *util.RateLimitInformation
	}{
		{
			name: "normal",
			res: &http.Response{
				Header: http.Header{
					"X-Rate-Limit-Limit":     []string{"1"},
					"X-Rate-Limit-Remaining": []string{"100"},
					"X-Rate-Limit-Reset":     []string{"100000000"},
				},
			},
			wantErr: false,
			expect: &util.RateLimitInformation{
				Limit:     1,
				Remaining: 100,
				ResetAt:   &resetTime,
			},
		},
		{
			name: "normal: limit is empty",
			res: &http.Response{
				Header: http.Header{
					"X-Rate-Limit-Limit":     []string{},
					"X-Rate-Limit-Remaining": []string{"100"},
					"X-Rate-Limit-Reset":     []string{"100000000"},
				},
			},
			wantErr: false,
			expect: &util.RateLimitInformation{
				Limit:     0,
				Remaining: 100,
				ResetAt:   &resetTime,
			},
		},
		{
			name: "normal: remaining is empty",
			res: &http.Response{
				Header: http.Header{
					"X-Rate-Limit-Limit":     []string{"1"},
					"X-Rate-Limit-Remaining": []string{},
					"X-Rate-Limit-Reset":     []string{"100000000"},
				},
			},
			wantErr: false,
			expect: &util.RateLimitInformation{
				Limit:     1,
				Remaining: 0,
				ResetAt:   &resetTime,
			},
		},
		{
			name: "normal: reset is empty",
			res: &http.Response{
				Header: http.Header{
					"X-Rate-Limit-Limit":     []string{"1"},
					"X-Rate-Limit-Remaining": []string{"100"},
					"X-Rate-Limit-Reset":     []string{},
				},
			},
			wantErr: false,
			expect: &util.RateLimitInformation{
				Limit:     1,
				Remaining: 100,
				ResetAt:   nil,
			},
		},
		{
			name: "error: invalid rate limit limit value",
			res: &http.Response{
				Header: http.Header{
					"X-Rate-Limit-Limit":     []string{"a"},
					"X-Rate-Limit-Remaining": []string{"100"},
					"X-Rate-Limit-Reset":     []string{"100000000"},
				},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "error: invalid rate limit remaining value",
			res: &http.Response{
				Header: http.Header{
					"X-Rate-Limit-Limit":     []string{"1"},
					"X-Rate-Limit-Remaining": []string{"a"},
					"X-Rate-Limit-Reset":     []string{"100000000"},
				},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "error: invalid rate limit reset value",
			res: &http.Response{
				Header: http.Header{
					"X-Rate-Limit-Limit":     []string{"1"},
					"X-Rate-Limit-Remaining": []string{"100"},
					"X-Rate-Limit-Reset":     []string{"a"},
				},
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ri, err := util.GetRateLimitInformation(c.res)
			if c.wantErr {
				assert.Error(tt, err)
				assert.Nil(tt, ri)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, ri)
		})
	}
}
