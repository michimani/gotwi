package types

import "github.com/michimani/gotwi/resources"

type ListTweetsOutput struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users  []resources.User  `json:"users,omitempty"`
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes,omitempty"`
	Meta   resources.TweetTimelineMeta `json:"meta"`
	Errors []resources.PartialError    `json:"errors,omitempty"`
}

func (r *ListTweetsOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

type ListMentionsOutput struct {
	Data     []resources.Tweet `json:"data"`
	Includes struct {
		Users  []resources.User  `json:"users,omitempty"`
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes,omitempty"`
	Meta   resources.TweetTimelineMeta `json:"meta"`
	Errors []resources.PartialError    `json:"errors,omitempty"`
}

func (r *ListMentionsOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
