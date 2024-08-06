package valueobject

import (
	"testing"
)

func TestSetPriorityFromString(t *testing.T) {
	priorityFromStringTests := []struct {
		inputStr string
		want     bool
	}{
		{"lowest", true},
		{"Lowest", true},
		{"LOW", true},
		{"medium", true},
		{"high", true},
		{"hiGhest", true},
		{"invalid", false},
	}

	for _, tt := range priorityFromStringTests {
		_, got := PriorityFromString(tt.inputStr)

		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}

func TestGetPriorityString(t *testing.T) {
	priorityString := Priority(PriorityHigh).ToString()

	want := "high"
	got := priorityString

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
