package repository

import (
	"github.com/google/uuid"
	"github.com/thiagodebastos/gofixit/domain/entity"
)

type Repository interface {
	Save(user *entity.User) error
	FindById(id uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Delete(id uuid.UUID) error
}
