package valueobject

import "fmt"

type InvalidStatusError struct {
	Value string
}

func (e *InvalidStatusError) Error() string {
	return fmt.Sprintf("invalid status: %s", e.Value)
}

type InvalidPriorityError struct {
	Value string
}

func (e *InvalidPriorityError) Error() string {
	return fmt.Sprintf("invalid priority %s", e.Value)
}
