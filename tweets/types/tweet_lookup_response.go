package types

import "github.com/michimani/gotwi/resources"

type TweetLookupResponse struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *TweetLookupResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type TweetLookupIDResponse struct {
	Data     resources.Tweet `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *TweetLookupIDResponse) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
