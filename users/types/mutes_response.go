package types

import "github.com/michimani/gotwi/resources"

// ListMutedUsersOutput is struct for response of `GET /2/users/:id/muting`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/get-users-muting
// more information:
type ListMutedUsersOutput struct {
	Data     []resources.User         `json:"data"`
	Meta     resources.PaginationMeta `json:"meta"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListMutedUsersOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

// CreateMutedUserOutput is struct for response of `POST /2/users/:id/muting`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/post-users-user_id-muting
// more information:
type CreateMutedUserOutput struct {
	Data struct {
		Muting bool `json:"muting"`
	} `json:"data"`
}

func (r *CreateMutedUserOutput) HasPartialError() bool {
	return false
}

// DeleteMutedUserOutput is struct for response of `DELETE /2/users/:source_user_id/muting/:target_user_id`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/mutes/api-reference/delete-users-user_id-muting
// more information:
type DeleteMutedUserOutput struct {
	Data struct {
		Muting bool `json:"muting"`
	} `json:"data"`
}

func (r *DeleteMutedUserOutput) HasPartialError() bool {
	return false
}
