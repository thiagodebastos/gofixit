package sqlite

import (
	"fmt"

	"github.com/thiagodebastos/gofixit/domain/entity"
	"zombiezen.com/go/sqlite"
)

type UserRepository struct {
	conn *sqlite.Conn
}

func NewUserRepository(conn *sqlite.Conn) *UserRepository {
	return &UserRepository{conn: conn}
}

func (r *UserRepository) GetUser(id int) (*entity.User, error) {
	stmt := r.conn.Prep("SELECT id, name FROM users WHERE id = ?")
	stmt.BindInt64(1, int64(id))
	hasRow, err := stmt.Step()
	if err != nil {
		return nil, err
	}
	if !hasRow {
		return nil, fmt.Errorf("no user found with id %d", id)
	}

	user := &entity.User{
		ID:   int(stmt.ColumnInt64(0)),
		Name: stmt.ColumnText(1),
	}

	return user, nil
}

func (r *UserRepository) CreateUser(user entity.User) error {
	stmt := r.conn.Prep("INSERT INTO users (id, userName, email, password, name) VALUES (?, ?, ?, ?, ?)")
	stmt.BindText(2, user.UserName())
	stmt.BindText(3, user.Email())
	stmt.BindText(4, user.Password())
	stmt.BindText(5, user.Name())
	_, err := stmt.Step()
	if err != nil {
		return err
	}
	return nil
}
