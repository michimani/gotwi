package response

import "github.com/michimani/gotwi/types/resource"

type Users struct {
	Data     []resource.User `json:"data"`
	Includes struct {
		Tweets []resource.Tweet `json:"tweets"`
	} `json:"includes"`
}

type UsersID struct {
	Data     resource.User `json:"data"`
	Includes struct {
		Tweets []resource.Tweet `json:"tweets"`
	} `json:"includes"`
}

type UsersByUsername struct {
	Data     resource.User `json:"data"`
	Includes struct {
		Tweets []resource.Tweet `json:"tweets"`
	} `json:"includes"`
}
