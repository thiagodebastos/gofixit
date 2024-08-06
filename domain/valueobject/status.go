package valueobject

type Status int

const (
	StatusOpen = iota
	StatusInProgress
	StatusResolved
	StatusClosed
	StatusReopened
)

var statusNames = map[Status]string{
	StatusOpen:       "open",
	StatusInProgress: "inprogress",
	StatusResolved:   "resolved",
	StatusClosed:     "closed",
	StatusReopened:   "reopened",
}

func NewStatus(value string) (Status, error) {
	for s, name := range statusNames {
		if name == value {
			return s, nil
		}
	}
	return Status(-1), &InvalidStatusError{Value: value}
}

// getter function that returns the IssueStatus value
func (s Status) Value() Status {
	return s
}

func (s Status) ToString() string {
	return statusNames[s]
}
