package gotwierrors

const (
	ErrorClientNotReady string = "Twitter client is not ready."
	ErrorParametersNil  string = "Parameter for %s is nil."
	ErrorNon200Status   string = "Status is not OK. %d: %s '%s'" // StatusCode, ErrorTitle, ErrorText
)
