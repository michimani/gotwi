package response

type ClientResponse struct {
	StatusCode int
	Status     string
	Body       []byte
}
