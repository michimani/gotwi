package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/compliance/batchcompliance/types"
	"github.com/stretchr/testify/assert"
)

func Test_ListJobsInput_SetAccessToken(t *testing.T) {
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
			p := &types.ListJobsInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_ListJobsInput_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.ListJobsInput
		expect string
	}{
		{
			name: "ok",
			params: &types.ListJobsInput{
				Type: types.ComplianceTypeTweets,
			},
			expect: endpoint + "?type=tweets",
		},
		{
			name: "ok: with type and status",
			params: &types.ListJobsInput{
				Type:   types.ComplianceTypeUsers,
				Status: types.ComplianceStatusFailed,
			},
			expect: endpoint + "?status=failed&type=users",
		},
		{
			name: "ng: has no required",
			params: &types.ListJobsInput{
				Status: types.ComplianceStatusFailed,
			},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpoint)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_ListJobsInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListJobsInput
	}{
		{
			name:   "empty params",
			params: &types.ListJobsInput{},
		},
		{
			name:   "nil",
			params: &types.ListJobsInput{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}

func Test_GetJobInput_SetAccessToken(t *testing.T) {
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
			p := &types.GetJobInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_GetJobInput_ResolveEndpoint(t *testing.T) {
	const endpointRoot = "test/endpoint/"
	const endpointBase = "test/endpoint/:id"
	cases := []struct {
		name   string
		params *types.GetJobInput
		expect string
	}{
		{
			name: "ok",
			params: &types.GetJobInput{
				ID: "test-id",
			},
			expect: endpointRoot + "test-id",
		},
		{
			name:   "ng: has no required",
			params: &types.GetJobInput{},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_GetJobInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.GetJobInput
	}{
		{
			name:   "empty params",
			params: &types.GetJobInput{},
		},
		{
			name:   "nil",
			params: &types.GetJobInput{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}

func Test_GetJobInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.GetJobInput
		expect map[string]string
	}{
		{
			name:   "normal: only required parameter",
			params: &types.GetJobInput{ID: "id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no required parameter",
			params: &types.GetJobInput{},
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

func Test_CreateJobInput_SetAccessToken(t *testing.T) {
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
			p := &types.CreateJobInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_CreateJobInput_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.CreateJobInput
		expect string
	}{
		{
			name:   "ok",
			params: &types.CreateJobInput{},
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

func Test_CreateJobInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateJobInput
		expect io.Reader
	}{
		{
			name: "ok: type",
			params: &types.CreateJobInput{
				Type: types.ComplianceTypeTweets,
			},
			expect: strings.NewReader(`{"type":"tweets"}`),
		},
		{
			name: "ok: name",
			params: &types.CreateJobInput{
				Name: gotwi.String("test-name"),
			},
			expect: strings.NewReader(`{"name":"test-name"}`),
		},
		{
			name: "ok: resumable",
			params: &types.CreateJobInput{
				Resumable: gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"resumable":true}`),
		},
		{
			name: "ok: all",
			params: &types.CreateJobInput{
				Type:      types.ComplianceTypeTweets,
				Name:      gotwi.String("test-name"),
				Resumable: gotwi.Bool(true),
			},
			expect: strings.NewReader(`{"type":"tweets","name":"test-name","resumable":true}`),
		},
		{
			name:   "ok: has no json parameters",
			params: &types.CreateJobInput{},
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

func Test_CreateJobInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.CreateJobInput
		expect map[string]string
	}{
		{
			name:   "ok",
			params: &types.CreateJobInput{},
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
