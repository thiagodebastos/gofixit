package entity

import "fmt"

type InvalidIssueStateTransitionError struct {
	From string
	To   string
}

func (e *InvalidIssueStateTransitionError) Error() string {
	return fmt.Sprintf("invalid issue state transition from %s to %s", e.From, e.To)
}

type IssueValidationError struct {
	Field   string
	Message string
}

func (e *IssueValidationError) Error() string {
	return fmt.Sprintf("issue validation error: %s - %s", e.Field, e.Message)
}
