package types

import "io"

type Parameters interface {
	SetAccessToken(token string)
	AccessToken() string
	ResolveEndpoint(endpointBase string) string
	Body() io.Reader
}
