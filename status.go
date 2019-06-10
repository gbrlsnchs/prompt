package prompt

// Status is a response status.
type Status int

const (
	// StatusAccept means input matches accepting answers.
	StatusAccept Status = iota
	// StatusDecline means input matches declining answers.
	StatusDecline
	// StatusNone means input matches neither accepting nor declining answers.
	StatusNone
)
