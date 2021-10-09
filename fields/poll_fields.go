package fields

type PollField string

const (
	DurationMinutes PollField = "duration_minutes"
	EndDatetime     PollField = "end_datetime"
	PollID          PollField = "id"
	Options         PollField = "options"
	VotingStatus    PollField = "voting_status"
)
