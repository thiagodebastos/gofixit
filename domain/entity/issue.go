package entity

import (
	"errors"

	"github.com/google/uuid"
	"github.com/thiagodebastos/gofixit/domain/valueobject/issue"
)

type Issue interface {
	ID() uuid.UUID
	Title() string
	Description() string
	Status() issue.Status
	ChangeStatus(newStatus issue.Status) error
	UpdateTitle(newTitle string) error
	UpdateDescription(newDescription string)
}

type issueEntity struct {
	title       string
	description string
	status      issue.Status
	id          uuid.UUID
}

var ErrInvalidStatus = errors.New("invalid status")

// NewIssue creates a new issue
func NewIssue(id uuid.UUID, title string, description string, status issue.Status) (*issueEntity, error) {
	if !issue.ValidStatus(status) {
		return nil, ErrInvalidStatus
	}
	return &issueEntity{
		id:          id,
		title:       title,
		status:      status,
		description: description,
	}, nil
}

func (i *issueEntity) ID() uuid.UUID        { return i.id }
func (i *issueEntity) Title() string        { return i.title }
func (i *issueEntity) Status() issue.Status { return i.status }
func (i *issueEntity) Description() string  { return i.description }

func (i *issueEntity) UpdateTitle(newTitle string) error {
	if i.Title() == "" {
		return errors.New("title cannot be empty")
	}

	i.title = newTitle
	return nil
}

func (i *issueEntity) ChangeStatus(newStatus issue.Status) error {
	switch i.status {
	case issue.StatusOpen:
		if newStatus == issue.StatusReopened {
			return errors.New("an Open ticket cannot be moved to Reopened")
		}
	case issue.StatusClosed:
		if newStatus != issue.StatusReopened && newStatus != issue.StatusResolved {
			return errors.New("an Closed ticket can only be moved to Reopened or Resolved")
		}
	case issue.StatusResolved:
		if newStatus != issue.StatusReopened {
			return errors.New("an Resolved ticket can only be moved to Reopened")
		}
	case issue.StatusReopened:
		if newStatus == issue.StatusOpen {
			return errors.New("a Reopened ticket cannot be moved to Opened")
		}
	case issue.StatusInProgress:
		// Allow any transition from InProgress or Reopened
	default:
		return errors.New("invalid current status")

	}

	i.status = newStatus
	return nil
}

func (i *issueEntity) ValidStatus(status issue.Status) bool {
	panic("not implemented")
}
