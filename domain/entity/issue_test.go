package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/thiagodebastos/gofixit/domain/valueobject/issue"
)

func TestNewIssue(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		description string
		status      issue.Status
		expectErr   bool
	}{
		{
			name:        "Valid Issue",
			title:       "Test Issue",
			description: "This is a test issue.",
			status:      issue.StatusOpen,
		}, {
			name:        "Invalid Status",
			title:       "Test issue",
			description: "This is a test issue.",
			status:      issue.StatusInvalid,
			expectErr:   true,
		}, {
			name:        "Invalid Title",
			title:       "",
			description: "",
			status:      issue.StatusOpen,
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
		status     issue.Status
		nextStatus issue.Status
		expectErr  bool
	}{
		{
			name:       "closed -> reopened",
			status:     issue.StatusClosed,
			nextStatus: issue.StatusReopened,
			expectErr:  false,
		}, {
			name:       "closed -> resolved",
			status:     issue.StatusClosed,
			nextStatus: issue.StatusResolved,
			expectErr:  false,
		}, {
			name:       "closed -> opened",
			status:     issue.StatusClosed,
			nextStatus: issue.StatusOpen,
			expectErr:  true,
		}, {
			name:       "closed -> inprogress",
			status:     issue.StatusClosed,
			nextStatus: issue.StatusInProgress,
			expectErr:  true,
		}, {
			name:       "resolved -> reopened",
			status:     issue.StatusResolved,
			nextStatus: issue.StatusReopened,
			expectErr:  false,
		}, {
			name:       "resolved -> open",
			status:     issue.StatusResolved,
			nextStatus: issue.StatusOpen,
			expectErr:  true,
		}, {
			name:       "open -> reopened",
			status:     issue.StatusOpen,
			nextStatus: issue.StatusReopened,
			expectErr:  true,
		}, {
			name:       "open -> inprogress",
			status:     issue.StatusOpen,
			nextStatus: issue.StatusInProgress,
			expectErr:  false,
		}, {
			name:       "open -> resolved",
			status:     issue.StatusOpen,
			nextStatus: issue.StatusResolved,
			expectErr:  false,
		}, {
			name:       "open -> closed",
			status:     issue.StatusOpen,
			nextStatus: issue.StatusClosed,
			expectErr:  false,
		}, {
			// TODO: inprogress should only move to reopened if it HAS been closed closed or resolved before
			name:       "inprogress -> reopened",
			status:     issue.StatusInProgress,
			nextStatus: issue.StatusReopened,
			expectErr:  false,
		}, {
			name:       "inprogress -> closed",
			status:     issue.StatusInProgress,
			nextStatus: issue.StatusClosed,
			expectErr:  false,
		}, {
			name:       "inprogress -> resolved",
			status:     issue.StatusInProgress,
			nextStatus: issue.StatusResolved,
			expectErr:  false,
		}, {
			// TODO: inprogress should only move to opened if it has NOT been closed closed or resolved before
			name:       "inprogress -> open",
			status:     issue.StatusInProgress,
			nextStatus: issue.StatusOpen,
			expectErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newIssue, err := NewIssue(uuid.New(), "Test Issue", "", tt.status)
			if err != nil {
				t.Fatalf("failed to create a new issue: %v", err)
			}

			err = newIssue.ChangeStatus(tt.nextStatus)
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
	newIssue, err := NewIssue(uuid.New(), "Test Title", "Test Description", issue.StatusOpen)
	if err != nil {
		t.Fatalf("failed to create a new issue: %v", err)
	}

	err = newIssue.UpdateTitle("")
	newTitle := newIssue.title
	if err != nil {
		t.Errorf("expected updateTitleError, got nil, %v", newTitle)
	}
}
