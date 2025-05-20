package resources

type ProcessingInfoState string

const (
	ProcessingInfoStatePending    ProcessingInfoState = "pending"
	ProcessingInfoStateInProgress ProcessingInfoState = "in_progress"
	ProcessingInfoStateSucceeded  ProcessingInfoState = "succeeded"
	ProcessingInfoStateFailed     ProcessingInfoState = "failed"
)

type ProcessingInfo struct {
	// Number of seconds to check again for status
	CheckAfterSecs int `json:"check_after_secs"`

	// Percent of upload progress
	ProgressPercent int `json:"progress_percent"`

	// State of upload
	State ProcessingInfoState `json:"state"`
}

type UploadedMedia struct {
	// The unique identifier of this Media.
	MediaID string `json:"id"`

	// The Media Key identifier for this attachment.
	MediaKey string `json:"media_key"`

	// Number of seconds after which upload session expires.
	ExpiresAfterSecs int `json:"expires_after_secs"`

	// Processing information for the media.
	ProcessingInfo ProcessingInfo `json:"processing_info"`

	// Size of the upload
	Size uint `json:"size"`
}
