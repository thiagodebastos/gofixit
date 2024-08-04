package entity

import (
	"testing"

	"github.com/google/uuid"
)

// TODO: test multiple issues with table-driven tests
// table-driven tests are common in go
func TestNewIssue(t *testing.T) {
	title := "Test Issue"
	description := "This is a test issue."
	issue := NewIssue(uuid.New(), title, description)

	if issue.id == uuid.Nil {
		t.Errorf("expected a valid UUID, got %v", issue.id)
	}

	if issue.Title == "" {
		t.Errorf("expected a valid title, got %v", issue.Title)
	}
}
