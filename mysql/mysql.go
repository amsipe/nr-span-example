package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// Store is used to access the DB
type Store struct {
	db *sqlx.DB
}

// New returns a new instance of Store
func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s *Store) GetUsers(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	err := s.db.SelectContext(ctx, &users, "SELECT * FROM example.users")
	return users, err
}

func (s *Store) GetUser(ctx context.Context, id int) (*User, error) {
	user := User{}
	err := s.db.GetContext(ctx, &user, "SELECT * FROM example.users where id = ?", id)
	return &user, err
}
