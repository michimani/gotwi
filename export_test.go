package gotwi

var (
	ExportNewRequest            = newRequest
	ExportResolveNon2XXResponse = resolveNon2XXResponse

	ExportGenerateOAthNonce     = generateOAthNonce
	ExportEndpointBase          = endpointBase
	ExportCreateParameterString = createParameterString
	ExportCreateSignatureBase   = createSignatureBase
	ExportCalculateSignature    = calculateSignature

	ExportWrapErr            = wrapErr
	ExportWrapWithAPIErr     = wrapWithAPIErr
	ExportNon2XXErrorSummary = non2XXErrorSummary
)
