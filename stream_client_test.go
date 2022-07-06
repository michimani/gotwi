package gotwi_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/stretchr/testify/assert"
)

func Test_newStreamClient(t *testing.T) {
	cases := []struct {
		name    string
		httpRes *http.Response
		wantErr bool
	}{
		{
			name: "ok",
			httpRes: &http.Response{
				Body: io.NopCloser(strings.NewReader(`{}`)),
			},
		},
		{
			name: "ng: closed body",
			httpRes: &http.Response{
				Close: true,
			},
			wantErr: true,
		},
		{
			name:    "ng: nil",
			httpRes: nil,
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			sc, err := gotwi.ExportNewStreamClient(c.httpRes)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(sc)
				return
			}

			asst.NoError(err)
			asst.NotNil(sc)
		})
	}
}

func Test_Receive(t *testing.T) {
	st, _ := gotwi.ExportNewStreamClient(&http.Response{
		Body: io.NopCloser(strings.NewReader(`{}`)),
	})

	cases := []struct {
		name   string
		st     *gotwi.StreamClient[*gotwi.MockResponse]
		close  bool
		expect bool
	}{
		{
			name:   "true",
			st:     st,
			expect: true,
		},
		{
			name:   "false",
			st:     st,
			close:  true,
			expect: false,
		},
		{
			name:   "false (nil)",
			st:     nil,
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			if c.close {
				c.st.Stop()
			}

			b := c.st.Receive()
			asst.Equal(c.expect, b)
		})
	}
}

func Test_Stop(t *testing.T) {
	st, _ := gotwi.ExportNewStreamClient(&http.Response{
		Body: io.NopCloser(strings.NewReader(`{}`)),
	})

	cases := []struct {
		name   string
		st     *gotwi.StreamClient[*gotwi.MockResponse]
		close  bool
		expect bool
	}{
		{
			name:   "ok",
			st:     st,
			expect: true,
		},
		{
			name:   "ok",
			st:     st,
			close:  true,
			expect: false,
		},
		{
			name:   "ok (nil)",
			st:     nil,
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			if c.close {
				c.st.Stop()
			}

			c.st.Stop()
		})
	}
}

func Test_Read(t *testing.T) {
	st, _ := gotwi.ExportNewStreamClient(&http.Response{
		Header: map[string][]string{
			"Content-Type": {"application/json;charset=UTF-8"},
		},
		Body: io.NopCloser(strings.NewReader(`{"text": "test"}`)),
	})

	cases := []struct {
		name    string
		st      *gotwi.StreamClient[*gotwi.MockResponse]
		close   bool
		wantErr bool
		expect  *gotwi.MockResponse
	}{
		{
			name: "ok",
			st:   st,
			expect: &gotwi.MockResponse{
				Text: "test",
			},
		},
		{
			name:    "ok (closed)",
			st:      st,
			close:   true,
			wantErr: false,
			expect:  nil,
		},
		{
			name:    "ok (nil)",
			st:      nil,
			wantErr: true,
			expect:  nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			if c.close {
				c.st.Stop()
			}

			st.Receive()
			out, err := c.st.Read()
			if c.wantErr {
				asst.Error(err)
				asst.Nil(out)
				return
			}

			asst.NoError(err)
			asst.Equal(c.expect, out)
		})
	}
}
