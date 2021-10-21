package types

import "github.com/michimani/gotwi/resources"

type TweetLikesLikingUsersResponse struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes,omitempty"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *TweetLikesLikingUsersResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type TweetLikesLikedTweetsResponse struct {
	Data     []resources.Tweet `json:"data"`
	Meta     resources.PaginationMeta
	Includes struct {
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes,omitempty"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *TweetLikesLikedTweetsResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type TweetLikesPostResponse struct {
	Data struct {
		Liked bool `json:"liked"`
	} `json:"data"`
}

func (r *TweetLikesPostResponse) HasPartialError() bool {
	return false
}

type TweetLikesDeleteResponse struct {
	Data struct {
		Liked bool `json:"liked"`
	} `json:"data"`
}

func (r *TweetLikesDeleteResponse) HasPartialError() bool {
	return false
}
