package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {
	id := uuid.New()
	userName := "newUser1234"
	email := "user.email@email.com"
	password := "pass1234!"
	name := "James"
	roles := []Role{{
		Name:        "developer",
		Permissions: []string{"create_issue", "delete_issue", "edit_issue"},
	}}

	user, err := CreateUser(userName, email, password, name, id, roles)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID() != id {
		t.Errorf("want %v got %v", id, user.ID())
	}
}
