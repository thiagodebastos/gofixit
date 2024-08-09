package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/thiagodebastos/gofixit/domain/valueobject"
)

func TestNewIssue(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		description string
		status      valueobject.Status
		priority    valueobject.Priority
		expectErr   bool
	}{
		{
			name:        "Valid Issue",
			title:       "Test Issue",
			description: "This is a test issue.",
			status:      valueobject.StatusOpen,
			priority:    valueobject.PriorityMedium,
			expectErr:   false,
		}, {
			name:        "Invalid Title",
			title:       "",
			description: "",
			priority:    valueobject.PriorityMedium,
			status:      valueobject.StatusOpen,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		id := uuid.New()
		newIssue, err := CreateIssue(id, tt.title, tt.description, tt.status, tt.priority)

		if tt.expectErr {
			if err == nil {
				t.Errorf("expected an error, got a valid issue")
			}
			return
		}

		if newIssue.ID() == uuid.Nil {
			t.Errorf("expected a valid UUID, got %v", newIssue.ID())
		}

		if newIssue.Title() == "" {
			t.Errorf("expected a valid title, got %v", newIssue.Title())
		}

		if newIssue.Status() != tt.status {
			t.Errorf("expected status to be \"Open\", got %v", newIssue.Status())
		}
	}
}

func TestSetIssueStatus(t *testing.T) {
	tests := []struct {
		name       string
		status     valueobject.Status
		priority   valueobject.Priority
		nextStatus valueobject.Status
		expectErr  bool
	}{
		{
			name:       "closed -> reopened",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusClosed,
			nextStatus: valueobject.StatusReopened,
			expectErr:  false,
		}, {
			name:       "closed -> resolved",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusClosed,
			nextStatus: valueobject.StatusResolved,
			expectErr:  false,
		}, {
			name:       "closed -> opened",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusClosed,
			nextStatus: valueobject.StatusOpen,
			expectErr:  true,
		}, {
			name:       "closed -> inprogress",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusClosed,
			nextStatus: valueobject.StatusInProgress,
			expectErr:  true,
		}, {
			name:       "resolved -> reopened",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusResolved,
			nextStatus: valueobject.StatusReopened,
			expectErr:  false,
		}, {
			name:       "resolved -> open",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusResolved,
			nextStatus: valueobject.StatusOpen,
			expectErr:  true,
		}, {
			name:       "open -> reopened",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusOpen,
			nextStatus: valueobject.StatusReopened,
			expectErr:  true,
		}, {
			name:       "open -> inprogress",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusOpen,
			nextStatus: valueobject.StatusInProgress,
			expectErr:  false,
		}, {
			name:       "open -> resolved",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusOpen,
			nextStatus: valueobject.StatusResolved,
			expectErr:  false,
		}, {
			name:       "open -> closed",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusOpen,
			nextStatus: valueobject.StatusClosed,
			expectErr:  false,
		}, {
			// TODO: inprogress should only move to reopened if it HAS been closed closed or resolved before
			name:       "inprogress -> reopened",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusInProgress,
			nextStatus: valueobject.StatusReopened,
			expectErr:  false,
		}, {
			name:       "inprogress -> closed",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusInProgress,
			nextStatus: valueobject.StatusClosed,
			expectErr:  false,
		}, {
			name:       "inprogress -> resolved",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusInProgress,
			nextStatus: valueobject.StatusResolved,
			expectErr:  false,
		}, {
			// TODO: inprogress should only move to opened if it has NOT been closed closed or resolved before
			name:       "inprogress -> open",
			priority:   valueobject.PriorityLow,
			status:     valueobject.StatusInProgress,
			nextStatus: valueobject.StatusOpen,
			expectErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newIssue, err := CreateIssue(uuid.New(), "Test Issue", "", tt.status, tt.priority)
			if err != nil {
				t.Fatalf("failed to create a new issue: %v", err)
			}

			err = newIssue.SetStatus(tt.nextStatus)
			if tt.expectErr {
				if err == nil {
					t.Error("expected an error but got none!")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				} else if newIssue.Status() != tt.nextStatus {
					t.Errorf("status not updated, expected %v, got %v", tt.nextStatus, newIssue.Status())
				}
			}
		})
	}
}

func TestSetIssueTitle(t *testing.T) {
	newIssue, err := CreateIssue(
		uuid.New(),
		"Test Title",
		"Test Description",
		valueobject.StatusOpen,
		valueobject.PriorityLow,
	)
	if err != nil {
		t.Fatalf("failed to create a new issue: %v", err)
	}

	err = newIssue.SetTitle("")
	newTitle := newIssue.Title()
	if err != nil {
		t.Errorf("expected updateTitleError, got nil, %v", newTitle)
	}
}

func TestSetIssuePriority(t *testing.T) {
	newIssue, err := CreateIssue(
		uuid.New(),
		"Test Title",
		"Test Description",
		valueobject.StatusOpen,
		valueobject.PriorityLow,
	)
	if err != nil {
		t.Fatalf("failed to create a new issue: %v", err)
	}

	err = newIssue.SetPriority(valueobject.PriorityHigh)
	if err != nil {
		t.Errorf("error setting issue priority: %v", err)
	}
}
