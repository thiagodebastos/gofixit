package entity

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusTodo  Status = "todo"
	StatusDoing Status = "doing"
	StatusDone  Status = "done"
)

type Issue struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      Status
	Title       string
	Description string
	id          uuid.UUID
}

func (issue *Issue) UpdateStatus(newStatus Status) *Issue {
	issue.Status = newStatus
	issue.UpdatedAt = time.Now()
	return issue
}

func (issue *Issue) UpdateDescription(newDescription string) *Issue {
	issue.Description = newDescription
	return issue
}

func (issue *Issue) UpdateTitle(newTitle string) *Issue {
	issue.Title = newTitle
	return issue
}

// NewIssue creates a new issue
func NewIssue(id uuid.UUID, title string, description string) Issue {
	issue := Issue{
		id:          id,
		Title:       title,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return issue
}

func (i Issue) ID() uuid.UUID {
	return i.id
}
