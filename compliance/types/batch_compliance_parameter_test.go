package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/compliance/types"
	"github.com/michimani/gotwi/resources"
	"github.com/stretchr/testify/assert"
)

func Test_BatchComplianceJobsPostParams_SetAccessToken(t *testing.T) {
	cases := []struct {
		name   string
		token  string
		expect string
	}{
		{
			name:   "normal",
			token:  "test-token",
			expect: "test-token",
		},
		{
			name:   "normal: empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.BatchComplianceJobsPostParams{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_BatchComplianceJobsPostParams_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.BatchComplianceJobsPostParams
		expect string
	}{
		{
			name:   "ok",
			params: &types.BatchComplianceJobsPostParams{},
			expect: endpoint,
		},
		{
			name:   "ok: nil",
			params: nil,
			expect: endpoint,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpoint)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_BatchComplianceJobsPostParams_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BatchComplianceJobsPostParams
		expect io.Reader
	}{
		{
			name: "ok: type",
			params: &types.BatchComplianceJobsPostParams{
				Type: resources.ComplianceTypeTweets,
			},
			expect: strings.NewReader(`{"type":"tweets"}`),
		},
		{
			name: "ok: name",
			params: &types.BatchComplianceJobsPostParams{
				Name: gotwi.String("test-name"),
			},
			expect: strings.NewReader(`{"name":"test-name"}`),
		},
		{
			name: "ok: resumable",
			params: &types.BatchComplianceJobsPostParams{
				Resumable: gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"resumable":true}`),
		},
		{
			name: "ok: all",
			params: &types.BatchComplianceJobsPostParams{
				Type:      resources.ComplianceTypeTweets,
				Name:      gotwi.String("test-name"),
				Resumable: gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"type":"tweets","name":"test-name","resumable":true}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.BatchComplianceJobsPostParams{},
			expect: strings.NewReader(`{}`),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, r)
		})
	}
}

func Test_BatchComplianceJobsPostParams_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.BatchComplianceJobsPostParams
		expect map[string]string
	}{
		{
			name:   "ok",
			params: &types.BatchComplianceJobsPostParams{},
			expect: map[string]string{},
		},
		{
			name:   "ok: nil",
			params: nil,
			expect: map[string]string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			m := c.params.ParameterMap()
			assert.Equal(tt, c.expect, m)
		})
	}
}
