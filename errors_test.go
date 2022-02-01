package gotwi_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/internal/gotwierrors"
	"github.com/michimani/gotwi/internal/util"
	"github.com/michimani/gotwi/resources"
	"github.com/stretchr/testify/assert"
)

func Test_wrapErr(t *testing.T) {
	cases := []struct {
		name    string
		err     error
		wantNil bool
	}{
		{
			name: "normal",
			err:  errors.New("error test"),
		},
		{
			name: "normal: wrapped",
			err:  gotwi.ExportWrapErr(errors.New("error test")),
		},
		{
			name:    "nil",
			err:     nil,
			wantNil: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert := assert.New(tt)

			ge := gotwi.ExportWrapErr(c.err)
			if c.wantNil {
				assert.Nil(ge)
				return
			}

			assert.NotNil(ge)
			assert.Equal("error test", ge.Error())
			assert.False(ge.OnAPI)

			un := ge.Unwrap()
			_, ok := un.(*gotwi.GotwiError)
			assert.False(ok)
		})
	}
}

func Test_wrapWithAPIErr(t *testing.T) {
	resetAt := time.Date(2021, 10, 24, 23, 59, 59, 59, time.UTC)

	cases := []struct {
		name         string
		err          *resources.Non2XXError
		expectErrMsg string
		wantNil      bool
	}{
		{
			name:         "normal: empty",
			err:          &resources.Non2XXError{},
			expectErrMsg: "The Twitter API returned a Response with a status other than 2XX series.",
		},
		{
			name: "normal: full",
			err: &resources.Non2XXError{
				APIErrors: []resources.ErrorInformation{
					{
						Message: "api-err-1",
						Code:    3,
						Label:   "api-err-label-1",
					},
					{
						Message: "api-err-2",
						Code:    13,
						Label:   "api-err-label-2",
					},
				},
				Title:      "non 2xx error title",
				Detail:     "non 2xx error detail",
				Type:       "non 2xx error type",
				Status:     "non 2xx error status",
				StatusCode: 500,
				RateLimitInfo: &util.RateLimitInformation{
					Limit:     100,
					Remaining: 20,
					ResetAt:   &resetAt,
				},
			},
			expectErrMsg: strings.Join([]string{
				"The Twitter API returned a Response with a status other than 2XX series.",
				"httpStatus=\"non 2xx error status\"",
				"httpStatusCode=500",
				"title=\"non 2xx error title\"",
				"detail=\"non 2xx error detail\"",
				"errorCode1=3 errorText1=\"Invalid coordinates.\" errorDescription1=\"Corresponds with HTTP 400. The coordinates provided as parameters were not valid for the request.\"",
				"errorCode2=13 errorText2=\"No location associated with the specified IP address.\" errorDescription2=\"Corresponds with HTTP 404. It was not possible to derive a location for the IP address provided as a parameter on the geo search request.\"",
				"rateLimit=100 rateLimitRemaining=20 rateLimitReset=\"2021-10-24 23:59:59.000000059 +0000 UTC\"",
			}, " "),
		},
		{
			name: "normal: partial",
			err: &resources.Non2XXError{
				APIErrors: []resources.ErrorInformation{
					{
						Message: "api-err-1",
						Code:    3,
						Label:   "api-err-label-1",
					},
				},
				Title:      "non 2xx error title",
				Detail:     "non 2xx error detail",
				Type:       "non 2xx error type",
				Status:     "non 2xx error status",
				StatusCode: 500,
			},
			expectErrMsg: strings.Join([]string{
				"The Twitter API returned a Response with a status other than 2XX series.",
				"httpStatus=\"non 2xx error status\"",
				"httpStatusCode=500",
				"title=\"non 2xx error title\"",
				"detail=\"non 2xx error detail\"",
				"errorCode1=3 errorText1=\"Invalid coordinates.\" errorDescription1=\"Corresponds with HTTP 400. The coordinates provided as parameters were not valid for the request.\"",
			}, " "),
		},
		{
			name:    "nil",
			err:     nil,
			wantNil: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert := assert.New(tt)

			ge := gotwi.ExportWrapWithAPIErr(c.err)
			if c.wantNil {
				assert.Nil(ge)
				return
			}

			assert.NotNil(ge)
			assert.True(ge.OnAPI)
			assert.Equal(c.expectErrMsg, ge.Error())
		})
	}
}

func Test_non2XXErrorSummary(t *testing.T) {
	resetAt := time.Date(2021, 10, 24, 23, 59, 59, 59, time.UTC)

	cases := []struct {
		name         string
		err          *resources.Non2XXError
		expectErrMsg string
	}{
		{
			name:         "normal: empty",
			err:          &resources.Non2XXError{},
			expectErrMsg: "The Twitter API returned a Response with a status other than 2XX series.",
		},
		{
			name: "normal: full",
			err: &resources.Non2XXError{
				APIErrors: []resources.ErrorInformation{
					{
						Message: "api-err-1",
						Code:    3,
						Label:   "api-err-label-1",
					},
					{
						Message: "api-err-2",
						Code:    13,
						Label:   "api-err-label-2",
					},
				},
				Title:      "non 2xx error title",
				Detail:     "non 2xx error detail",
				Type:       "non 2xx error type",
				Status:     "non 2xx error status",
				StatusCode: 500,
				RateLimitInfo: &util.RateLimitInformation{
					Limit:     100,
					Remaining: 20,
					ResetAt:   &resetAt,
				},
			},
			expectErrMsg: strings.Join([]string{
				"The Twitter API returned a Response with a status other than 2XX series.",
				"httpStatus=\"non 2xx error status\"",
				"httpStatusCode=500",
				"title=\"non 2xx error title\"",
				"detail=\"non 2xx error detail\"",
				"errorCode1=3 errorText1=\"Invalid coordinates.\" errorDescription1=\"Corresponds with HTTP 400. The coordinates provided as parameters were not valid for the request.\"",
				"errorCode2=13 errorText2=\"No location associated with the specified IP address.\" errorDescription2=\"Corresponds with HTTP 404. It was not possible to derive a location for the IP address provided as a parameter on the geo search request.\"",
				"rateLimit=100 rateLimitRemaining=20 rateLimitReset=\"2021-10-24 23:59:59.000000059 +0000 UTC\"",
			}, " "),
		},
		{
			name: "normal: partial",
			err: &resources.Non2XXError{
				APIErrors: []resources.ErrorInformation{
					{
						Message: "api-err-1",
						Code:    3,
						Label:   "api-err-label-1",
					},
				},
				Title:      "non 2xx error title",
				Detail:     "non 2xx error detail",
				Type:       "non 2xx error type",
				Status:     "non 2xx error status",
				StatusCode: 500,
			},
			expectErrMsg: strings.Join([]string{
				"The Twitter API returned a Response with a status other than 2XX series.",
				"httpStatus=\"non 2xx error status\"",
				"httpStatusCode=500",
				"title=\"non 2xx error title\"",
				"detail=\"non 2xx error detail\"",
				"errorCode1=3 errorText1=\"Invalid coordinates.\" errorDescription1=\"Corresponds with HTTP 400. The coordinates provided as parameters were not valid for the request.\"",
			}, " "),
		},
		{
			name:         "nil",
			err:          nil,
			expectErrMsg: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			assert := assert.New(tt)

			s := gotwi.ExportNon2XXErrorSummary(c.err)
			assert.Equal(c.expectErrMsg, s)
		})
	}
}

func Test_GotwiErrorError(t *testing.T) {
	cases := []struct {
		name   string
		e      *gotwi.GotwiError
		expect string
	}{
		{
			name:   "normal",
			e:      gotwi.ExportWrapErr(errors.New("error test")),
			expect: "error test",
		},
		{
			name:   "normal: empty",
			e:      &gotwi.GotwiError{},
			expect: gotwierrors.ErrorUndefined,
		},
		{
			name:   "normal: nil",
			e:      nil,
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			a := assert.New(tt)
			e := c.e.Error()
			a.Equal(c.expect, e)
		})
	}
}

func Test_Unwrap(t *testing.T) {
	cases := []struct {
		name    string
		ge      *gotwi.GotwiError
		wantNil bool
	}{
		{
			name: "normal",
			ge:   &gotwi.GotwiError{},
		},
		{
			name:    "nil",
			ge:      nil,
			wantNil: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			a := assert.New(tt)
			e := c.ge.Unwrap()

			if c.wantNil {
				a.Nil(e)
				return
			}

			_, ok := e.(*gotwi.GotwiError)
			a.False(ok)
		})

	}

}

func Test_GotwiAccessNon2XXErrorFields(t *testing.T) {
	resetAt := time.Date(2021, 10, 24, 23, 59, 59, 59, time.UTC)

	n2e := &resources.Non2XXError{
		APIErrors: []resources.ErrorInformation{
			{
				Message: "api-err-1",
				Code:    3,
				Label:   "api-err-label-1",
			},
			{
				Message: "api-err-2",
				Code:    13,
				Label:   "api-err-label-2",
			},
		},
		Title:      "non 2xx error title",
		Detail:     "non 2xx error detail",
		Type:       "non 2xx error type",
		Status:     "non 2xx error status",
		StatusCode: 500,
		RateLimitInfo: &util.RateLimitInformation{
			Limit:     100,
			Remaining: 20,
			ResetAt:   &resetAt,
		},
	}

	a := assert.New(t)

	ge := gotwi.ExportWrapWithAPIErr(n2e)

	a.NotNil(ge)
	a.Equal("non 2xx error title", ge.Title)
	a.Equal("non 2xx error detail", ge.Detail)
	a.Equal("non 2xx error type", ge.Type)
	a.Equal("non 2xx error status", ge.Status)
	a.Equal(500, ge.StatusCode)
	a.Equal("api-err-1", ge.APIErrors[0].Message)
	a.Equal(resources.ErrorCode(3), ge.APIErrors[0].Code)
	a.Equal("api-err-label-1", ge.APIErrors[0].Label)
	a.Equal("api-err-2", ge.APIErrors[1].Message)
	a.Equal(resources.ErrorCode(13), ge.APIErrors[1].Code)
	a.Equal("api-err-label-2", ge.APIErrors[1].Label)
	a.Equal(100, ge.RateLimitInfo.Limit)
	a.Equal(20, ge.RateLimitInfo.Remaining)
	a.Equal(resetAt, *ge.RateLimitInfo.ResetAt)
}
