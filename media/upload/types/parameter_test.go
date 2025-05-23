package types_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
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

func Test_InitializeInput_ResolveEndpoint(t *testing.T) {
	cases := []struct {
		name         string
		endpointBase string
		expect       string
	}{
		{
			name:         "normal",
			endpointBase: "https://api.twitter.com/2/media/upload",
			expect:       "https://api.twitter.com/2/media/upload",
		},
		{
			name:         "empty",
			endpointBase: "",
			expect:       "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.InitializeInput{}
			endpoint := p.ResolveEndpoint(c.endpointBase)
			assert.Equal(tt, c.expect, endpoint)
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
	hash := func(s string) string {
		h := sha256.New()
		h.Write([]byte(s))
		return hex.EncodeToString(h.Sum(nil))
	}

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
			expect:     hash("test-media-id1"),
		},
		{
			name:       "empty mediaID",
			mediaID:    "",
			segmentIdx: 1,
			expect:     hash("1"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			p := &types.AppendInput{
				MediaID:      c.mediaID,
				SegmentIndex: c.segmentIdx,
			}
			boundary := p.GenerateBoundary()
			asst.Equal(c.expect, boundary)
		})
	}
}

func Test_AppendInput_Body(t *testing.T) {
	cases := []struct {
		name     string
		params   *types.AppendInput
		boundary string
		wantErr  bool
		expect   string
	}{
		{
			name: "error: boundary is not set",
			params: &types.AppendInput{
				MediaID:      "test-media-id",
				Media:        strings.NewReader("test-media"),
				SegmentIndex: 1,
			},
			wantErr: true,
		},
		{
			name: "ok",
			params: &types.AppendInput{
				MediaID:      "test-media-id",
				Media:        strings.NewReader("test-media"),
				SegmentIndex: 1,
			},
			boundary: "test-boundary",
			wantErr:  false,
			expect:   "--test-boundary\r\nContent-Disposition: form-data; name=\"media\"; filename=\"media\"\r\nContent-Type: application/octet-stream\r\n\r\ntest-media\r\n--test-boundary\r\nContent-Disposition: form-data; name=\"segment_index\"\r\n\r\n1\r\n--test-boundary--\r\n",
		},
	}

	for _, c := range cases {
		types.Exported_SetBoundary(c.params, c.boundary)

		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			b, err := c.params.Body()
			if c.wantErr {
				asst.Error(err)
				return
			}

			asst.NoError(err)
			buf := new(bytes.Buffer)
			buf.ReadFrom(b)
			asst.Equal(c.expect, buf.String())
		})
	}
}

func Test_AppendInput_Boundary(t *testing.T) {
	cases := []struct {
		name     string
		boundary string
		expect   string
	}{
		{
			name:     "normal",
			boundary: "test-boundary",
			expect:   "test-boundary",
		},
		{
			name:     "empty",
			boundary: "",
			expect:   "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.AppendInput{}
			types.Exported_SetBoundary(p, c.boundary)
			assert.Equal(tt, c.expect, p.Boundary())
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

func Test_AppendInput_ResolveEndpoint(t *testing.T) {
	cases := []struct {
		name         string
		endpointBase string
		mediaID      string
		expect       string
	}{
		{
			name:         "normal",
			endpointBase: "https://api.twitter.com/2/media/upload/:mediaID",
			mediaID:      "test-media-id",
			expect:       "https://api.twitter.com/2/media/upload/test-media-id",
		},
		{
			name:         "empty mediaID",
			endpointBase: "https://api.twitter.com/2/media/upload/:mediaID",
			mediaID:      "",
			expect:       "https://api.twitter.com/2/media/upload/",
		},
		{
			name:         "empty endpoint",
			endpointBase: "",
			mediaID:      "test-media-id",
			expect:       "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.AppendInput{
				MediaID: c.mediaID,
			}
			endpoint := p.ResolveEndpoint(c.endpointBase)
			assert.Equal(tt, c.expect, endpoint)
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

func Test_FinalizeInput_ResolveEndpoint(t *testing.T) {
	cases := []struct {
		name         string
		endpointBase string
		mediaID      string
		expect       string
	}{
		{
			name:         "normal",
			endpointBase: "https://api.twitter.com/2/media/upload/:mediaID",
			mediaID:      "test-media-id",
			expect:       "https://api.twitter.com/2/media/upload/test-media-id",
		},
		{
			name:         "empty mediaID",
			endpointBase: "https://api.twitter.com/2/media/upload/:mediaID",
			mediaID:      "",
			expect:       "https://api.twitter.com/2/media/upload/",
		},
		{
			name:         "empty endpoint",
			endpointBase: "",
			mediaID:      "test-media-id",
			expect:       "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.FinalizeInput{
				MediaID: c.mediaID,
			}
			endpoint := p.ResolveEndpoint(c.endpointBase)
			assert.Equal(tt, c.expect, endpoint)
		})
	}
}
