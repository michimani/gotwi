package types

import (
	"encoding/json"
	"io"
	"strings"
)

type MediaCategory string

const (
	MediaCategoryAmplifyVideo MediaCategory = "amplify_video"
	MediaCategoryTweetGIF     MediaCategory = "tweet_gif"
	MediaCategoryTweetImage   MediaCategory = "tweet_image"
	MediaCategoryTweetVideo   MediaCategory = "tweet_video"
	MediaCategoryDMGIF        MediaCategory = "dm_gif"
	MediaCategoryDMImage      MediaCategory = "dm_image"
	MediaCategoryDMVideo      MediaCategory = "dm_video"
	MediaCategorySubtitles    MediaCategory = "subtitles"
)

type MediaType string

const (
	MediaTypeMP4       MediaType = "video/mp4"
	MediaTypeWebM      MediaType = "video/webm"
	MediaTypeMP2T      MediaType = "video/mp2t"
	MediaTypeQuickTime MediaType = "video/quicktime"
	MediaTypeSRT       MediaType = "text/srt"
	MediaTypeVTT       MediaType = "text/vtt"
	MediaTypeJPEG      MediaType = "image/jpeg"
	MediaTypeGIF       MediaType = "image/gif"
	MediaTypeBMP       MediaType = "image/bmp"
	MediaTypePNG       MediaType = "image/png"
	MediaTypeWebP      MediaType = "image/webp"
	MediaTypePJPEG     MediaType = "image/pjpeg"
	MediaTypeTIFF      MediaType = "image/tiff"
	MediaTypeGLTF      MediaType = "model/gltf-binary"
	MediaTypeUSDZ      MediaType = "model/vnd.usdz+zip"
)

// InitializeInput is the input for the Initialize endpoint.
type InitializeInput struct {
	accessToken string

	// Unique identifier of this User. This is returned as a string in order to avoid complications
	// with languages and tools that cannot handle large integers.
	AdditionalOwners []string `json:"additional_owners,omitempty"`

	// A string enum value which identifies a media use-case. This identifier is used to enforce use-case specific constraints
	// (e.g. file size, video duration) and enable advanced features.
	MediaCategory MediaCategory `json:"media_category,omitempty"`

	//The type of media.
	MediaType MediaType `json:"media_type,omitempty"`

	// Whether this media is shared or not.
	Shared bool `json:"shared,omitempty"`

	// The total size of the media upload in bytes.
	TotalBytes int `json:"total_bytes,omitempty"`
}

func (p *InitializeInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *InitializeInput) AccessToken() string {
	return p.accessToken
}

func (p *InitializeInput) ResolveEndpoint(endpointBase string) string {
	return endpointBase
}

func (p *InitializeInput) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *InitializeInput) ParameterMap() map[string]string {
	return map[string]string{}
}
