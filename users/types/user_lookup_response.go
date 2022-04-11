package types

import "github.com/michimani/gotwi/resources"

// ListUsersOutput is struct for response of `GET /2/users`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users
type ListUsersOutput struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListUsersOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

// GetUserOutput is struct for response of `GET /2/users/:id`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
type GetUserOutput struct {
	Data     resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *GetUserOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

// ListUsersByUsernamesOutput is struct for response of `GET /2/users/by`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by
type ListUsersByUsernamesOutput struct {
	Data     []resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListUsersByUsernamesOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

// GetUserByUsernameOutput is struct for response of `GET /2/users/by/username/:username`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
type GetUserByUsernameOutput struct {
	Data     resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *GetUserByUsernameOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}

// GetMeOutput is struct for response of `GET /2/users/me`.
// more information: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
type GetMeOutput struct {
	Data     resources.User `json:"data"`
	Includes struct {
		Tweets []resources.Tweet `json:"tweets"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *GetMeOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
