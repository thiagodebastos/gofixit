package valueobject

import "strings"

type Priority int

// Priority can be Highest, High, Medium, Low, Lowest
const (
	PriorityLowest = iota
	PriorityLow
	PriorityMedium
	PriorityHigh
	PriorityHighest
)

var priorityNames = map[Priority]string{
	PriorityLowest:  "lowest",
	PriorityLow:     "low",
	PriorityMedium:  "medium",
	PriorityHigh:    "high",
	PriorityHighest: "highest",
}

func NewPriority(p Priority) (Priority, error) {
	return p, nil
}

func PriorityFromString(s string) (Priority, bool) {
	s = strings.ToLower(s)
	for p, name := range priorityNames {
		if name == s {
			return Priority(p), true
		}
	}
	return Priority(-1), false
}

func ToString(p Priority) string {
	return priorityNames[p]
}
