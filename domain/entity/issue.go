package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/thiagodebastos/gofixit/domain/valueobject"
)

type Issue interface {
	ID() uuid.UUID
	Title() string
	Description() string
	Status() valueobject.Status
	Priority() valueobject.Priority

	SetStatus(newStatus valueobject.Status) error
	SetTitle(newTitle string) error
	SetDescription(newDescription string)
	SetPriority(newPriority valueobject.Priority) error

	validateIssueCreation() error
	validateStatusTransition(newStatus valueobject.Status) error
}

type issueEntity struct {
	title       string
	description string
	status      valueobject.Status
	id          uuid.UUID
	priority    valueobject.Priority
}

var ErrInvalidStatus = errors.New("invalid status")

// NewIssue is a factory responsible for creating a new Issue
func NewIssue(id uuid.UUID,
	title string,
	description string,
	status valueobject.Status,
	priority valueobject.Priority,
) (Issue, error) {
	issue := &issueEntity{
		id:          id,
		title:       title,
		status:      status,
		description: description,
		priority:    priority,
	}

	if err := issue.validateIssueCreation(); err != nil {
		return nil, err
	}

	return issue, nil
}

func (i *issueEntity) ID() uuid.UUID                  { return i.id }
func (i *issueEntity) Title() string                  { return i.title }
func (i *issueEntity) Status() valueobject.Status     { return i.status }
func (i *issueEntity) Description() string            { return i.description }
func (i *issueEntity) Priority() valueobject.Priority { return i.priority }

func (i *issueEntity) SetTitle(newTitle string) error {
	if i.Title() == "" {
		return errors.New("title cannot be empty")
	}

	i.title = newTitle
	return nil
}

func (i *issueEntity) SetStatus(newStatus valueobject.Status) error {
	if err := i.validateStatusTransition(newStatus); err != nil {
		return err
	}

	i.status = newStatus
	return nil
}

func (i *issueEntity) SetDescription(newDescription string) {
	i.description = newDescription
}

func (i *issueEntity) SetPriority(newPriority valueobject.Priority) error {
	priority, error := valueobject.NewPriority(newPriority)

	if error != nil {
		return error
	}

	i.priority = priority
	return nil
}

func (i *issueEntity) validateIssueCreation() error {
	if i.title == "" {
		return &IssueValidationError{Field: "title", Message: "title field cannot be empty"}
	}
	return nil
}

func (i *issueEntity) validateStatusTransition(newStatus valueobject.Status) error {
	switch i.status {
	case valueobject.StatusOpen:
		if newStatus == valueobject.StatusReopened {
			return &InvalidIssueStateTransitionError{
				From: i.Status().ToString(),
				To:   newStatus.ToString(),
			}
		}
	case valueobject.StatusClosed:
		if newStatus != valueobject.StatusReopened && newStatus != valueobject.StatusResolved {
			return &InvalidIssueStateTransitionError{
				From: i.Status().ToString(),
				To:   newStatus.ToString(),
			}
		}
	case valueobject.StatusResolved:
		if newStatus != valueobject.StatusReopened {
			return &InvalidIssueStateTransitionError{
				From: i.Status().ToString(),
				To:   newStatus.ToString(),
			}
		}
	case valueobject.StatusReopened:
		if newStatus == valueobject.StatusOpen {
			return &InvalidIssueStateTransitionError{
				From: i.Status().ToString(),
				To:   newStatus.ToString(),
			}
		}
	case valueobject.StatusInProgress:
	// Allow any transition from InProgress or Reopened
	default:
		return &IssueValidationError{
			Field:   "status",
			Message: "invalid current status",
		}
	}
	return nil
}
