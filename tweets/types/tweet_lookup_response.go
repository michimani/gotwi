package types

import "github.com/michimani/gotwi/resources"

type TweetLookupTweetsResponse struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *TweetLookupTweetsResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type TweetLookupTweetsIDResponse struct {
	Data     resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *TweetLookupTweetsIDResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
