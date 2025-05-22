package types_test

import (
	"io"
	"strings"
	"testing"

	"github.com/michimani/gotwi/media/upload/types"
	"github.com/stretchr/testify/assert"
)

func Test_InitializeInput_SetAccessToken(t *testing.T) {
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
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.InitializeInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_InitializeInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.InitializeInput
		expect io.Reader
	}{
		{
			name: "ok: has all parameters",
			params: &types.InitializeInput{
				AdditionalOwners: []string{"owner1", "owner2"},
				MediaCategory:    types.MediaCategoryTweetImage,
				MediaType:        types.MediaTypeJPEG,
				Shared:           true,
				TotalBytes:       1024,
			},
			expect: strings.NewReader(`{"additional_owners":["owner1","owner2"],"media_category":"tweet_image","media_type":"image/jpeg","shared":true,"total_bytes":1024}`),
		},
		{
			name: "ok: has some parameters",
			params: &types.InitializeInput{
				MediaCategory: types.MediaCategoryTweetImage,
				MediaType:     types.MediaTypeJPEG,
				TotalBytes:    1024,
			},
			expect: strings.NewReader(`{"media_category":"tweet_image","media_type":"image/jpeg","total_bytes":1024}`),
		},
		{
			name:   "ok: has no parameters",
			params: &types.InitializeInput{},
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

func Test_InitializeInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.InitializeInput
		expect map[string]string
	}{
		{
			name:   "normal: has parameters",
			params: &types.InitializeInput{MediaCategory: types.MediaCategoryTweetImage},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameters",
			params: &types.InitializeInput{},
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

func Test_AppendInput_SetAccessToken(t *testing.T) {
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
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.AppendInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_AppendInput_GenerateBoundary(t *testing.T) {
	cases := []struct {
		name       string
		mediaID    string
		segmentIdx int
		expect     string
	}{
		{
			name:       "normal",
			mediaID:    "test-media-id",
			segmentIdx: 1,
			expect:     "test-media-id1",
		},
		{
			name:       "empty mediaID",
			mediaID:    "",
			segmentIdx: 1,
			expect:     "1",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.AppendInput{
				MediaID:      c.mediaID,
				SegmentIndex: c.segmentIdx,
			}
			boundary := p.GenerateBoundary()
			assert.NotEmpty(tt, boundary)
		})
	}
}

func Test_AppendInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.AppendInput
		expect error
	}{
		{
			name: "error: boundary is not set",
			params: &types.AppendInput{
				MediaID:      "test-media-id",
				Media:        strings.NewReader("test-media"),
				SegmentIndex: 1,
			},
			expect: assert.AnError,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			_, err := c.params.Body()
			assert.Error(tt, err)
		})
	}
}

func Test_AppendInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.AppendInput
		expect map[string]string
	}{
		{
			name:   "normal: has parameters",
			params: &types.AppendInput{MediaID: "test-media-id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameters",
			params: &types.AppendInput{},
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

func Test_FinalizeInput_SetAccessToken(t *testing.T) {
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
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.FinalizeInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_FinalizeInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.FinalizeInput
		expect io.Reader
	}{
		{
			name:   "normal: has parameters",
			params: &types.FinalizeInput{MediaID: "test-media-id"},
			expect: nil,
		},
		{
			name:   "normal: has no parameters",
			params: &types.FinalizeInput{},
			expect: nil,
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

func Test_FinalizeInput_ParameterMap(t *testing.T) {
	cases := []struct {
		name   string
		params *types.FinalizeInput
		expect map[string]string
	}{
		{
			name:   "normal: has parameters",
			params: &types.FinalizeInput{MediaID: "test-media-id"},
			expect: map[string]string{},
		},
		{
			name:   "normal: has no parameters",
			params: &types.FinalizeInput{},
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
