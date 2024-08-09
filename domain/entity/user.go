package entity

import (
	"github.com/google/uuid"
	"github.com/thiagodebastos/gofixit/domain/validation"
)

type Role struct {
	Name        string
	Permissions []string
}

type User interface {
	ID() uuid.UUID
	UserName() string
	Email() string
	Password() string
	Name() string
	Roles() []Role
}

type userEntity struct {
	userName string
	email    string
	password string
	name     string
	roles    []Role
	id       uuid.UUID
}

// Email implements User.
func (u *userEntity) Email() string {
	return u.email
}

// ID implements User.
func (u *userEntity) ID() uuid.UUID {
	return u.id
}

// Name implements User.
func (u *userEntity) Name() string {
	return u.name
}

// Password implements User.
func (u *userEntity) Password() string {
	return u.password
}

// Roles implements User.
func (u *userEntity) Roles() []Role {
	return u.roles
}

// UserName implements User.
func (u *userEntity) UserName() string {
	return u.userName
}

func CreateUser(
	userName string,
	email string,
	password string,
	name string,
	id uuid.UUID,
	roles []Role,
) (User, error) {
	if err := validation.ValidatePassword(password); err != nil {
		return nil, err
	}
	return &userEntity{
		userName: userName,
		email:    email,
		password: password,
		name:     name,
		id:       id,
		roles:    []Role{},
	}, nil
}
