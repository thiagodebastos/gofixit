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
		expectErr   bool
	}{
		{
			name:        "Valid Issue",
			title:       "Test Issue",
			description: "This is a test issue.",
			status:      valueobject.StatusOpen,
		}, {
			name:        "Invalid Status",
			title:       "Test issue",
			description: "This is a test issue.",
			status:      valueobject.StatusInvalid,
			expectErr:   true,
		}, {
			name:        "Invalid Title",
			title:       "",
			description: "",
			status:      valueobject.StatusOpen,
		},
	}

	for _, tt := range tests {
		id := uuid.New()
		newIssue, err := NewIssue(id, tt.title, tt.description, tt.status)

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

func TestUpdateIssueStatus(t *testing.T) {
	tests := []struct {
		name       string
		status     valueobject.Status
		nextStatus valueobject.Status
		expectErr  bool
	}{
		{
			name:       "closed -> reopened",
			status:     valueobject.StatusClosed,
			nextStatus: valueobject.StatusReopened,
			expectErr:  false,
		}, {
			name:       "closed -> resolved",
			status:     valueobject.StatusClosed,
			nextStatus: valueobject.StatusResolved,
			expectErr:  false,
		}, {
			name:       "closed -> opened",
			status:     valueobject.StatusClosed,
			nextStatus: valueobject.StatusOpen,
			expectErr:  true,
		}, {
			name:       "closed -> inprogress",
			status:     valueobject.StatusClosed,
			nextStatus: valueobject.StatusInProgress,
			expectErr:  true,
		}, {
			name:       "resolved -> reopened",
			status:     valueobject.StatusResolved,
			nextStatus: valueobject.StatusReopened,
			expectErr:  false,
		}, {
			name:       "resolved -> open",
			status:     valueobject.StatusResolved,
			nextStatus: valueobject.StatusOpen,
			expectErr:  true,
		}, {
			name:       "open -> reopened",
			status:     valueobject.StatusOpen,
			nextStatus: valueobject.StatusReopened,
			expectErr:  true,
		}, {
			name:       "open -> inprogress",
			status:     valueobject.StatusOpen,
			nextStatus: valueobject.StatusInProgress,
			expectErr:  false,
		}, {
			name:       "open -> resolved",
			status:     valueobject.StatusOpen,
			nextStatus: valueobject.StatusResolved,
			expectErr:  false,
		}, {
			name:       "open -> closed",
			status:     valueobject.StatusOpen,
			nextStatus: valueobject.StatusClosed,
			expectErr:  false,
		}, {
			// TODO: inprogress should only move to reopened if it HAS been closed closed or resolved before
			name:       "inprogress -> reopened",
			status:     valueobject.StatusInProgress,
			nextStatus: valueobject.StatusReopened,
			expectErr:  false,
		}, {
			name:       "inprogress -> closed",
			status:     valueobject.StatusInProgress,
			nextStatus: valueobject.StatusClosed,
			expectErr:  false,
		}, {
			name:       "inprogress -> resolved",
			status:     valueobject.StatusInProgress,
			nextStatus: valueobject.StatusResolved,
			expectErr:  false,
		}, {
			// TODO: inprogress should only move to opened if it has NOT been closed closed or resolved before
			name:       "inprogress -> open",
			status:     valueobject.StatusInProgress,
			nextStatus: valueobject.StatusOpen,
			expectErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newIssue, err := NewIssue(uuid.New(), "Test Issue", "", tt.status)
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

func TestUpdateIssueTitle(t *testing.T) {
	newIssue, err := NewIssue(
		uuid.New(),
		"Test Title",
		"Test Description",
		valueobject.StatusOpen,
	)
	if err != nil {
		t.Fatalf("failed to create a new issue: %v", err)
	}

	err = newIssue.UpdateTitle("")
	newTitle := newIssue.title
	if err != nil {
		t.Errorf("expected updateTitleError, got nil, %v", newTitle)
	}
}

func TestSetIssuePriority(t *testing.T) {
	newIssue, err := NewIssue(
		uuid.New(),
		"Test Title",
		"Test Description",
		valueobject.StatusOpen,
	)
	if err != nil {
		t.Fatalf("failed to create a new issue: %v", err)
	}

	err = newIssue.SetPriority(valueobject.PriorityHigh)
	if err != nil {
		t.Errorf("error setting issue priority: %v", err)
	}
}
