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
	ChangeStatus(newStatus valueobject.Status) error
	UpdateTitle(newTitle string) error
	UpdateDescription(newDescription string)
}

type issueEntity struct {
	title       string
	description string
	status      valueobject.Status
	id          uuid.UUID
	priority    valueobject.Priority
}

var ErrInvalidStatus = errors.New("invalid status")

// NewIssue creates a new issue
func NewIssue(id uuid.UUID, title string, description string, status valueobject.Status) (*issueEntity, error) {
	return &issueEntity{
		id:          id,
		title:       title,
		status:      status,
		description: description,
	}, nil
}

func (i *issueEntity) ID() uuid.UUID              { return i.id }
func (i *issueEntity) Title() string              { return i.title }
func (i *issueEntity) Status() valueobject.Status { return i.status }
func (i *issueEntity) Description() string        { return i.description }

func (i *issueEntity) UpdateTitle(newTitle string) error {
	if i.Title() == "" {
		return errors.New("title cannot be empty")
	}

	i.title = newTitle
	return nil
}

func (i *issueEntity) SetStatus(newStatus valueobject.Status) error {
	switch i.status {
	case valueobject.StatusOpen:
		if newStatus == valueobject.StatusReopened {
			return errors.New("an Open ticket cannot be moved to Reopened")
		}
	case valueobject.StatusClosed:
		if newStatus != valueobject.StatusReopened && newStatus != valueobject.StatusResolved {
			return errors.New("an Closed ticket can only be moved to Reopened or Resolved")
		}
	case valueobject.StatusResolved:
		if newStatus != valueobject.StatusReopened {
			return errors.New("an Resolved ticket can only be moved to Reopened")
		}
	case valueobject.StatusReopened:
		if newStatus == valueobject.StatusOpen {
			return errors.New("a Reopened ticket cannot be moved to Opened")
		}
	case valueobject.StatusInProgress:
		// Allow any transition from InProgress or Reopened
	default:
		return errors.New("invalid current status")

	}

	i.status = newStatus
	return nil
}

func (i *issueEntity) SetPriority(newPriority valueobject.Priority) error {
	i.priority = newPriority
	return nil
}
