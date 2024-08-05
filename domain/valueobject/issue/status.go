package issue

import "errors"

type Status int

const (
	StatusInvalid = iota
	StatusOpen
	StatusInProgress
	StatusResolved
	StatusClosed
	StatusReopened
)

var statusName = map[Status]string{
	StatusInvalid:    "invalid",
	StatusOpen:       "open",
	StatusInProgress: "inprogress",
	StatusResolved:   "resolved",
	StatusClosed:     "closed",
	StatusReopened:   "reopened",
}

var (
	ErrInvalidStatus = errors.New("invalid status")
	validStatuses    = map[Status]bool{
		StatusInvalid:    false,
		StatusOpen:       true,
		StatusClosed:     true,
		StatusInProgress: true,
		StatusResolved:   true,
		StatusReopened:   true,
	}
)

// constructor function that instantiates a Status object and validates the input
func NewStatus(value Status) (Status, error) {
	if !validStatuses[value] {
		return StatusInvalid, ErrInvalidStatus
	}

	// validate that new status is always of type status
	return value, nil
}

// getter function that returns the IssueStatus value
func (s Status) Value() Status {
	return s
}

func ValidStatus(value Status) bool {
	return validStatuses[value]
}
