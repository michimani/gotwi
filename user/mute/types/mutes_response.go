package types

import "github.com/michimani/gotwi/resources"

// ListsOutput is struct for response of `GET /2/users/:id/muting`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
// more information:
type ListsOutput struct {
	Data     []resources.User         `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListsOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

// CreateOutput is struct for response of `POST /2/users/:id/muting`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/post-users-user_id-muting
// more information:
type CreateOutput struct {
	Data struct {
		Muting bool `json:"muting"`
	} `json:"data"`
}

func (r *CreateOutput) HasPartialError() bool {
	return false
}

// DeleteOutput is struct for response of `DELETE /2/users/:source_user_id/muting/:target_user_id`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/delete-users-user_id-muting
// more information:
type DeleteOutput struct {
	Data struct {
		Muting bool `json:"muting"`
	} `json:"data"`
}

func (r *DeleteOutput) HasPartialError() bool {
	return false
}
